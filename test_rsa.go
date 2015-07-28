package main

import (
	"log"
	"qiniu/codec"
)

const (
	pRIVATE_KEY = `
-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAtolbxCWZyqQUXA3X5fueJwZ+qsMp4UNtdXhoXHJRVXGRfbBT
lMQB/nH1DXP1FJ4QNuIu7eeOxe44dPNAAgPdEY4Ys9LVwsPhuZ58kwUK0ZUZbLXY
+xS2yhzWhYsZywztBMKa1t7YcQSnsaehHoDVixAFJ50Gcl8qWFfn8E55MY5Y6C75
BJDXp3rDrayWFi24GpgqRRlECbRwk3PN9YQf104yktj530x81zR2PAiQLV4cIRt4
11eZ7f9fJEQywuEnVpSuVsjK+LnX4KiTWxTLGaa63NWdXM4Jwe1qn6SkkuU9ZSyY
Y1bwonr7FfeOueFQ3VWogOpzhERdp0Ws2V//lQIDAQABAoIBAQCgsdXunOeCROaD
j9BUcCnv00Dp1fxlinWvZ7wAPsepf9yEmRlLqy7SMMJ6AG5uoyRFHOQRnrvLNgfP
tWHRJFOXI9BNZru2xblPLt4ek97NWQRT9tc5Wyf8UFzuneGsJwn2Gdg0d2R2QpHa
zWcDMopL7WOMVymYwHzH30OaA4uf+mM/oa16xkB7ixWd18pqOQvwrHpWWURwLovu
NsDxWSQ+oSM7yjIVG73sEIHwEdOE5VvPDMd6LhTCSOQKZ9L80Jre9z4hwr2EaEnp
AfoYsCQzmdUldjTRjUbv6unxDdXIQccB2bw5CqA5XnFHhgU57p3whvlsysK0RuEm
j0r1fa01AoGBAN3M976nWGnUe4JlrVVsPBboNGI8UY3vciPHx+cR27BK1100XP5W
EX/6gaoY/fUq3WkV2qECglyQ7FbxI6V3GTXtrjY8GHgRu2rWd6VFjDU0Yf/nQfpl
RXS3NABFcOsNyso0OmjKRY+D+41BiF35D7TeKeFW+n5yE7g7gDM/455XAoGBANKu
iZtwBY+GbE9bKwhWwk6Xk5rMbrnIjKOb6WLwBPgo8lZJIYX/6lsskJZP3cdvfg5h
YTbKXiSf/gDHjvoZs3QpFKDOVdPRY+2gXJqsDHY7v3d+8/hEbtTVZklYPwtl115K
H5l+0GfFOWmouQ5vyecTpPA4ubc2gVwrw40p7wXzAoGBAKqcvAWf3Far58XKSKbo
9t/4BjN6ipFPmtEDIDYSepcFOtyrJs1Nj3COVaduSguIyX/IG8C2mWhy4hmOrAjf
sDjXd6aoW3ogybXI+4faE5vpi2i5jvr5Y5AATLPYtp9YoKEhw7xPu2pF7/4cZrVC
nF5YdoarzUvunFSfEGJbxs9JAoGAJMkq58QIhIXxFW4StnMHnFdlA2tcjf3RaKPJ
fWfxRi9IGP7N5qrHjcHbQROS4sa52OLx6XIuO/Dfld1CPrMMHWUq3+UHIWP3Mb+F
S9BsoJxQExpMmPXB8FGOeZH5+BCBKUqB9/gnhWbvXl6CaV3lf/5zFyqgargOoDxX
+abvwDcCgYAIpujLEeRfs5trOQovHF3NPEY8X9r59t+BJ3p+baDbC4n/v7+OETdz
E9gKsIerePhlRulpQNr5MBuPOvKdnwBn36rmuPFKd2QkY0YSc9GVIbmc2Iq3ihX5
2yggRe4ntA3NhezBSAZD50Qf7NCVKzMp92wnhPsMBxkq3/jSX5+e1A==
-----END RSA PRIVATE KEY-----`
	pUBLICK_KEY = `
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtolbxCWZyqQUXA3X5fue
JwZ+qsMp4UNtdXhoXHJRVXGRfbBTlMQB/nH1DXP1FJ4QNuIu7eeOxe44dPNAAgPd
EY4Ys9LVwsPhuZ58kwUK0ZUZbLXY+xS2yhzWhYsZywztBMKa1t7YcQSnsaehHoDV
ixAFJ50Gcl8qWFfn8E55MY5Y6C75BJDXp3rDrayWFi24GpgqRRlECbRwk3PN9YQf
104yktj530x81zR2PAiQLV4cIRt411eZ7f9fJEQywuEnVpSuVsjK+LnX4KiTWxTL
Gaa63NWdXM4Jwe1qn6SkkuU9ZSyYY1bwonr7FfeOueFQ3VWogOpzhERdp0Ws2V//
lQIDAQAB
-----END PUBLIC KEY-----
`
)

func test_rsa() {
	pubErr, priErr := codec.RSA.Init(pUBLICK_KEY, pRIVATE_KEY)
	log.Println("init error:", pubErr, priErr)

	str, err := codec.RSA.String("OKOK", codec.MODE_PRIKEY_ENCRYPT)
	log.Println("prikey encrypt:", str, err)
	str, err = codec.RSA.String(str, codec.MODE_PUBKEY_DECRYPT)
	log.Println("pubkey decrypt:", str, err)
	str, err = codec.RSA.String(str, codec.MODE_PUBKEY_ENCRYPT)
	log.Println("pubkey encrypt:", str, err)
	str, err = codec.RSA.String(str, codec.MODE_PRIKEY_DECRYPT)
	log.Println("private encrypt:", str, err)

}
