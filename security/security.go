package security

import (
	"bytes"
	"fmt"
	"strings"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/md5"
	"crypto/hmac"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"../random"
)


//Hmacmd5 Hmac-md5 crypto
func Hmacmd5(data string, key string) string {
	mac := hmac.New(md5.New, []byte(key))
	mac.Write([]byte(data))
	hash := mac.Sum(nil)
	return fmt.Sprintf("%x", hash)
}

//Md5 md5
func Md5(data string) string {
	hash := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", hash)
}

//Md5 md5
func Md5_16(data string) string {
	hash := Md5(data)
	return hash[8: 8 + 16]
}

//Sha256 sha256
func Sha256(data string) string {
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

//Sha1 sha1
func Sha1(input string) string {
	hash := sha1.Sum([]byte(input))
	return fmt.Sprintf("%x", hash)
}


//Hmacsha1 Hmac-sha1 crypto
func Hmacsha1(data string, key string) string {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(data))
	hash := mac.Sum(nil)
	return fmt.Sprintf("%x", hash)
}


//SaltPassword saltpassword
func SaltPassword(inputPassword string) string {
	hash := random.String(32)
	return Sha1(inputPassword + hash) + ":" + hash;
}

//VerifySaltPassword verify salt password
func VerifySaltPassword(inputPassword, saltPassword string) bool {
	hashValues := strings.Split(saltPassword, ":")
	if len(hashValues) != 2 {
		return false
	}
	checkHash := Sha1(inputPassword + hashValues[1])
	//log.Printf("%v, %v, %v, %v", checkHash, hashValues[0], hashValues[1], checkHash == hashValues[0])
	return checkHash == hashValues[0]
}

//AesEncode aes encode
func AesEncode(data, key string) (string, error) {
	hashKey := Md5(key)
	bys, err := AesEncrypt([]byte(data), []byte(hashKey))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bys), nil
}

//AesDecode aes decode
func AesDecode(data, key string) (string, error) {
	hashKey := Md5(key)
	unBase64, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	ret, err := AesDecrypt(unBase64, []byte(hashKey))
	if err != nil {
		return "", err
	}
	return string(ret), nil
}

//AesEncrypt aes encrypt
func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	iv := key[:blockSize]
	fmt.Println("AesEncrypt key:", string(key), "|", string(iv))
	origData = Pkcs5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

//AesDecrypt aes decrypt
func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	iv := key[:blockSize]
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = Pkcs5UnPadding(origData)
	return origData, nil
}

//Pkcs5Padding pkcs5padding
func Pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext) % blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//Pkcs5UnPadding pkcs5 unpadding
func Pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length - 1])
	end := length - unpadding
	if end > (length - 1) {
		end = length - 1
	}
	if end < 0 {
		end = 0
	}
	return origData[:end]
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext) % blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	suffix := byte(0)
	paddingLen := 0
	length := len(origData)
	for i := (length - 1); i >= 0; i--{
		if origData[i] != suffix{
			break
		}
		paddingLen++
	}
	return origData[:(length - paddingLen)]
}

// Padding补全
func PKCS7Pad(data []byte) []byte {
	padding := aes.BlockSize - len(data) % aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

func PKCS7UPad(data []byte) []byte {
	padLength := int(data[len(data)-1])
	return data[:len(data)-padLength]
}