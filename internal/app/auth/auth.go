package auth

import (
	"bytes"
	"crypto/aes"
	"encoding/hex"
	"encoding/json"
	"time"
)

type Authenticator struct {
	Key []byte
}

type Token struct {
	Email   string    `json:"email"`
	Issued  time.Time `json:"issued"`
	Expires time.Time `json:"expires"`
}

func NewAuthenticator(privateKey string) *Authenticator {
	return &Authenticator{Key: []byte(privateKey)}
}

func NewToken(email string) *Token {
	return &Token{Email: email}
}

func (auth *Authenticator) Encode(token *Token) (string, error) {

	token.Issued = time.Now()
	token.Expires = time.Now().Add(time.Hour * 12)

	var cipher, err = aes.NewCipher(auth.Key)
	if err != nil {
		return "", err
	}

	js, err := json.Marshal(token)
	if err != nil {
		return "", err
	}

	var ciphertext []byte
	ciphertext = reslice(js, ciphertext, cipher.Encrypt)
	return hex.EncodeToString(ciphertext), nil
}

func (auth *Authenticator) Decode(tokenHex string) (Token, error) {
	var token Token

	var ciphertext, err = hex.DecodeString(tokenHex)
	if err != nil {
		return token, err
	}

	cipher, err := aes.NewCipher(auth.Key)
	if err != nil {
		return token, err
	}

	var decryptedToken = make([]byte, 0)
	decryptedToken = reslice(ciphertext, decryptedToken, cipher.Decrypt)

	err = json.Unmarshal(decryptedToken, &token)

	return token, err
}

type CryptFunc func([]byte, []byte)

func reslice(input, output []byte, cb CryptFunc) []byte {
	var block = make([]byte, aes.BlockSize)

	for len(input) > 0 {
		if len(input) < aes.BlockSize {
			input = PKCS5Padding(input, aes.BlockSize)
			cb(block, input)
			output = append(output, block...)
		} else {
			cb(block, input[0:aes.BlockSize])
			output = append(output, block...)
		}
		input = input[aes.BlockSize:]
	}

	return output
}

func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
