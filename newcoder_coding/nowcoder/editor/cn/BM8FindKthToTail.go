/**
输入一个长度为 n 的链表，设链表中的元素的值为 ai ，返回该链表中倒数第k个节点。 如果该链表长度小于k，请返回一个长度为 0 的链表。 
 
 数据范围：，， 要求：空间复杂度 ，时间复杂度 进阶：空间复杂度 ，时间复杂度 
 例如输入{1,2,3,4,5},2时，对应的链表结构如下图所示： 其中蓝色部分为该链表的最后2个结点，所以返回倒数第2个结点（也即结点值为4的结点）即可，系统
会打印后面所有的节点来比较。
 
 Related Topics 链表 
示例:
输入:{1,2,3,4,5},2
输出:{4,5}

*/
package nowcoder.editor.cn; //根据实际修改



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
 * @param pHead ListNode类 
 * @param k int整型 
 * @return ListNode类
*/
//func FindKthToTail( pHead *ListNode ,  k int ) *ListNode {
//    // write code here
//    if pHead==nil||k==0{
//        return nil
//    }
//
//    fast,slow:=pHead,pHead
//    n:=1
//    for {
//        if fast == nil{
//            if n<k{ //还没有走到第k个 所以长度小于k
//                return nil
//            }
//            return slow
//        }
//
//        fast =fast.Next
//        n++
//        if n>=k{
//            slow = slow.Next
//        }
//    }
//}
func FindKthToTail( pHead *ListNode ,  k int ) *ListNode {
    // write code here
    if pHead == nil || k == 0 {
        return nil
    }

    fast, slow := pHead, pHead
    for i:=0;i<k;i++{
        if fast==nil{
            return nil
        }
        fast = fast.Next

    }

    for fast!=nil{
        slow = slow.Next
        fast=fast.Next
    }
    return  slow
}
//nowcoder submit region end(Prohibit modification and deletion)