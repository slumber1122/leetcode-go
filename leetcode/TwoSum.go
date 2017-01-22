package leetcode

func TwoSum(nums []int, target int) []int {
	var z []int
	num_dict := make(map[int]int)
	for idx, num := range nums {
		num_dict[num] = idx
	}
	for idx, num := range nums {
		key := target - num
		if val, ok := num_dict[key]; ok  && val != idx {
			z = append(z, idx, val)
			return z
		}
	}
	return z
}
