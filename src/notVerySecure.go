package main

import (
	"fmt"
	"regexp"
)

/*
func Test_alphanumeric(t *testing.T) {
	fmt.Println(alphanumeric("Mazinkaiser\\")) // [\]
}
*/

func alphanumeric(str string) bool {
	if len(str) == 0 {
		return false
	}

	reg := regexp.MustCompile(`\W`)
	if reg == nil {
		return false
	}

	result := reg.FindAllString(str, -1)


	if len(result) > 0 {
		return false
	}

	return true
}
