package main

// 35. 搜索插入位置
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	i, j := 0, len(nums)-1

	if target < nums[i] {
		return 0
	}

	if target > nums[j] {
		return j + 1
	}

	mid := (i + j) / 2
	for {
		if nums[i] == target {
			return i
		}
		if nums[j] == target {
			return j
		}
		if nums[mid] == target {
			return mid
		}

		if j-i <= 1 {
			if nums[j] < target {
				return j + 1
			}

			if nums[i] > target {
				return i - 1
			}

			if target > nums[i] {
				return i + 1
			}
		}

		if nums[mid] < target {
			i = mid
		} else {
			j = mid
		}
		mid = (i + j) / 2
	}
}

/**
	[1, 3, 5, 6, 7] 2
	l	r	m
1.	0	5	2
2.	0	2	1
3.	0	1	0
*/
func searchInsert1(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func main() {
	println(searchInsert1([]int{1, 3, 5, 6, 7}, 2))

}
