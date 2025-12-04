/**
给定一个单链表的头结点pHead(该头节点是有值的，比如在下图，它的val是1)，长度为n，反转该链表后，返回新链表的表头。 
 数据范围： 要求：空间复杂度 ，时间复杂度 。 
 如当输入链表{1,2,3}时， 经反转后，原链表变为{3,2,1}，所以对应的输出为{3,2,1}。 以上转换过程如下图所示： 
 Related Topics 链表 
示例:
输入:{1,2,3}
输出:{3,2,1}

*/
package nowcoder.editor.cn;  //根据实际修改
//nowcoder submit region begin(Prohibit modification and deletion)
package main
//import "fmt"
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
 * @return ListNode类
*/
func ReverseList( head *ListNode ) *ListNode {
    // write code here
    if head == nil{
        return nil
    }
    var pre *ListNode
    cur:=head
    for cur!=nil{
        next:=cur.Next
        cur.Next=pre
        pre = cur
        cur=next
    }
    return pre


}
//nowcoder submit region end(Prohibit modification and deletion)
