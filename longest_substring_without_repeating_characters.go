package leetcode

func LengthOfLongestSubstring(s string) int {

	if len(s) == 1 {
		return 1
	}

	var maxlen int = 0
	var j int = 0
	var k int = 0

	for ; k <= len(s); k++ {
		sub := s[j:k]

		if checkNonRepeating(sub) {
			maxlen = len(sub)
		} else {
			j++
		}
	}

	return maxlen
}

func checkNonRepeating(s string) bool {
	m := make(map[rune]bool)

	for _, r := range s {
		_, ok := m[r]
		if !ok {
			m[r] = true
		} else {
			return false
		}
	}
	return true
}
