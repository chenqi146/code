package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	temp := 0
	for x != 0 {
		// 每次乘10且加上取余数
		temp = temp*10 + x%10
		x = x / 10
		// 判断是否溢出
		if temp < math.MinInt32 || temp > math.MaxInt32 {
			return 0
		}
	}
	return temp
}

func main() {
	fmt.Print(reverse(123))
}
