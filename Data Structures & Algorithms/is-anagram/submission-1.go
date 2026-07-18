func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	countS, countT := make(map[rune]int), make(map[rune]int)
	runeS, runeT := []rune(s), []rune(t)

	for i, ch := range runeS {
		countS[ch]++
		countT[runeT[i]]++
	}

	for k, v := range countS {
		if countT[k] != v {
			return false
		}
	}
	return true
}
