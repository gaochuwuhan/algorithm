/**
输入一个单链表，链表中一个或多个连续的整数组成一个子链表。求所有子链表和的最大值。 
 1.你能不使用额外的O(n)空间复杂度来完成该题吗？ 2.你能使用递归解法来完成该题吗？ 3.如果你能使用递归解法完成本题，本题也可以看成二叉树中的最大路径和
的一个子问题，可以尝试递归解法该题之后，去突破二叉树中的最大路径和 
 数据范围：链表长度满足 ，链表中的值满足 
 
 
 Related Topics 动态规划 记忆化搜索 链表 
示例:
输入:{1,-2,3,10,-4,7,2,-5}
输出:18

*/

package nowcoder

//nowcoder submit region begin(Prohibit modification and deletion)
package main
//import "fmt"
//import "math"

import . "nc_tools"


/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 
 * @param head ListNode类 
 * @return int整型
*/

func Max(x,y int)int{
    if x>y{
        return x
    }
    return y
}
func FindGreatestSumOfSubArray( head *ListNode ) int {
    // write code here
    // 遍历时记录一个num[i]最大值，如果碰到累加后低于这个值，就不要前面的累计和
    maxNodeVal:=-99999
    maxSum:=-99999
    for head!=nil{
        maxNodeVal = Max(head.Val,head.Val+maxNodeVal)
        maxSum = Max(maxSum,maxNodeVal)
        head=head.Next
    }

    return maxSum

}
//nowcoder submit region end(Prohibit modification and deletion)
