package minhash

import (
	"bytes"
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/sha3"
)

// CHAKE based family of hash functions
type CShakeHashFamily struct {
	// Output lentgh that will be taken from each function in family
	HashOutputLength int
	hashes           []sha3.ShakeHash
}

// Interface implementation guard
var _ MinHasher = (*CShakeHashFamily)(nil)

func (hf *CShakeHashFamily) MinHash(dataParts [][]byte) ([]byte, error) {
	result := make([]byte, 0, len(hf.hashes)*hf.HashOutputLength)
	for _, hash := range hf.hashes {
		minHash := make([]byte, 0, hf.HashOutputLength)
		for _, part := range dataParts {
			_, err := hash.Write(part)
			if err != nil {
				return nil, fmt.Errorf("write to cshake hash function: %w", err)
			}

			h := make([]byte, hf.HashOutputLength)
			_, err = hash.Read(h)
			if err != nil {
				return nil, fmt.Errorf("read from cshake hash function: %w", err)
			}
			hash.Reset()

			if len(minHash) == 0 {
				minHash = h
				continue
			}

			if bytes.Compare(h, minHash) == -1 {
				minHash = h
			}
		}
		result = append(result, minHash...)
	}
	return result, nil
}

func NewCShakeHashFamily(signature [][]byte, hashOutputLength int) (*CShakeHashFamily, error) {
	// Input validation
	if hashOutputLength <= 0 {
		return nil, fmt.Errorf("hash output length must be greater than zero")
	}
	if len(signature) == 0 {
		return nil, fmt.Errorf("hash signature size must be greater than zero")
	}
	for _, signaturePart := range signature {
		if len(signaturePart) == 0 {
			return nil, fmt.Errorf("all signature parts must be greater than zero")
		}
	}

	// Create hash family
	hashes := make([]sha3.ShakeHash, 0, len(signature))
	for _, signaturePart := range signature {
		// Here I intentionally pass signature parts to N argument
		// because we need a hash function family and not a
		// one function with separate domains.
		hashes = append(hashes, sha3.NewCShake128(signaturePart, nil))
	}

	return &CShakeHashFamily{hashes: hashes, HashOutputLength: hashOutputLength}, nil
}

// Just get random signature if you don't have one or don't care.
func GetCSHAKERandomSignature(size, length int) ([][]byte, error) {
	signature := make([][]byte, 0, size)
	for i := 0; i < size; i++ {
		signaturePart := make([]byte, length)
		_, err := rand.Read(signaturePart)
		if err != nil {
			return nil, err
		}
		signature = append(signature, signaturePart)
	}
	return signature, nil
}
