package sort_learn

func quickSort(nums []int) []int {
	//选择基准值
	l := 0
	r := len(nums) - 1
	rank(nums, l, r)
	return nums
}

func partition(nums []int, l, r int) int {
	//核心实现
	i := l
	j := r
	baseI := i
	for i < j {
		for i < j && nums[j] >= nums[baseI] {
			j--
		}
		for i < j && nums[i] <= nums[baseI] {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	//此时i=j
	nums[baseI], nums[i] = nums[i], nums[baseI]
	return i
}

func rank(nums []int, l, r int) {
	if l >= r {
		return
	}
	mid := partition(nums, l, r)
	rank(nums, l, mid-1) //mid不再需要参与排序
	rank(nums, mid+1, r)
}
