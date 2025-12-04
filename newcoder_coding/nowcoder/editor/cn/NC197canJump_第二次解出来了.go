/**
给定一个非负整数数组nums，假定最开始处于下标为0的位置，数组里面的每个元素代表下一跳能够跳跃的最大长度。如果能够跳到数组最后一个位置，则返回true，否则返
回false。 数据范围:


 Related Topics 动态规划
示例:
输入:[2,1,3,3,0,0,100]
输出:true

*/
//nowcoder submit region begin(Prohibit modification and deletion)
package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param nums int整型一维数组
 * @return bool布尔型
 */
func canJump(nums []int) bool {
	// write code here
	// 遍历，看最远能跳的地方reach_i
	//在往后遍历看当前i是否在reach_i范围内，超出范围说明当前数不能继续遍历了return false。未超出则继续更新reach_i
	if len(nums) == 1 { //由于题干说每个元素>=0,所以开始就到了终点
		return true
	}
	reach_i := nums[0]

	for i := 1; i < len(nums); i++ {
		if reach_i >= len(nums)-1 {
			return true
		}
		if i > reach_i { //比如i=1,reach_i=0
			return false
		}
		//继续更新reach_i,以较远的为reach_i
		reach_i = max(reach_i, i+nums[i])

	}

	return false

}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
