// Package crypto provides handling of the crypto algorithms
package crypto

import (
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
