import (
	"slices"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sRunes, tRunes := []rune(s), []rune(t)

	slices.Sort(sRunes)
	slices.Sort(tRunes)

	for i := range sRunes {
		if sRunes[i] != tRunes[i] {
			return false
		}
	}
	return true
}
