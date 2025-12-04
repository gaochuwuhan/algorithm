/**
将给出的链表中的节点每 k 个一组翻转，返回翻转后的链表
 如果链表中的节点数不是 k 的倍数，将最后剩下的节点保持原样
 你不能更改节点中的值，只能更改节点本身。 
 数据范围： ， ，链表中每个元素都满足 
 要求空间复杂度 ，时间复杂度 例如： 给定的链表是 对于 , 你应该返回 对于 , 你应该返回 
 
 Related Topics 链表 
示例:
输入:{1,2,3,4,5},2
输出:{2,1,4,3,5}

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
 * @param k int整型 
 * @return ListNode类
*/
func reverseKGroup( head *ListNode ,  k int ) *ListNode {
    // write code here
    //假设长度<k，不翻转
    tail := head
    //tail为第k+1个元素
    for i:=0;i<k;i++{
        if tail == nil{
            return head
        }
        tail = tail.Next
    }
    //找到了第k+1个
    var pre *ListNode
    cur:=head
    for cur!=tail{
        next:=cur.Next
        cur.Next=pre

        pre=cur
        cur=next
    }

    //遍历反转过后 pre就是这组翻转过来的第一个
    //令这组最后一个的next为下一组的翻转过来的第一个，也就是递归的返回值
    head.Next = reverseKGroup(tail,k)
    //遍历反转过后 pre就是这组翻转过来的第一个
    return pre
}
//nowcoder submit region end(Prohibit modification and deletion)
