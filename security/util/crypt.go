package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

const LOCAL_IV = "0102030405060708"

func Encrypt(input, key string) (string, error) {
	bKey, err := Base64Decode(key)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(bKey) //选择加密算法
	if err != nil {
		return "", err
	}
	plantText := PKCS5Padding([]byte(input), block.BlockSize())
	blockModel := cipher.NewCBCEncrypter(block, []byte(LOCAL_IV))
	ciphertext := make([]byte, len(plantText))
	blockModel.CryptBlocks(ciphertext, plantText)
	return Base64Encode(ciphertext), nil
}

func Decrypt(input, key string) (string, error) {
	bKey, err := Base64Decode(key)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(bKey) //选择加密算法
	if err != nil {
		return "", err
	}

	ciphertext, err := Base64Decode(input)
	if err != nil {
		return "", err
	}
	blockModel := cipher.NewCBCDecrypter(block, []byte(LOCAL_IV))
	plantText := make([]byte, len(ciphertext))
	blockModel.CryptBlocks(plantText, ciphertext)
	plantText = PKCS5Unpadding(plantText)
	return string(plantText), nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5Unpadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
