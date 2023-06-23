package minhash

type MinHasher interface {
	MinHash(dataParts [][]byte) ([]byte, error)
}
