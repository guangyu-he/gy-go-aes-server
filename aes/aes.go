package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

// PKCS7填充
func PKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// 去除PKCS7填充
func PKCS7UnPadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}

// AES加密
func AESEncrypt(plainText, key []byte) (string, error) {
	// 创建AES块
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 使用PKCS7填充明文
	plainText = PKCS7Padding(plainText, block.BlockSize())

	// 生成随机IV
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// 创建CBC加密模式
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[aes.BlockSize:], plainText)

	// 返回Base64编码的加密结果
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// AES解密
func AESDecrypt(encrypted string, key []byte) (string, error) {
	// Base64解码
	cipherText, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	// 创建AES块
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 提取IV
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	// 创建CBC解密模式
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)

	// 去除PKCS7填充
	plainText := PKCS7UnPadding(cipherText)

	return string(plainText), nil
}
