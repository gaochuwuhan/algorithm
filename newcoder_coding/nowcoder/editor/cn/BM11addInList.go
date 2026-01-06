/**
假设链表中每一个节点的值都在 0 - 9 之间，那么链表整体就可以代表一个整数。 给定两个这种链表，请生成代表两个整数相加值的结果链表。 数据范围：，链表任意值
 
 要求：空间复杂度 ，时间复杂度 
 
 例如：链表 1 为 9->3->7，链表 2 为 6->3，最后生成新的结果链表为 1->0->0->0。 
 
 Related Topics 链表 模拟 
示例:
输入:[9,3,7],[6,3]
输出:{1,0,0,0}

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
 * @param head1 ListNode类 
 * @param head2 ListNode类 
 * @return ListNode类
*/
func addInList( head1 *ListNode ,  head2 *ListNode ) *ListNode {
    // write code here
	if head1==nil{
		return head2
	}
	if head2==nil{
		return head1
	}
	h1,h2:=make([]*ListNode,0),make([]*ListNode,0)
	cur1:=head1
	cur2:=head2
	for cur1!=nil{
		h1 = append(h1,cur1)
		cur1=cur1.Next
	}
	for cur2!=nil{
		h2 = append(h2,cur2)
		cur2=cur2.Next
	}

	i1:=len(h1)-1
	i2:=len(h2)-1
	flag:=false // 前一个指针是否需要进位
	var afterPre *ListNode
	for i1>=0 && i2>=0{
		node1:=h1[i1]
		node2:=h2[i2]
		node:=&ListNode{Val: node2.Val+node1.Val}
		if flag{
			node.Val++
		}
		if node.Val>=10{
			node.Val=node.Val-10
			flag=true
		}else{
			flag=false
		}
		if afterPre==nil{
			afterPre=node
			i1--
			i2--
			continue
		}
		tmp:=afterPre
		node.Next=tmp
		afterPre=node//上一次的连续的链接链表
		i1--
		i2--
	}
	//剩余节点还没有加
	if i1>=0{
		for j1:=i1;j1>=0;j1--{
			node:=&ListNode{Val: h1[j1].Val}
			if flag{
				node.Val++
			}
			if node.Val>=10{
				node.Val=node.Val-10
				flag=true
			}else{
				flag=false
			}
			temp:=afterPre
			node.Next=temp
			afterPre=node
		}
	}
	if i2>=0{
		for j1:=i2;j1>=0;j1--{
			node:=&ListNode{Val: h2[j1].Val}
			if flag{
				node.Val++
			}
			if node.Val>=10{
				node.Val=node.Val-10
				flag=true
			}else{
				flag=false
			}
			temp:=afterPre
			node.Next=temp
			afterPre=node
		}
	}
	//判断头之前的相加值是否有进位
	if flag{
		node:=&ListNode{
			Val: 1,
		}
		tmp:=afterPre
		node.Next=tmp
		afterPre=node
	}
	return afterPre

}
//nowcoder submit region end(Prohibit modification and deletion)
