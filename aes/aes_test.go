package aes

import (
	"fmt"
	"testing"
)

func TestAES(t *testing.T) {
	key := []byte("examplekey123456") // AES-128 (16字节)
	plainText := "Hello, AES in Golang!"

	// 加密
	encrypted, err := AESEncrypt([]byte(plainText), key)
	if err != nil {
		fmt.Println("加密失败:", err)
		return
	}
	fmt.Println("加密结果:", encrypted)

	// 解密
	decrypted, err := AESDecrypt(encrypted, key)
	if err != nil {
		fmt.Println("解密失败:", err)
		return
	}
	fmt.Println("解密结果:", decrypted)
}
