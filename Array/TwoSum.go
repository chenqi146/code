package Array


/**
 暴力解法
 */
func twoSum(nums []int, target int) []int {

	length := len(nums)
	for i := 0; i < length- 1; i++ {
		for j:= i + 1; j < length; j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum1(nums []int, target int) []int {
	saveMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		temp := target - nums[i]
		if _, ok := saveMap[temp]; ok {
			return []int{saveMap[temp], i}
		}
		saveMap[nums[i]] = i
	}
	return nil
}