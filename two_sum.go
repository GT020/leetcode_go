package leetcode

func TwoSum(nums []int, target int) []int {
	mem := make(map[int]int)
	result := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		val, ok := mem[target-nums[i]]
		if ok {
			result[0] = val
			result[1] = i
			return result
		} else {
			mem[nums[i]] = i
		}
	}
	return result
}
