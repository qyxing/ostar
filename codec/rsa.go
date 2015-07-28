package codec

import (
	"bytes"
	"crypto/rsa"
	"encoding/base64"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"time"
)

const (
	MODE_PUBKEY_ENCRYPT = iota //公钥加密
	MODE_PUBKEY_DECRYPT        //公钥解密
	MODE_PRIKEY_ENCRYPT        //私钥加密
	MODE_PRIKEY_DECRYPT        //私钥解密
)

var (
	RSA = &rSASecurity{}
)

type rSASecurity struct {
	isFileInit bool            //true: 从文件读取密钥钥, false:字符串密钥
	ifCache    bool            //isFileInit＝＝true时有效。 true:只在初始化时读取密钥, false:每次都从文件读取
	pubStr     string          //isFileInit＝=true:公钥文件路径， isFileInit＝=false:公钥字符串
	priStr     string          //isFileInit＝=true:私钥文件路径， isFileInit＝=false:私钥字符串
	pubkey     *rsa.PublicKey  //公钥
	prikey     *rsa.PrivateKey //私钥
	pubModTime time.Time       //isFileInit＝＝true&&ifCache＝＝false时有效, 公钥文件最后的修改时间
	priModTime time.Time       //isFileInit＝＝true&&ifCache＝＝false时有效, 私钥文件最后的修改时间
}

func (this *rSASecurity) String(in string, mode int) (string, error) {
	var inByte []byte
	var err error
	if mode == MODE_PRIKEY_ENCRYPT || mode == MODE_PUBKEY_ENCRYPT {
		inByte = []byte(in)
	} else if mode == MODE_PRIKEY_DECRYPT || mode == MODE_PUBKEY_DECRYPT {
		inByte, err = base64.StdEncoding.DecodeString(in)
		if err != nil {
			return "", err
		}
	} else {
		return "", errors.New("mode not found")
	}
	inByte, err = this.Byte(inByte, mode)
	if err != nil {
		return "", err
	}
	if mode == MODE_PRIKEY_ENCRYPT || mode == MODE_PUBKEY_ENCRYPT {
		return base64.StdEncoding.EncodeToString(inByte), nil
	} else {
		return string(inByte), nil
	}
}

func (this *rSASecurity) Byte(in []byte, mode int) ([]byte, error) {
	out := bytes.NewBuffer(nil)
	err := this.IO(bytes.NewReader(in), out, mode)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(out)
}

func (this *rSASecurity) File(srcPath, distPath string, mode int) error {
	in, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer in.Close()
	isSuccess := false
	defer func() {
		if !isSuccess {
			os.Remove(distPath)
		}
	}()
	out, err := os.OpenFile(distPath, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0664)
	if err != nil {
		return err
	}
	defer out.Close()
	err = this.IO(in, out, mode)
	if err == nil {
		isSuccess = true
	}
	return err
}

func (this *rSASecurity) IO(in io.Reader, out io.Writer, mode int) error {
	switch mode {
	case MODE_PUBKEY_ENCRYPT:
		if key, err := this.getPubKey(); err != nil {
			return err
		} else {
			return pubKeyIO(key, in, out, true)
		}
	case MODE_PUBKEY_DECRYPT:
		if key, err := this.getPubKey(); err != nil {
			return err
		} else {
			return pubKeyIO(key, in, out, false)
		}
	case MODE_PRIKEY_ENCRYPT:
		if key, err := this.getPriKey(); err != nil {
			return err
		} else {
			return priKeyIO(key, in, out, true)
		}
	case MODE_PRIKEY_DECRYPT:
		if key, err := this.getPriKey(); err != nil {
			return err
		} else {
			return priKeyIO(key, in, out, false)
		}
	default:
		return errors.New("mode not found")
	}
}

func NewRSASecurity(pubStr, priStr string) (rsa *rSASecurity, pubKeyErr, priKeyErr error) {
	rsa = &rSASecurity{}
	pubKeyErr, priKeyErr = rsa.Init(pubStr, priStr)
	return
}

func NewRSASecurityByFile(pubFile, priFile string, ifCache bool) (rsa *rSASecurity, pubKeyErr, priKeyErr error) {
	rsa = &rSASecurity{}
	pubKeyErr, priKeyErr = rsa.InitByFile(pubFile, priFile, ifCache)
	return
}

func (this *rSASecurity) Init(pubStr, priStr string) (pubkeyErr, prikeyErr error) {
	this.isFileInit = false
	this.pubStr = pubStr
	this.priStr = priStr
	this.ifCache = true
	this.pubkey, pubkeyErr = getPubKey([]byte(this.pubStr))
	this.prikey, prikeyErr = getPriKey([]byte(this.priStr))
	return
}

func (this *rSASecurity) InitByFile(pubFile, priFile string, ifCache bool) (pubkeyErr, prikeyErr error) {
	this.isFileInit = true
	this.pubStr = pubFile
	this.priStr = priFile
	this.ifCache = ifCache
	_, pubkeyErr = this.getPubKey()
	_, prikeyErr = this.getPriKey()
	return
}

func (this *rSASecurity) getPubKey() (*rsa.PublicKey, error) {
	if this.isFileInit && !this.ifCache {
		f, err := os.Stat(this.pubStr)
		if err != nil {
			return nil, err
		}
		if f.ModTime().Equal(this.pubModTime) {
			if this.pubkey == nil {
				return nil, ErrPublicKey
			}
			return this.pubkey, nil
		} else {
			in, err := ioutil.ReadFile(this.pubStr)
			if err != nil {
				return nil, err
			}
			this.pubkey, err = getPubKey(in)
			if err == nil {
				this.pubModTime = f.ModTime()
			}
			return this.pubkey, err
		}
	} else {
		if this.pubkey == nil {
			return nil, ErrPublicKey
		}
		return this.pubkey, nil
	}
}

func (this *rSASecurity) getPriKey() (*rsa.PrivateKey, error) {
	if this.isFileInit && !this.ifCache {
		f, err := os.Stat(this.priStr)
		if err != nil {
			return nil, err
		}
		if f.ModTime().Equal(this.priModTime) {
			if this.prikey == nil {
				return nil, ErrPrivateKey
			}
			return this.prikey, nil
		} else {
			in, err := ioutil.ReadFile(this.priStr)
			if err != nil {
				return nil, err
			}
			this.prikey, err = getPriKey(in)
			if err == nil {
				this.priModTime = f.ModTime()
			}
			return this.prikey, err
		}
	} else {
		if this.prikey == nil {
			return nil, ErrPrivateKey
		}
		return this.prikey, nil
	}
}
