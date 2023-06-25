package helpers

// Given items = []byte("abcdefgh") and chunkSize = 3
// function will return [][]byte{{'a', 'b', 'c'}, {'d', 'e', 'f'}, {'g', 'h'}}
func Chunkify[T any](chunkSize int, items []T) [][]T {
	result := make([][]T, 0, len(items)/chunkSize+1)
	for len(items) > chunkSize {
		result = append(result, items[:chunkSize])
		items = items[chunkSize:]
	}
	result = append(result, items)
	return result
}

// Given items = []byte("abcd") and chunkSize = 3
// function will return [][]byte{{'a', 'b', 'c'}, {'b', 'c', 'd'}}
func Shinglets[T any](size int, items []T) [][]T {
	step := 1
	result := make([][]T, 0, len(items)-size+1)
	for len(items) > size {
		result = append(result, items[:size])
		items = items[step:]
	}
	result = append(result, items)
	return result
}
