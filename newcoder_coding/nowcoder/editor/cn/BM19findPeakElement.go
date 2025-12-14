/**
给定一个长度为n的数组nums，请你找到峰值并返回其索引。数组可能包含多个峰值，在这种情况下，返回任何一个所在位置即可。 1.峰值元素是指其值严格大于左右相邻值
的元素。严格大于即不能有等于 2.假设 nums[-1] = nums[n] = 3.对于所有有效的 i 都有 nums[i] != nums[i + 1] 4
.你可以使用O(logN)的时间复杂度实现此问题吗？ 
 数据范围： 
 
 
 如输入[2,4,1,2,7,8,4]时，会形成两个山峰，一个是索引为1，峰值为4的山峰，另一个是索引为5，峰值为8的山峰，如下图所示： 
 Related Topics 数组 查找 
示例:
输入:[2,4,1,2,7,8,4]
输出:1 

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
 * @return int整型
*/
func findPeakElement( nums []int ) int {
    // write code here
    l:=0
    r:=len(nums)-1
    for l<r{
        mid:=(l+r)/2
        if nums[mid]<nums[mid+1]{
            l=mid+1
        }else{
            r=mid //mid本身可能是峰值
        }
    }
    return l
}
//nowcoder submit region end(Prohibit modification and deletion)
