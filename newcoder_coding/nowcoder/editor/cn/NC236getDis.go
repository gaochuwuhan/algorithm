/**
有一个长为 n 的数组 A ，求满足 0 ≤ a ≤ b < n 的 A[b] - A[a] 的最大值。 给定数组 A 及它的大小 n ，请返回最大差值。 
 数据范围： ，数组中的值满足 
 Related Topics 贪心 动态规划 模拟 
示例:
输入:[5,1],2
输出:0 

*/
package nowcoder.editor.cn;  //根据实际修改
//nowcoder submit region begin(Prohibit modification and deletion)
package main
//import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 
 * @param A int整型一维数组 
 * @param n int整型 
 * @return int整型
*/
// 这道题最好动手画个图，去web上查看完整题干，长度多一些 会发现每次遍历到新元素时，其实就是在找他之前几个元素的最小值
//写的程序需要把当前值和之前的最小值进行比较，再把max差值记录即可
func Min(x,y int) int{
	if x<y{
		return x
	}
	return y
}
func Max(x,y int)int{
	if x>y{
		return x
	}
	return y
}
func getDis( A []int ,  n int ) int {
    // write code here

	maxSubVal:=0
	minVal:=A[0]
	for i:=1;i<n;i++{
		maxSubVal = Max(maxSubVal,A[i]-minVal)
		minVal = Min(A[i],minVal)
	}
	return maxSubVal
}
//nowcoder submit region end(Prohibit modification and deletion)
