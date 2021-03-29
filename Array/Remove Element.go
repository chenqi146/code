package main

import "fmt"

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	for j := len(nums) - 1; ; {
		if i > j {
			break
		}
		if nums[i] == val && nums[j] == val {
			j--
			continue
		}

		if nums[i] == val {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		}
		i++
	}
	return i
}

func main() {
	ints := []int{0, 1, 2, 2, 3, 0, 4, 2}
	println(removeElement(ints, 2))

	fmt.Printf("%#v", ints)
}
