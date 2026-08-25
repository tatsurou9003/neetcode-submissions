func isPalindrome(s string) bool {
	left := 0
	right := len(s)-1

	for left < right {
		if !isalnum(s[left]) {
			left++
			continue
		}
		if !isalnum(s[right]) {
			right--
			continue
		}

	if toLower(s[left]) != toLower(s[right]) {
		return false
	}

		left++
		right--
	}
	return true
}

func isalnum(b byte) bool {
	return ('a' <= b && b <= 'z') || ('A' <= b && b <= 'Z') || ('0' <= b && b <= '9')
}

func toLower(b byte) byte {
	if 'A' <= b && b <= 'Z' {
		return b + ('a'-'A')
	}
	return b
}

/*
stringにindexでアクセスするとbyteになる
''はrune型(4バイト)
*/
