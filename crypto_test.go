package utils

import (
	"testing"
)

func TestEncryption(t *testing.T) {
	text := "this is the text to encrypt"

	key, err := Crypto().RandomNumberGenerator(32)
	Test().Nil(t, err)

	encrypted, err := Crypto().Encrypt([]byte(text), key)
	Test().Nil(t, err)

	plaintext, err := Crypto().Decrypt(encrypted, key)
	Test().Nil(t, err)
	Test().Equals(t, plaintext, []byte(text))
}
