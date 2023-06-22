package jaccard

func JaccardSimilarity[T comparable](first, second []T) float32 {
	return float32(intersectionLen(first, second)) / float32(unionLen(first, second))
}

func unionLen[T comparable](first, second []T) int {
	all := make([]T, 0, len(first)+len(second))
	all = append(all, first...)
	all = append(all, second...)

	result := make(map[T]struct{}, len(first)+len(second))
	for _, elem := range all {
		if _, ok := result[elem]; !ok {
			result[elem] = struct{}{}
		}
	}

	return len(result)
}

func intersectionLen[T comparable](first, second []T) int {
	firstSet := make(map[T]struct{}, len(first))
	for _, elem := range first {
		firstSet[elem] = struct{}{}
	}

	secondSet := make(map[T]struct{}, len(first))
	for _, elem := range first {
		secondSet[elem] = struct{}{}
	}

	result := 0
	for k := range firstSet {
		if _, ok := secondSet[k]; ok {
			result++
		}
	}

	return result
}
