package snippet

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"log"
	"strings"
	"testing"
)

var key = "a secret code must not be empty!"

func Encoder(uid string) string {
	return string(base64Encode([]byte(AesEncrypt(uid, key))))
}

func Decoder(utoken string) (string, error) {
	// decode base64
	enbyte, err := base64Decode([]byte(utoken))
	if err != nil {
		log.Println(err)
		return "", err
	}
	// decode aes
	return AesDecrypt(string(enbyte), key), nil
}

func TestLezenBAES(t *testing.T) {
	tests := []struct {
		uid string
	}{
		{"ZRZ"},
		{"hinabook"},
		{"51665133"},
		{"luckypeachchina"},
		{"chaosmiao"},
		{"bandaocvs"},
		{"boyvintagestore"},
	}
	for i, ts := range tests {
		token := Encoder(ts.uid)
		res, err := Decoder(token)
		log.Printf("%d-th uid=%s\t\ttoken=%s\n", i+1, ts.uid, token)
		// log.Println(""ts.uid, "= ", token)
		if err != nil || strings.Compare(ts.uid, res) != 0 {
			t.Fatal("expected=", ts.uid, "got=", res)
		}
	}
}

func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}

func TestBASE64(t *testing.T) {
	// encode
	hello := "/JLzx3LuO3Sr/eQrqD+6EA=="
	debyte := base64Encode([]byte(hello))
	fmt.Println(string(debyte))
	// decode
	enbyte, err := base64Decode(debyte)
	if err != nil {
		fmt.Println(err.Error())
	}

	if hello != string(enbyte) {
		fmt.Println("hello is not equal to enbyte")
	}

	fmt.Println(string(enbyte))
}

func TestAES(t *testing.T) {
	orig := "123456"
	key := "a secret code must not be empty!"
	fmt.Println("原文：", orig)
	encryptCode := AesEncrypt(orig, key)
	fmt.Println("密文：", encryptCode)
	decryptCode := AesDecrypt(encryptCode, key)
	fmt.Println("解密结果：", decryptCode)
}
func AesEncrypt(orig string, key string) string {
	// 转成字节数组
	origData := []byte(orig)
	k := []byte(key)
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = PKCS7Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)
	return base64.StdEncoding.EncodeToString(cryted)
}
func AesDecrypt(cryted string, key string) string {
	// 转成字节数组
	crytedByte, _ := base64.StdEncoding.DecodeString(cryted)
	k := []byte(key)
	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	orig = PKCS7UnPadding(orig)
	return string(orig)
}

//补码
//AES加密数据块分组长度必须为128bit(byte[16])，密钥长度可以是128bit(byte[16])、192bit(byte[24])、256bit(byte[32])中的任意一个。
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
