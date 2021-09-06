package main

import (
	"strings"
)

/*
func Test_HighestScoringWord(t *testing.T) {

	str := "rjdpeqcnow chcilz kachyfpm"
	fmt.Println(HighestScoringWord(str))
}
*/

// 随机单词 第一个数值最大的单词
func HighestScoringWord(s string) string {
	wordSlice := strings.Split(s, " ")
	expMap := map[string]int{
		"a": 1, "b": 2, "c": 3, "d": 4, "e": 5,
		"f": 6, "g": 7, "h": 8, "i": 9, "j": 10,
		"k": 11, "l": 12, "m": 13, "n": 14, "o": 15,
		"p": 16, "q": 17, "r": 18, "s": 19, "t": 20,
		"u": 21, "v": 22, "w": 23, "x": 24, "y": 25,
		"z": 26,
	}

	res := make([]int, 0, 0)

	for _, value := range wordSlice {
		num := 0
		for _, v := range value {
			num += expMap[string(v)]
		}
		res = append(res, num)
	}

	if len(res) == 1 {
		return wordSlice[0]
	}

	var strLoop func(data []int) []int
	strLoop = func(data []int) []int {
		if len(data) > 1 {
			switch {
			case data[0] < data[len(data)-1]:
				data = data[1:]

			case data[0] > data[len(data)-1]:
				data = data[:len(data)-1]

			case data[0] == data[len(data)-1]:
				if len(data) == 2 {
					return (data[:1])
				} else {
					data = data[:len(data)-1]
				}
			}

			return strLoop(data)
		}

		return data
	}

	loopResult := strLoop(res)

	for key, val := range res {
		if val == loopResult[0] {
			return wordSlice[key]
		}
	}

	return wordSlice[0]
}
