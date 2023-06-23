package minhash

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"log"

	"golang.org/x/crypto/sha3"
)

const (
	cshakeCustomisationSize = 10
	defaultHashOutputLength = 20
)

type HashGetter interface {
	Hash(dataParts [][]byte) ([]byte, error)
}

type CShakeHashFamily struct {
	hashOutputLength int
	hashes           []sha3.ShakeHash
}

func (hf *CShakeHashFamily) Hash(dataParts [][]byte) ([]byte, error) {
	result := make([]byte, 0, len(hf.hashes)*hf.hashOutputLength)
	for _, hash := range hf.hashes {
		minHash := make([]byte, 0, hf.hashOutputLength)
		for _, part := range dataParts {
			_, err := hash.Write(part)
			if err != nil {
				return nil, fmt.Errorf("write to cshake hash function: %w", err)
			}

			h := make([]byte, hf.hashOutputLength)
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

func NewCShakeFamilyHashGetter(size, hashOutputLength int) (HashGetter, error) {
	if size <= 0 {
		return nil, fmt.Errorf("hash family size must be greater tnan zero")
	}

	if hashOutputLength <= 0 {
		return nil, fmt.Errorf("hash output length must be greater tnan zero")
	}

	hashes := make([]sha3.ShakeHash, 0, size)
	for i := 0; i < size; i++ {
		cshakeCustomization := make([]byte, cshakeCustomisationSize)
		_, err := rand.Read(cshakeCustomization)
		if err != nil {
			log.Fatal(err.Error())
		}

		hashes = append(hashes, sha3.NewCShake128(cshakeCustomization, []byte{}))
	}

	return &CShakeHashFamily{hashes: hashes, hashOutputLength: hashOutputLength}, nil
}
