package main

/*
func Test_HighestScoringWord(t *testing.T) {

	str := "iiisdoso" // 0+1+1+1  9  8 [8] 64 [8,64]
	fmt.Println(ParseDeadfish(str))
}
*/

// ParseDeadfish
func ParseDeadfish(data string) []int {
	num := 0
	result := make([]int, 0, 0)

	for _, value := range data {
		switch string(value) {
		case "i":
			num += 1
		case "d":
			num -= 1
		case "s":
			num *= num
		case "o":
			result = append(result, num)
		default:
			continue
		}
	}

	return result
}
