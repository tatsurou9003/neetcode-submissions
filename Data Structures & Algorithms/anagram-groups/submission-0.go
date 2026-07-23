import (
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	res := make(map[string][]string)

	for _, s := range strs {
		runeS := []rune(s)
		slices.Sort(runeS)
		sortedS := string(runeS)
		res[sortedS] = append(res[sortedS], s)
	}

	var result [][]string
	for _, group := range res {
		result = append(result, group)
	}
	return result
}
