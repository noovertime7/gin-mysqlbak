package main

import (
	"fmt"
	"github.com/yuchenfw/gocrypt"
	"github.com/yuchenfw/gocrypt/rsa"
)

func main() {
	secretInfo := rsa.RSASecret{
		PublicKey:          "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAyoiAraTnAbCoqGVOKugFDM2/ms2szXmb3zTOU3ByicH/XPZqy7Eougbs8OQQIoNW4xKw8PNyWf0lfr90qBfPj27INn6N7umVmbHCNCKkQ4frPn46xesw1ywtc2GhOEzZlC8ajlnzBUkj5FJZcrNjXfFmfsQcFQP0g/o/3CAUpk1BXFXt7eZsaYdyn0m7fMoyFt1wlF8egQeGYYE98vtKsvrII51HK8vOEf+5VXU4UZxGfvyzS3A8kuNEkKEh1n9mazjfPBT0KGSiOGh7Nugks+jjfswSgXRK/b2eP3fS7U625rbS798pKxnoS2E0Pgpzdk5fWoNgAlG/n2F9oI2/kQIDAQAB",
		PublicKeyDataType:  gocrypt.Base64,
		PrivateKey:         "MIIEowIBAAKCAQEAyoiAraTnAbCoqGVOKugFDM2/ms2szXmb3zTOU3ByicH/XPZqy7Eougbs8OQQIoNW4xKw8PNyWf0lfr90qBfPj27INn6N7umVmbHCNCKkQ4frPn46xesw1ywtc2GhOEzZlC8ajlnzBUkj5FJZcrNjXfFmfsQcFQP0g/o/3CAUpk1BXFXt7eZsaYdyn0m7fMoyFt1wlF8egQeGYYE98vtKsvrII51HK8vOEf+5VXU4UZxGfvyzS3A8kuNEkKEh1n9mazjfPBT0KGSiOGh7Nugks+jjfswSgXRK/b2eP3fS7U625rbS798pKxnoS2E0Pgpzdk5fWoNgAlG/n2F9oI2/kQIDAQABAoIBAF378hqiR0CVhe5+9EMc4BsM7zka8HF5WUe+7W/y4nPivmmZP/29/DQ3OoSekI4zfIJrDgkCL7JqspeaqLvIMN1Sfz4qhBq18mIcBw7CdI+R5yxcz1FAzq1LJtxAFdxWbTFCmoQsYYW2Zx1wyWlcrWPOvc1dm9p0t2b3HeM8T9jLdY+D0Bm9zmAS0nwTuDBxYS77DB9Ncl6pWLLd197/5IoN1/nunFuzpkiwMPI9RF7lgrnUthc/1Gfnylz5/tXCiQsEVSbAdbMXt9nsV0RgVeMcPq/aUqTMLS2lIV8JySWDrRQi4yPHU0hIjcp6ggo53YMuncJZweI/wwkJexojz0ECgYEA5QzRObpU0CryfJ7qa97/USIKHbvl6PuQG9OLyUeP9bG0edidQhUrR4EZwjIl73O8CTJ0bB24wAKZZEOK3eJeqG/N0q+CiD83ygr8pSZzpE1xvqQp32IgXtgvm7/UmT8cfAp05Z3bF4jcA8uXwodBz4NsVGijlO78PsCooLsArM0CgYEA4lz5pXDEN3w5JwkbspLnUSUS738hne8YM0PchCaww+8sXLS9GLL2CHcvwh6Tv9Mee7r6SdbDI73x118y68WEDDhidiYZCLhXJN2v12ezJOMqH5m9wVJzQOGNv6kPV1EW1WlWxoJQGxCdzbZMLxtTbyTZe3+iAVG++8u6NWMV3dUCgYA1dm1rnQto321kGy+6Z/2OMXTNBeufGwDDDfilzZdTkNwASMhEAW7trLuXcV8bahcsymMUTUevQawOFBnYupq/lAEluSOtq5vZBAF+huAdLJptFiJT6rKFkM5j+z2jW3DJnyMz6UmXT7GTDTVqCWoaBqIFfbsY60NjXlK92YhJzQKBgQDWfQjktbSHasLw9RV0oPRklD+cBhfBgfOpZ+0En3CxR+j+MxhW1gSBQwZS5wxTIGXrEeHlo4UmUe5diExE0dRsi+ToVPM1qw6P1SuwbQd3tXSNmu0NyOWCnfblm/j4YNLFB1p9IK9s5dLRQKJxpG/ribw15FuK6n2QM5vOyIPIvQKBgE5PUzRUCCVsjKAxZOfaZQatMbSzAUSB3bNmUw+F3pDq8ibs6XXvtySowG2femlPDNL7mDMuUc9kYrtTFTQNrEsQGB55wBopX3UxzRjpXJoAQ/d+RPdrSJC7xJyu+URoFI6ae0I3bx1BzjctYU0Rv5DUh+j9leMH5N2S9vHb+vqu",
		PrivateKeyType:     gocrypt.PKCS1,
		PrivateKeyDataType: gocrypt.Base64,
	}
	handleRSA := rsa.NewRSACrypt(secretInfo) //RSA

	encrypt, err := handleRSA.Encrypt("testaaaaa", gocrypt.HEX)
	if err != nil {
		fmt.Println("encrypt error :", err)
		return
	}
	fmt.Println("encrypt data :", encrypt)
	decrypt, err := handleRSA.Decrypt(encrypt, gocrypt.HEX)
	if err != nil {
		fmt.Println("decrypt error :", err)
		return
	}
	fmt.Println("decrypt data :", decrypt)
}
