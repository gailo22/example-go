package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/tuya/tuya-pulsar-sdk-go/pkg/tylog"
)

const (
	IV = "357608fc0d2941cb8a7a8f363c09a2d9"
)

func main() {

	str := "f0a4461763f5ad4b6b90da6425adc6bc0fa40fab01dfcf7da7af374866b18cd6e018f55eb0717ce8c3cf6e2adb6d2f3fcd8a4fe94f57b07a79ee5397ecf43b64a208e6fe5b217cbf21c91445a1cd87ad4cb66d68bd4b64e29129bb51b0d349ae5a040f14a4175f5c9214333084863ba3"

	// fmt.Println(decypt(str, IV))

	// accessKey := "42ee90ce9aed4bbf98c27d75e44bf6c8"
	accessKey := "2b3ca2a1e8314a3895775913444c8091"
	key := accessKey[8:24]

	de, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("base64 decode failed", tylog.ErrorField(err))
	}
	fmt.Println("decoded: ", string(de))
	data := Data{}
	err = json.Unmarshal(de, &data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("data: ", data)

	dee := data.Files[0].Data
	fmt.Println("dee: ", dee)
	// dc := tyutils.EcbDecrypt([]byte(dee), []byte(key))
	dc := decypt([]byte(dee), []byte(IV), []byte(key))
	fmt.Println("decypted data:", string(dc))
}

func decypt(data []byte, iv []byte, key []byte) []byte {

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// iv := data[0:aes.BlockSize]

	buffer := make([]byte, len(data)-len(iv))
	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypter.CryptBlocks(buffer, data[len(iv):])

	return PKCS7Unpadding(buffer)
}

func PKCS7Padding(ciphertext []byte) []byte {
	padding := aes.BlockSize - len(ciphertext)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7Unpadding(plantText []byte) []byte {
	length := len(plantText)
	padding := int(plantText[length-1])
	return plantText[:(length - padding)]
}

type Data struct {
	V     string `json:"v"`
	Files []File `json:"files"`
	Cmd   string `json:"cmd"`
	Type  string `json:"type"`
}

type File struct {
	Data  string `json:"data"`
	KeyID string `json:"keyId"`
	Iv    string `json:"iv"`
}
