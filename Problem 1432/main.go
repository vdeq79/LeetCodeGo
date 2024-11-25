package main

import (
	"fmt"
	"math"
)

func maxDiff(num int) int {

	var a []int
	var b []int

	for num > 0 {
		a = append(a, num%10)
		b = append(b, num%10)
		num /= 10
	}

	x_a := -1
	x_b := -1
	result_a := 0

	for i := len(a) - 1; i >= 0; i-- {
		if x_a == -1 && a[i] < 9 {
			x_a = a[i]
		}

		if x_a != -1 && a[i] == x_a {
			a[i] = 9
		}

		result_a += a[i] * int(math.Pow(10, (float64(i))))
	}

	s_b := 0
	if b[len(b)-1] > 1 {
		x_b = b[len(b)-1]
		b[len(b)-1] = 1
		s_b = 1
	}

	result_b := b[len(b)-1] * int(math.Pow(10, (float64(len(b)-1))))

	for i := len(b) - 2; i >= 0; i-- {
		if x_b == -1 && b[i] > 1 {
			x_b = b[i]
		}

		if x_b != -1 && b[i] == x_b {
			b[i] = s_b
		}

		result_b += b[i] * int(math.Pow(10, (float64(i))))
	}

	return result_a - result_b
}

func main() {
	fmt.Println(maxDiff(111))
}
