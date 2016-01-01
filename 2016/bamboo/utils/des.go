// des util
// since 2016-01-01

package utils

import (
    "bytes"
    "crypto/cipher"
    "crypto/des"
    "encoding/base64"
)

var (
    de2key []byte
    iv     []byte
)

func init() {
    c := NewConfig("../conf/app.properties")
    de2key = []byte(c.GetValue("app.des.secretKey"))
    iv = []byte(c.GetValue("app.del.iv"))
}

// 3DES encode
// Read de2key and iv from an config file
func TripleDESEncode(plaintext []byte) string {

    block, err := des.NewTripleDESCipher(de2key)
    if err != nil {
        panic(err)
    }

    bs := block.BlockSize()
    plaintext = PKCS5Padding(plaintext, bs)

    if len(plaintext)%bs != 0 {
        panic("plaintext is not a multiple of the block size")
    }

    ciphertext := make([]byte, len(plaintext))
    mode := cipher.NewCBCEncrypter(block, iv)
    mode.CryptBlocks(ciphertext, plaintext)

    base := base64.StdEncoding.EncodeToString(ciphertext)

    return base
}

// 3DES decode
// Just like encode
func TripleDESDecode(plaintext string) string {

    ciphertext, err := base64.StdEncoding.DecodeString(plaintext)
    if err != nil {
        panic(err)
    }

    block, err := des.NewTripleDESCipher(de2key)
    if err != nil {
        panic(err)
    }

    bs := block.BlockSize()
    if len(ciphertext) < bs {
        panic("ciphertext too short")
    }

    // CBC mode always works in whole blocks.
    if len(ciphertext)%bs != 0 {
        panic("ciphertext is not a multiple of the block size")
    }

    mode := cipher.NewCBCDecrypter(block, iv)

    // CryptBlocks can work in-place if the two arguments are the same.
    mode.CryptBlocks(ciphertext, ciphertext)

    ciphertext = PKCS5UnPadding(ciphertext)
    return string(ciphertext)
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
    padding := blockSize - len(ciphertext)%blockSize
    padtext := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
    length := len(origData)
    unpadding := int(origData[length-1])
    return origData[:(length - unpadding)]
}
