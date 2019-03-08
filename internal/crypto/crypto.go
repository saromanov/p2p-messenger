// Package crypto provides handling of the crypto algorithms
package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	"github.com/pkg/errors"
	"golang.org/x/crypto/curve25519"
	"golang.org/x/crypto/ed25519"
)

// MakeKeyPair provides creating of the key pair for peers
func MakeKeyPair() ([32]byte, [32]byte, error) {
	pubKey, privKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		return [32]byte{}, [32]byte{}, err
	}

	var (
		publicKey  [32]byte
		privateKey [32]byte
	)

	copy(publicKey[:], pubKey[:])
	copy(privateKey[:], privKey[:])
	curve25519.ScalarBaseMult(&publicKey, &privateKey)

	return publicKey, privateKey, nil
}

// Encrypt provides encrypting of the data
func Encrypt(data, key []byte) ([]byte, error) {
	block, _ := aes.NewCipher(key)
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get cipher block")
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, errors.Wrap(err, "unable to read bytes")
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}
