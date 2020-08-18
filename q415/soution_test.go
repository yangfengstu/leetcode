package q415

import (
	"fmt"
	"testing"
)

func TestAddString(t *testing.T) {
	testCase := [][2]string{
		[2]string{"12345", "1234"},
	}
	for _, v := range testCase {
		testRes := addStrings(v[0], v[1])
		fmt.Println(testRes)
	}
}
