package main

import (
	"fmt"
	"strconv"
	"testing"
)

/*
func Test_num(t *testing.T) {

	n := 132189
	fmt.Println(DigitalRoot(n))
}
*/

// 数字根
func DigitalRoot(n int) int {
	var addNum func(n int) int

	addNum = func(n int) int {
		def := 0

		for _, v := range []rune(strconv.Itoa(n)) {
			r, _ := strconv.Atoi(string(v))
			def += r
		}

		if def < 10 {
			return def
		}

		return addNum(def)
	}

	return addNum(n)
}
