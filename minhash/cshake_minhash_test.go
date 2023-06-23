package minhash

import (
	"testing"
)

func TestRandomSignature(t *testing.T) {
	t.Parallel()

	familySize := 5
	signaturePartLength := 10

	signature, err := GetCSHAKERandomSignature(familySize, signaturePartLength)
	if err != nil {
		t.Errorf("get cshake random signature: %v", err)
	}

	if len(signature) != familySize {
		t.Errorf("signature length mismatch: expected - %d, got - %d", familySize, len(signature))
	}

	for _, signaturePart := range signature {
		if len(signature) != familySize {
			t.Errorf("signature length part mismatch: expected - %d, got - %d", signaturePartLength, len(signaturePart))
		}
	}
}

func TestHashing(t *testing.T) {
	t.Parallel()

	familySize := 5
	signaturePartLength := 10
	signature, err := GetCSHAKERandomSignature(familySize, signaturePartLength)
	if err != nil {
		t.Errorf("get cshake random signature: %v", err)
	}

	hashFamily, err := NewCShakeHashFamily(signature, 10)
	if err != nil {
		t.Errorf("create new cshake family: %v", err)
	}

	h, err := hashFamily.MinHash([][]byte{[]byte("some random piece of data"), []byte("really random 100%")})
	if err != nil {
		t.Errorf("get hash from cshake family: %v", err)
	}

	if len(h) == 0 {
		t.Error("hash family generated an empty hash")
	}
}
