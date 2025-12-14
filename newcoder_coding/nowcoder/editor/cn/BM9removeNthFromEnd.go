/**
给定一个链表，删除链表的倒数第 n 个节点并返回链表的头指针
 例如， 给出的链表为: , .
 删除了链表的倒数第 个节点之后,链表变为. 
 数据范围： 链表长度 ，链表中任意节点的值满足 要求：空间复杂度 ，时间复杂度 
 备注： 题目保证 一定是有效的
 
 Related Topics 链表 双指针 
示例:
输入:{1,2},2 
输出:{2} 

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
 * @param n int整型 
 * @return ListNode类
*/
func removeNthFromEnd( head *ListNode ,  n int ) *ListNode {
    // write code here
    if head == nil || head.Next==nil{
        return nil
    }
    newH:=&ListNode{Next: head}
    slowPre,fast:=newH,head

    //先找fast
    for i:=0;i<n;i++{
        fast = fast.Next
    }
    for fast!=nil{
        fast = fast.Next
        slowPre = slowPre.Next
    }
    slow:=slowPre.Next
    slowN:=slow.Next
    slowPre.Next = slowN
    return newH.Next



}
//nowcoder submit region end(Prohibit modification and deletion)
