package crypto

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
)

// Signer handles cryptographic signing operations for transactions
type Signer struct {
	privateKey *ecdsa.PrivateKey
}

// NewSigner creates a new signer with a given private key
func NewSigner(privKey *ecdsa.PrivateKey) *Signer {
	return &Signer{
		privateKey: privKey,
	}
}

// SignMessage signs arbitrary data and returns hex-encoded signature
func (s *Signer) SignMessage(data []byte) (string, error) {
	hash := sha256.Sum256(data)

	r, sig, err := ecdsa.Sign(rand.Reader, s.privateKey, hash[:])
	if err != nil {
		return "", fmt.Errorf("signing failed: %w", err)
	}

	// Encode signature as hex
	signature := append(r.Bytes(), sig.Bytes()...)
	return "0x" + hex.EncodeToString(signature), nil
}

// VerifySignature verifies a signature against public key
func VerifySignature(pubKey *ecdsa.PublicKey, data []byte, signature string) bool {
	hash := sha256.Sum256(data)

	// Decode hex signature
	sigBytes, err := hex.DecodeString(signature[2:]) // Remove 0x prefix
	if err != nil {
		return false
	}

	// Split into r and s
	r := new(big.Int).SetBytes(sigBytes[:len(sigBytes)/2])
	s := new(big.Int).SetBytes(sigBytes[len(sigBytes)/2:])

	return ecdsa.Verify(pubKey, hash[:], r, s)
}
