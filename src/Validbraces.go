package main

/*
func Test_ValidBraces(t *testing.T) {

	str := "()[{()}]{}()"
	fmt.Println(ValidBraces(str))
}
*/

// 待优化.可使用堆栈方法或者匹配删除方式
func ValidBraces(str string) bool {
	data := []rune(str)
	result := make([]string, 0, 0)

	exp := []string{"(", "[", "{", ")", "]", "}"}
	for key := range data {
		if string(data[key]) == exp[0] || string(data[key]) == exp[1] || string(data[key]) == exp[2] {
			result = append(result, string(data[key]))
			continue
		}

		if len(result) == 0 {
			return false
		}

		switch {
		case string(data[key]) == exp[3]:
			left := len(result) - 1
			if result[left] == exp[0] {
				result = result[:len(result)-1]
			} else {
				return false
			}
		case string(data[key]) == exp[4]:
			left := len(result) - 1
			if result[left] == exp[1] {
				result = result[:len(result)-1]
			} else {
				return false
			}
		case string(data[key]) == exp[5]:
			left := len(result) - 1
			if result[left] == exp[2] {
				result = result[:len(result)-1]
			} else {
				return false
			}
		default:
			return false
		}
	}

	if len(result) == 0 {
		return true
	}

	return false
}
