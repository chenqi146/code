package main

import "math"

/**
	https://leetcode-cn.com/problems/median-of-two-sorted-arrays/

	[1, 2, 3, 4]     9  5 2
	[1, 6, 8, 9, 11]  1 1 2 3 4 6 8 9 11
[1], 9, 10, 11, 12]   k = 4 , k = 2 , k=1
[3, 4,] 6, 7]
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	if length == 0 {
		return 0
	}

	r := make([]int, length, length)
	i, j, z := 0, 0, 0
	for ; i < len(nums1) && j < len(nums2); z++ {
		if nums1[i] <= nums2[j] {
			r[z] = nums1[i]
			i++
		} else {
			r[z] = nums2[j]
			j++
		}
	}

	if i >= len(nums1) {
		for ; j < len(nums2); j, z = j+1, z+1 {
			r[z] = nums2[j]
		}
	} else {
		for ; i < len(nums1); i, z = i+1, z+1 {
			r[z] = nums1[i]
		}
	}

	i2 := (len(r) - 1) / 2
	if len(r)%2 == 1 {
		return float64(r[i2])
	} else {
		return float64(r[i2]+r[i2+1]) / 2
	}
}

/***
		  L1   R1
num1  1   3  | 5    10                  4
num2  1   2   4  |  7    9   14         6
		     L2     R2

要满足  L1 <= R2     L2 <= R1   时找到中位数

切割线num1    cut1  左边有几个元素
切割线num2    cut2  左边有几个元素

当  L1 > R2 时   切割线左移   (cutL, cut1 - 1)
当  L2 > R1 时   切割线右移   (cut1 + 1, cutR)

*/
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	// 一直在最小的数组切分  达到最大效率
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays1(nums2, nums1)
	}

	length := len(nums1) + len(nums2)
	cut1 := 0
	cut2 := 0
	// 二分的两个范围值
	cutL := 0
	cutR := len(nums1)

	for cut1 <= len(nums1) {
		cut1 = (cutR-cutL)/2 + cutL
		cut2 = length/2 - cut1
		// 默认给最小值  如果没有划分线  那就是左边不需要划分
		L1 := math.MinInt64
		if cut1 != 0 {
			L1 = nums1[cut1-1]
		}
		L2 := math.MinInt64
		if cut2 != 0 {
			L2 = nums2[cut2-1]
		}

		// 默认给最大值  如果没有划分线  那就是右边不需要划分
		R1 := math.MaxInt64
		if cut1 != len(nums1) {
			R1 = nums1[cut1]
		}
		R2 := math.MaxInt64
		if cut2 != len(nums2) {
			R2 = nums2[cut2]
		}

		if L1 > R2 {
			cutR = cut1 - 1
		} else if L2 > R1 {
			cutL = cut1 + 1
		} else {
			// 偶数
			if length%2 == 0 {
				return (math.Max(float64(L1), float64(L2)) + math.Min(float64(R1), float64(R2))) / 2
			} else {
				return math.Min(float64(R1), float64(R2))
			}
		}
	}
	return -1
}

func main() {
	println(findMedianSortedArrays1([]int{1, 3}, []int{2}))
}
