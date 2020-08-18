package q415

func palindromePairs(words []string) [][]int {
	var res [][]int
	res = [][]int{}
	reverseStr := []string{}
	reverseStrMap := map[string]int{}
	for _, v := range words {
		reverseStr = append(reverseStr, reverse(v))
	}
	for k, v := range reverseStr {
		reverseStrMap[v] = k
	}
	for k1, word1 := range words {
		len1 := len(word1)
		for k2, words2 := range words {

		}
	}
	return res
}

func reverse(s string) string {
	len := len(s)
	reverseStr := []byte(s)
	for i := 0; i < (len-1)/2; i++ {
		reverseStr[i], reverseStr[len-i-1] = reverseStr[len-i-1], reverseStr[i]
	}
	return string(reverseStr)
}
func isPalindrome(s string, left, right int) bool {
	for i := 0; i < (right-left-1)/2; i++ {
		if s[left+1] != s[right-1] {
			return false
		}
	}
	return true
}
