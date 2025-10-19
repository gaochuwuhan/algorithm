package basic

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func TestSlice(t *testing.T) {
	x := []int{1}
	x = x[:len(x)-1]
	t.Log(x)
}

var (
	leftArr   = []string{}
	leftSets  = []string{"(", "{", "["}
	rightSets = []string{")", "}", "]"}
	matchRule = map[string]string{
		"}": "{",
		")": "(",
		"]": "[",
	}
)

func IsValid(s string) bool {
	//设置一个left栈，仅存储左括号
	if len(s) == 1 {
		return false
	}
	for _, runeS := range s {
		currentS := string(runeS)
		if slices.Contains(leftSets, currentS) {
			leftArr = append(leftArr, currentS)
			continue
		}
		//右括号检查匹配
		matched := matchAndRmLastLeft(currentS)
		if !matched {
			return false
		}

	}
	if len(leftArr) > 0 {
		return false
	}
	return true
}

func matchAndRmLastLeft(right string) bool {
	//从leftArr中找到匹配的左括号，则返回true,并移除该左括号
	if len(leftArr) == 0 {
		return false
	}
	lastLeft := leftArr[len(leftArr)-1]
	if matchRule[right] == lastLeft {
		//移除
		leftArr = leftArr[:len(leftArr)-1]
		return true
	}
	return false
}

func TestIsValid(t *testing.T) {
	input := "([])"
	res := IsValid(input)
	assert.True(t, res)
}
