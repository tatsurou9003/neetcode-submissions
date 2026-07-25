import (
	"cmp"
	"slices"
)

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}

	keys := make([]int, 0, len(count))
	for key := range count {
		keys = append(keys, key)
	}

	slices.SortFunc(keys, func(a, b int) int {
		return cmp.Compare(count[b], count[a])
	})

	return keys[:k]
}
