package sort_learn

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestSelectSort(t *testing.T) {
	nums := []int{
		3, 2, 4, 7, 8, 1, 99, 0, 4,
	}
	res := SelectSort(nums)
	fmt.Println(res)
	ok := reflect.DeepEqual(res, []int{0, 1, 2, 3, 4, 4, 7, 8, 99})
	assert.True(t, ok)
}

func TestMaoPao(t *testing.T) {
	nums := []int{
		3, 2, 4, 7, 8, 1, 99, 0, 4,
	}
	res := MaoPao(nums)
	fmt.Println(res)
	ok := reflect.DeepEqual(res, []int{0, 1, 2, 3, 4, 4, 7, 8, 99})
	fmt.Println(res[1:2])

	assert.True(t, ok)
}

func TestInsertion(t *testing.T) {
	nums := []int{
		3, 2, 4, 7, 8, 1, 99, 0, 4,
	}
	res := Insertion(nums)
	fmt.Println(res)
	ok := reflect.DeepEqual(res, []int{0, 1, 2, 3, 4, 4, 7, 8, 99})
	assert.True(t, ok)
}

func TestGuiBing(t *testing.T) {
	nums := []int{
		3, 2, 4, 7, 8, 1, 99, 99, 999,
	}
	res := quickSort(nums)
	fmt.Println(res)
	ok := reflect.DeepEqual(res, []int{0, 1, 2, 3, 4, 4, 7, 8, 99})
	assert.True(t, ok)
}
