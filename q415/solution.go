package q415

import (
	"fmt"
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	a := '1'
	b := '2'
	fmt.Println(a + b - '0' - '0')
	flag := 0
	var res string
	l1 := len(num1) - 1
	l2 := len(num2) - 1
	for i, j := l1, l2; i >= 0 || j >= 0 || flag != 0; i, j = i-1, j-1 {

		numsi, err := getNumsByIndex(num1, i)
		if err != nil {
			return "0"
		}
		numsj, err := getNumsByIndex(num2, j)
		if err != nil {
			return "0"
		}
		tmpRes := numsi + numsj + flag
		if tmpRes >= 10 {
			res = strconv.Itoa(tmpRes-10) + res
			flag = 1
		} else {
			res = strconv.Itoa(tmpRes) + res
			flag = 0
		}
	}
	return string(res)
}
func getNumsByIndex(nums string, index int) (int, error) {
	if index < 0 {
		return 0, nil
	}
	num, err := strconv.Atoi(string(nums[index]))
	if err != nil {
		return 0, err
	}

	return num, nil
}
