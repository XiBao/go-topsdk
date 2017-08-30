package util

import (
	"crypto/hmac"
	"crypto/md5"
)

func HmacMd5(input []byte, key string) ([]byte, error) {
	bKey, err := Base64Decode(key)
	if err != nil {
		return nil, err
	}

	mac := hmac.New(md5.New, bKey)
	mac.Write(input)
	return mac.Sum(nil), nil
}

func HmacMD5EncryptToBase64(input, key string, compressLen int) (string, error) {
	encryptRet, err := HmacMd5([]byte(input), key)
	if err != nil {
		return "", err
	}
	if compressLen > 0 {
		encryptRet = Compress(encryptRet, compressLen)
	}
	return Base64Encode(encryptRet), nil
}
