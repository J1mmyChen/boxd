package crypto

import (
	"crypto/sha256"

	"golang.org/x/crypto/ripemd160"
)

// Ripemd160 calculates the RIPEMD160 digest of buf
func Ripemd160(buf []byte) []byte {
	hasher := ripemd160.New()
	hasher.Write(buf)
	return hasher.Sum(nil)
}

// Sha256 calculates the sha256 digest of buf
func Sha256(buf []byte) []byte {
	digest := sha256.Sum256(buf)
	return digest[:]
}
