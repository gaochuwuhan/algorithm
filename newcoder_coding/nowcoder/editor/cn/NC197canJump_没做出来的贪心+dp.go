/**
给定一个非负整数数组nums，假定最开始处于下标为0的位置，数组里面的每个元素代表下一跳能够跳跃的最大长度。如果能够跳到数组最后一个位置，则返回true，否则返
回false。 数据范围: 
 
 
 Related Topics 动态规划 
示例:
输入:[2,1,3,3,0,0,100]
输出:true 

*/
package nowcoder.editor.cn;  //根据实际修改
//nowcoder submit region begin(Prohibit modification and deletion)
package main
//import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 
 * @param nums int整型一维数组 
 * @return bool布尔型
*/
//
func canJump( nums []int ) bool {
    // write code here
	//能否跳刀最后一个位置的关键是当前index的值 num[i]>=len(num)-1-i

	if len(nums) == 1{
		return true
	}
	longestIdx:=0
	end:=len(nums)-1
	for i,v:=range nums{
		 if i>longestIdx{
			 return false
		 }
		 if i+v>longestIdx{
			 longestIdx=i+v
		 }
		if longestIdx>=end{
			return true
		}
	}
	return false

}


//nowcoder submit region end(Prohibit modification and deletion)
