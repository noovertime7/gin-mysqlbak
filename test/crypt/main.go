package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"log"
)

func main() {
	var test = `Nb3tfHl6sFUyB72+/uMYSKxqSoniiWmmDl4FMBfvObf7skJOcYwl7do8/VPHRCku9BYiihP44kbOGZp5bVtnIsLaG6uYygvyr8k1ZtsXE627H77QKfG9X7l96/PCyI9ZaMemJL3vB66HbvxEM6KThvPPqV8kQK1H7F5Efoxugkv7qPHWuVDWchkTwafytOEqa6K3t6A+5NQbeEA5AQhfKZJ/BDlmRF/4F363VQbEa55mS9hbXtmK5gG0wNe6BGmW40hSpUs4vfIsnVRehljXMn9KkKZZ+eN3y0uL/TcpJ0HTF1nZvdMKVywNwBS8D1/fmaEGyJS5K6gH/185yaorwQ==`
	fmt.Println(RsaDecode(test))
}

func RsaEncode(plain string) string {
	msg := []byte(plain)
	publicKey :=
		`-----BEGIN Public Key-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA783z1uGIn78nmRnlfoP5
RsnbmdhB+tDoyxAKAPVSrWenMdgJFvn2RcaiQBeH2BbKNQcP/ygYeOEC2Zhb+EZH
AwyT0Zo34vZQAgVWOi5zVDNwxu4VPfc+25+YNAr+Zy/txqJlWF5EbPyqjSZGQQI/
xwZSvgW76ucz2vj2LKoyddESpbmV0QQYVxvQ1gHoBODHXuc6dDwQZM9cPWh/N/nH
9I45Ty0ZWCTOj9qCk/92ChLaI/hY4552yLDMrzsqEqrL0KsCQKqnTyCeUgqvHFxG
K0zZt35ob0C7p1FYJ4HHHfTHV6T0a8U1CYua5NJNi5mP+Nb6KIyuWCHrQGeKeN75
CwIDAQAB
-----END Public Key-----
`
	pubBlock, _ := pem.Decode([]byte(publicKey))
	pubKeyValue, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if err != nil {
		log.Println(err)
		return ""
	}
	pub := pubKeyValue.(*rsa.PublicKey)
	encryText, err := rsa.EncryptPKCS1v15(rand.Reader, pub, msg)
	if err != nil {
		log.Println(err)
		return ""
	}
	return base64.StdEncoding.EncodeToString(encryText)
}

func RsaDecode(strPlainText string) string {
	plainText, err := base64.StdEncoding.DecodeString(strPlainText)
	privateKey := `-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDvzfPW4YifvyeZ
GeV+g/lGyduZ2EH60OjLEAoA9VKtZ6cx2AkW+fZFxqJAF4fYFso1Bw//KBh44QLZ
mFv4RkcDDJPRmjfi9lACBVY6LnNUM3DG7hU99z7bn5g0Cv5nL+3GomVYXkRs/KqN
JkZBAj/HBlK+Bbvq5zPa+PYsqjJ10RKluZXRBBhXG9DWAegE4Mde5zp0PBBkz1w9
aH83+cf0jjlPLRlYJM6P2oKT/3YKEtoj+FjjnnbIsMyvOyoSqsvQqwJAqqdPIJ5S
Cq8cXEYrTNm3fmhvQLunUVgngccd9MdXpPRrxTUJi5rk0k2LmY/41voojK5YIetA
Z4p43vkLAgMBAAECggEAGYDz3/SYjTTNR1EjwTLT/h1Vx6TiT4SMXZxVFAkDRAH0
HC73uIPZF06qztoxzl+OHdmkb+SZxbFYlj/H+D2xK7fYuMAIFZFQrQZYR1PNXDUk
V2PWyoJTIxR0IYTRzbOhPDDlSmKsMFMm6GAbtKpCki4v+pmthKwWLcGwnMt1FmZz
OmUsVRH7kMk/WhW3ooPw7/YThMImsfBh6mPOuiFoHDtG6slZbya4av7+FT9zWNXr
S3dNkUdTzxCmp/SWZZIezRTwGL/itou5m7j0J5OnBCaSq/IVGU5Qz2vmrw9x/42j
HeusYxJdiRABnGswNmq/LxUv25PnrJmP47167TxXsQKBgQD60ePjL2MTltDAvSw+
W/2BCpVSbR8nMqQz4Sy5o5qdcvgN+qy88vUmZ1TaVzfQzDrOsu8HC1pnO5U/IHvN
w9VdWc5LPcBjFEVbYmDP/u9CVSumhcovzn3vzyo0pPkI8SoXo/Ov55L21+7EZqhk
Wz2ecFGFYbaMkFxt60aAUzm9DwKBgQD0wdKox3K3hNYKK+TzxRxx5LNUkUyIFOSR
9sxxQKAFM/be7ohJRyLkHDPuv3pra3zzz43yQUcIVoL92iSSg46tmABMb49FrudN
iM1bKlyJ52hG3Jy/9aBEUxhL4FW5b9862NVAJ693SprnOb7JUqDvVfgD/bfSR9YS
spcilQi8RQKBgDGwGNwtzAvaabp0/2nPIXZJ2XD9yxkh0COy7QBNp9ifKQLj8Qpx
ex1DhSzH8He9rby6991GY39l71gVIFGQBRm8K8D+F7nJD0BeSd2KnG1goAgaIwa7
enliafMJ54fc9sC+Kw1i69tYxaVEzQEsjhVwHMY2qEpKYvUnkczkL+EDAoGBAOxP
MoAISScM1sbtSFlfxy3jnI0a2CRO37xZ51u1BIrheAoXjXgKeZJ59F4feye5mOOh
UjBCfD19oW1Mr6DR6fCK4EbsMw0ZaHNAkNzjUoNG7DP2jlT75ufwvnWLu9iVPZcI
gSQ27L+lRVfYNe8UmxNZEmNwFImvF+3nheo6sDttAoGASKudjo9aFw66fRMVoF+M
n3YlkD42axD4iPy3qSid5Ray/lloRcckuE/cf4oeJv9mYX1ihDTJa8/kVokGs4Bu
B42tPAAKInL+pUjagGH+w1YjpWBshp9ct9m3zrobbHhLzzUQPufNL7Im9mWw+zP4
qiwZPi0OAWEGVZqy5c3Ai8g=
-----END PRIVATE KEY-----`
	bytePrivateKey := []byte(privateKey)
	priBlock, _ := pem.Decode(bytePrivateKey)
	priKey, err := x509.ParsePKCS8PrivateKey(priBlock.Bytes)
	if err != nil {
		log.Println(err)
		return ""
	}
	decryptText, err := rsa.DecryptPKCS1v15(rand.Reader, priKey.(*rsa.PrivateKey), plainText)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(decryptText)
}
