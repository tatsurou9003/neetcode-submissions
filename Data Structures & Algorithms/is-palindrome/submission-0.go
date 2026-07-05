func isPalindrome(s string) bool {
	newStr := ""
	for _, c := range s {
		if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
			newStr += string(c)
		} else if 'A' <= c && c <= 'Z' {
			newStr += string(c + 'a' - 'A')
		}
	}

	reversedStr := reverse(newStr)
	return newStr == reversedStr
}

func reverse(s string) string {
	bytes := []byte(s)
	n := len(bytes)
	for i := 0; i < n/2; i++ {
		bytes[i], bytes[n-1-i] = bytes[n-1-i], bytes[i]
	}

	return string(bytes)
}

/*
indexを前から読んでいく、後ろから読んでいくもので==, !=でfalse
*/
