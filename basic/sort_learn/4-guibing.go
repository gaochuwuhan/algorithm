package sort_learn

// O(log2N) 递归没写上来，打算参数传left和right的切片，其实只要传递索引l、r就行
func GuiBing(nums []int) []int {
	//遍历了len(nums)-1次
	split(nums, 0, len(nums)-1)
	return nums
}

func merge(nums []int, l, mid, r int) {

	//使用临时数组合并两个排序后的结果
	tmp := make([]int, 0)
	curL := l
	curR := mid + 1
	for curL <= mid && curR <= r {
		if nums[curR] < nums[curL] {
			tmp = append(tmp, nums[curR])
			curR++
		} else {
			tmp = append(tmp, nums[curL])
			curL++
		}

	}
	// 结束之后curL和curR都是还没判断过的位置，有剩余的直接放到tmp后面
	if curL <= mid {
		for _, v := range nums[curL : mid+1] {
			tmp = append(tmp, v)
		}
	}
	if curR <= r {
		for _, v := range nums[curR : r+1] {
			tmp = append(tmp, v)
		}
	}
	//将新的tmp拷贝会原来的num位置
	copy(nums[l:r+1], tmp)
}

func split(nums []int, l, r int) {
	if l >= r {
		return
	}
	mid := (l + r) / 2
	split(nums, l, mid)
	split(nums, mid+1, r)
	merge(nums, l, mid, r)
}
