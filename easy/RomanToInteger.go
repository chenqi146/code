package main

func romanToInt(s string) int {

	pre := s[0]
	num := 0
	switch pre {
	case 'I':
		num += 1
	case 'V':
		num += 5
	case 'X':
		num += 10
	case 'L':
		num += 50
	case 'C':
		num += 100
	case 'D':
		num += 500
	case 'M':
		num += 1000
	}
	for i := 1; i < len(s); i++ {

		switch s[i] {
		case 'I':
			num += 1
		case 'V':
			if pre == 'I' {
				num += 3
			} else {
				num += 5
			}
		case 'X':
			if pre == 'I' {
				num += 8
			} else {
				num += 10
			}
		case 'L':
			if pre == 'X' {
				num += 30
			} else {
				num += 50
			}
		case 'C':
			if pre == 'X' {
				num += 80
			} else {
				num += 100
			}
		case 'D':
			if pre == 'C' {
				num += 300
			} else {
				num += 500
			}
		case 'M':
			if pre == 'C' {
				num += 800
			} else {
				num += 1000
			}
		}
		pre = s[i]
	}
	return num
}

// 优化版
func romanToInt1(s string) int {
	pre := 0
	r := 0
	for i := len(s) - 1; i >= 0; i-- {
		cur := getRomanToInt(s[i])
		if cur >= pre {
			r += cur
		} else {
			r -= cur
		}
		pre = cur
	}

	return r
}

func getRomanToInt(r byte) int {
	switch r {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}

func main() {
	println(romanToInt1("MCMXCIV"))
}
