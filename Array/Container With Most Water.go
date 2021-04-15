package main

import "math"

/**

Area = max(min(height[i], height[j]) * (j - i))
max(min(height[0], height[1]) * (1 - 0))
*/
func maxArea(height []int) int {

	i := 0
	j := 1

	area := 0.0
	max := area
	for i < j {
		if i >= len(height) {
			break
		}

		if j >= len(height) {
			i++
			j = i + 1
			continue
		}
		area = math.Min(float64(height[i]), float64(height[j])) * float64(j-i)
		if area > max {
			max = area
		}

		j++
	}
	return int(max)
}

/**
最优解
基本的表达式: area = min(height[i], height[j]) * (j - i)
使用两个指针，值小的指针向内移动，这样就减小了搜索空间 因为面积取决于指针的距离与值小的值乘积，如果值大的值向内移动，距离一定减小,
而求面积的另外一个乘数一定小于等于值小的值，因此面积一定减小，而我们要求最大的面积，因此值大的指针不动，而值小的指针向内移动遍历
*/
func maxArea1(height []int) int {

	left, right := 0, len(height)-1
	max, area := 0, 0
	for left < right {
		if hl, hr, distance := height[left], height[right], right-left; hl > hr {
			area = hr * distance
			right--
		} else {
			area = hl * distance
			left++
		}

		if max < area {
			max = area
		}
	}
	return max
}

func main() {
	println(maxArea1([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
