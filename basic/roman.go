package basic

import "slices"

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

func RomanToInt(s string) int {
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

func longestCommonPrefix(strs []string) string {
	commPrefix := ""
	firstS := strs[0]
	//firstSlen:=len(firstS)
	hasRuneF := true
	for fi, runeF := range firstS {
		//遍历其余字符串的字符
		if !hasRuneF {
			break
		}
		for i := 1; i < len(strs); i++ {
			strI := strs[i]
			if fi <= len(strI)-1 {
				if string(strI[fi]) != string(runeF) {
					hasRuneF = false
					break
				}
			} else {
				hasRuneF = false
				break
			}
			//判断是最后一个字符串，说明都有前缀
			if i == len(strs)-1 {
				commPrefix = commPrefix + string(runeF)
			}
		}
	}
	return commPrefix
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	// write code here
	if root == nil {
		return false
	}
	curSum := 0
	has := false
	walk(root, &curSum, sum, &has)
	return has
}

func walk(tree *TreeNode, curSum *int, sum int, has *bool) {
	if *has {
		return
	}
	if tree.Left == nil && tree.Right == nil {
		if *curSum == sum {
			*has = true
		}
		return
	}
	if *curSum+tree.Val > sum {
		return
	}
	*curSum = *curSum + tree.Val
	if tree.Left != nil {
		walk(tree.Left, curSum, sum, has)
	}
	walk(tree.Right, curSum, sum, has)
}
