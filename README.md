# ostar
pUBLICK_KEY:公钥字符串

pRIVATE_KEY: 私钥字符串

pubErr, priErr := codec.RSA.Init(pUBLICK_KEY, pRIVATE_KEY)

//pubErr, priErr = codec.RSA.InitByFile(pubFile, priFile, false)

log.Println("init error:", pubErr, priErr)

私钥加密

str, err := codec.RSA.String("golang", codec.MODE_PRIKEY_ENCRYPT)

公钥解密

str, err = codec.RSA.String(str, codec.MODE_PUBKEY_DECRYPT)

公钥加密

str, err = codec.RSA.String(str, codec.MODE_PUBKEY_ENCRYPT)

私钥解密

str, err = codec.RSA.String(str, codec.MODE_PRIKEY_DECRYPT)


byte操作

codec.RSA.Byte(in []byte, mode int)(out []byte, err error)

字符串操作

codec.RSA.String(in string, mode int)(out string, err error)

文件操作

codec.RSA.File(inFile, outFile string)(err error)

ReaderWrite操作

codec.RSA.IO(in io.Reader, out io.Writer)(err error)

