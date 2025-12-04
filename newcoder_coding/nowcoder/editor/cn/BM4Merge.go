/**
输入两个递增的链表，单个链表的长度为n，合并这两个链表并使新链表中的节点仍然是递增排序的。
 数据范围： ，
 要求：空间复杂度 ，时间复杂度 
 
 如输入{1,3,5},{2,4,6}时，合并后的链表为{1,2,3,4,5,6}，所以对应的输出为{1,2,3,4,5,6}，转换过程如下图所示： 
 
 或输入{-1,2,4},{1,3,4}时，合并后的链表为{-1,1,2,3,4,4}，所以对应的输出为{-1,1,2,3,4,4}，转换过程如下图所示： 
 
 Related Topics 链表 2021 
示例:
输入:{1,3,5},{2,4,6}
输出:{1,2,3,4,5,6}

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
 * @param pHead1 ListNode类 
 * @param pHead2 ListNode类 
 * @return ListNode类
*/
func Merge( pHead1 *ListNode ,  pHead2 *ListNode ) *ListNode {
    // write code here
    if pHead1 ==nil{
        return pHead2
    }
    if pHead2 == nil{
        return pHead1
    }
    cur1:=pHead1
    cur2:=pHead2
    newHead:=&ListNode{}
    cur:=newHead
    for {
        if cur1 == nil || cur2 == nil{
            break
        }
        if cur1.Val<cur2.Val{
            cur.Next = cur1
            cur1 = cur1.Next
        }else{
            cur.Next = cur2
            cur2 = cur2.Next
        }
        cur = cur.Next

    }
    if cur1 != nil{
        cur.Next = cur1
    }
    if cur2 != nil{
        cur.Next = cur2
    }
    return newHead.Next
}
//nowcoder submit region end(Prohibit modification and deletion)
