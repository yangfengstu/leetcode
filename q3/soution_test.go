package q3

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	testCase := map[string]int{
		"qweerrtyy":  3,
		"djkjplioad": 8,
	}

	for k, v := range testCase {
		testRes := lengthOfLongestSubstring(k)
		if testRes != v {
			t.Error("case", k, "未通过。测试结果是：", testRes, "正确结果是", v)
		} else {
			t.Log("case", k, "通过。测试结果是：", testRes, "正确结果是", v)
		}
	}
}
