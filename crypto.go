package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

type CryptoFunctions struct{}

func Crypto() CryptoFunctions {
	return CryptoFunctions{}
}

func (c CryptoFunctions) HexToString(key []byte) string {
	return hex.EncodeToString(key)
}

func (c CryptoFunctions) StringToHex(key string) ([]byte, error) {
	return hex.DecodeString(key)
}

func (c CryptoFunctions) RandomNumberGenerator(size int) ([]byte, error) {
	bytes := make([]byte, size)
	_, err := rand.Read(bytes)
	if err != nil {
		return []byte{}, err
	}

	return bytes, nil
}

func (c CryptoFunctions) Encrypt(text, key []byte) ([]byte, error) {
	// create a new cipher
	// 16 bytes key = AES-128
	// 24 bytes key = AES-192
	// 32 bytes key = AES-256
	if len(key) != 32 {
		return []byte{}, fmt.Errorf("to use AES-256 the key should be 32 bytes, got %d", len(key))
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}

	// using Galois/Counter mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return []byte{}, err
	}

	// create a nonce
	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return []byte{}, err
	}

	// encrypt
	encrypted := gcm.Seal(nonce, nonce, text, nil)
	return encrypted, nil
}

func (c CryptoFunctions) Decrypt(encrypted, key []byte) ([]byte, error) {
	// create a new cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}

	// using Galois/Counter mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return []byte{}, err
	}

	nonceSize := gcm.NonceSize()
	if len(encrypted) < nonceSize {
		return []byte{}, fmt.Errorf("encrypted text should be at least %d bytes", nonceSize)
	}

	nonce, cipherText := encrypted[:nonceSize], encrypted[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return []byte{}, err
	}

	return plaintext, nil
}
