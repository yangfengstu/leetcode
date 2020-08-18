package q3

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
func lengthOfLongestSubstring(s string) int {
	var sSlice []rune
	sSlice = []rune(s)
	if len(sSlice) <= 0 {
		return 0
	}
	max := 1
	for i, char := range sSlice {
		charMap := map[rune]int{
			char: i,
		}
		for j := i + 1; j < len(sSlice); j++ {
			if _, ok := charMap[sSlice[j]]; ok {
				if max < j-i {
					max = j - i
				}
				break
			} else {
				charMap[sSlice[j]] = j
			}
			if max < j-i+1 {
				max = j - i + 1
			}
		}
	}
	return max
}

func lengthOfLongestSubstringSecond(s string) int {
	return 0
}
