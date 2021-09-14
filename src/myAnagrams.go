package main

import (
	"strings"
)

/*
func Test_Anagrams(t *testing.T) {
	fmt.Println(Anagrams("aabb", []string{"aabb", "aabb", "aabb", "bbaa", "bbaa", "bbaa"})) // ["aabb","bbaa"]
}
*/

func Anagrams(word string, words []string) []string {
	tempMap := make(map[string]int, 0)
	result := make([]string, 0, 0)
	temp := ""

	for key, value := range words {
		if len(value) != len(word) {
			continue
		}
		temp = value
		for _, val := range word {
			temp = strings.Replace(temp, string(val), "", -1)

			if len(temp) == 0 {
				tempMap[words[key]] += 1
			}
		}
	}

	for k := range tempMap {
		result = append(result, k)
	}

	if len(result) == 0 {
		return nil
	}

	return result
}
