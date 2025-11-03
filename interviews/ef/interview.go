package ef

func longestNonrepeatedSubstrings(s string) []string {
	if len(s) == 0 {
		return nil
	}

	left, right := 0, 0
	charCnt := make(map[byte]int)
	charCnt[s[0]] = 1

	maxLen := 1
	ans := []string{string(s[0])}

	var repeatedChar byte

	right++
	for right < len(s) {
		for right < len(s) && charCnt[s[right]] == 0 {
			charCnt[s[right]]++
			right++
		}

		if right < len(s) {
			repeatedChar = s[right]
		}

		newLen := right - left
		if newLen > maxLen {
			maxLen = newLen
			ans = []string{s[left:right]}
		} else if newLen == maxLen {
			ans = append(ans, s[left:right])
		}

		if right == len(s) {
			break
		}

		for left < right && s[left] != repeatedChar {
			charCnt[s[left]]--
			left++
		}

		// s[left] == repeatedChar
		charCnt[s[left]]--
		left++
	}

	return ans
}
