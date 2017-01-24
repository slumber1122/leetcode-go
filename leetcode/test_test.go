package leetcode

import "testing"

func TestTwoSum(t *testing.T) {
	arrs := [4]int{2, 7, 11, 15}
	res := twoSum(arrs[:], 9)
	if res[0] != 0 && res[1] != 1 {
		t.Error(`error`)
	}
}
