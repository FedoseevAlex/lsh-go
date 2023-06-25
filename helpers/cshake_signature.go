package helpers

import "math/rand"

// Use seeded pseudo random number generator to generate signature for
// CSHAKE hash function family
func NewPseudoRandomFamilySignature(seed int64, rows, cols int) ([][]byte, error) {
	rng := rand.New(rand.NewSource(seed))
	signature := make([][]byte, 0, rows)
	for row := 0; row < rows; row++ {
		s := make([]byte, cols)
		_, err := rng.Read(s)
		if err != nil {
			return nil, err
		}
		signature = append(signature, s)
	}

	return signature, nil
}
