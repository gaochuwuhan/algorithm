package sort_learn

// O(n2) 没有理解
func Insertion(nums []int) []int {
	//遍历了len(nums)-1次
	for i := 1; i < len(nums); i++ {
		j := i - 1 //从当前元素向挨个向前看
		cur := nums[i]
		for j >= 0 && cur < nums[j] {
			nums[j+1] = nums[j] //看到比自己大的就向右移动
			//向左移动
			j = j - 1
		}
		//j是上一个发现>cur的位置，在for循环里-1了所以cur要插入j+1位置（没有理解）
		nums[j+1] = cur

	}
	return nums
}
