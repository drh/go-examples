package mapkeys

import (
	"cmp"
	"slices"
)

func SortByKey[M ~map[K]V, K cmp.Ordered, V any](m M) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	slices.SortStableFunc(keys, func(a, b K) int { return cmp.Compare(a, b) })
	return keys
}

func SortByValue[M ~map[K]V, K cmp.Ordered, V cmp.Ordered](m M) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	slices.SortStableFunc(keys, func(a, b K) int { return cmp.Compare(m[b], m[a]) })
	return keys
}
