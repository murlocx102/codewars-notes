package main

import (
	"math"
)

/*
func Test_ChooseBestSum(t *testing.T) {
	fmt.Println(ChooseBestSum(163, 3, []int{50, 55, 56, 57, 58})) // 163
}
*/

// ChooseBestSum
func ChooseBestSum(t, k int, ls []int) int {
	combination := make([][]int, 0, 0)
	optimal := -1

	// 穷举全组合
	var a = int(math.Pow(float64(2), float64(len(ls))) - 1)
	var num []int
	for i := 1; i <= a; i++ {
		s := i
		num = []int{}
		for k := 0; s > 0; k++ {
			if s&1 == 1 {
				num = append(num, ls[k])
			}

			s >>= 1
		}
		combination = append(combination, num)
	}

	// 遍历所有组合(去掉不符合数量的).获取最大值
	for _, value := range combination {
		num := 0
		if len(value) < k || len(value) > k {
			continue
		}

		for _, v := range value {
			num += v
		}

		if num <= t {
			if num > optimal {
				optimal = num
			}
		}
	}

	return optimal
}
