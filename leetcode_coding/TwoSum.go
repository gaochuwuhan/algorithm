
//IMPORTANT!! Submit Code Region Begin(Do not remove this line)
func twoSum(nums []int, target int) []int {
    // 时间复杂O(n2)
	for i,val := range nums {
		for j:=i+1;j<len(nums);j++{
			if val+nums[j]==target{
				return []int{i,j}
			}
		}
	}
	return []int{0,0}
}
//IMPORTANT!! Submit Code Region End(Do not remove this line)