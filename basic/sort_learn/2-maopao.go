package sort_learn

import "fmt"

// O(n2) 可以解答上来
func MaoPao(nums []int) []int {
	//两两交换发生于每组遍历nums，遍历了len(nums)-1次
	for r := range len(nums) - 1 {
		fmt.Println(r)
		for i := 0; i < len(nums)-1; i++ {
			if nums[i+1] < nums[i] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}
	return nums
}
