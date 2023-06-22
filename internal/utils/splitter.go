package utils

import "strings"

func SplitByWhitespace(s string) [][]byte {
	parts := strings.Split(s, " ")

	result := make([][]byte, 0, len(parts))
	for _, part := range parts {
		result = append(result, []byte(part))
	}
	return result
}

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
func ChunkifyOverlap[T any](chunkSize int, items []T) [][]T {
	result := make([][]T, 0, len(items)-chunkSize+1)
	for len(items) > chunkSize {
		result = append(result, items[:chunkSize])
		items = items[1:]
	}
	result = append(result, items)
	return result
}
