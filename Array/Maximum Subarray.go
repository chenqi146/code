package main

/**
给定一个整数数组 nums，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。



示例 1：

输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组[4,-1,2,1] 的和最大，为6 。
示例 2：

输入：nums = [1]
输出：1
示例 3：

输入：nums = [0]
输出：0
示例 4：

输入：nums = [-1]
输出：-1
示例 5：

输入：nums = [-100000]
输出：-100000


提示：

1 <= nums.length <= 3 * 104
-105 <= nums[i] <= 105
f(n)为到第n位结束的子数组的最大和
	f(n) = max(f(n - 1), nums[n])
	f(1) = nums[1]
*/

func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {

		// 如果当前位加上之前的比当前要大  继续累加 否则就丢弃 继续移动
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}

		// 如果当前位比最大值大则替换
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func main() {

}
