// IMPORTANT!! Submit Code Region Begin(Do not remove this line)
package main

var charNumMap = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

var composeCharList = []rune{
	'I', 'X', 'C',
}

var composeMatch = map[rune][]rune{
	'I': {'V', 'X'},
	'X': {'L', 'C'},
	'C': {'D', 'M'},
}

func checkCompose(romanChar rune, tempRoman *[]rune) (int, bool) {
	if len(*tempRoman) == 0 {
		*tempRoman = append(*tempRoman, romanChar)
		return 0, false
	}
	lengL := len(*tempRoman) - 1
	if (*tempRoman)[lengL] == 'I' || (*tempRoman)[lengL] == 'X' || (*tempRoman)[lengL] == 'C' {
		//查看是否可以搭配
		if slices.Contains(composeMatch[(*tempRoman)[lengL]], romanChar) {
			//可以搭配则返回组合后的结果并清空temp
			sum := charNumMap[romanChar] - charNumMap[(*tempRoman)[lengL]]
			*tempRoman = []rune{}
			return sum, true
		} else {
			sum := charNumMap[(*tempRoman)[lengL]]
			*tempRoman = append(*tempRoman, romanChar)
			//说明上一个是可以相加的返回上一个数对应的值
			return sum, false
		}
	} else {
		sum := charNumMap[(*tempRoman)[lengL]]
		*tempRoman = append(*tempRoman, romanChar)
		return sum, false
	}
}

func romanToInt(s string) int {
	tempRomanList := []rune{}
	sum := 0
	lastId := len(s) - 1
	for i, romanChar := range s {
		partSum, composed := checkCompose(romanChar, &tempRomanList)
		sum = sum + partSum
		//最后一个
		if lastId == i && !composed {
			sum = charNumMap[romanChar] + sum
		}
	}

	return sum

}

//IMPORTANT!! Submit Code Region End(Do not remove this line)
