//给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。 
//
// 如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。 
//
// 
//
// 示例 1： 
// 
// 
//输入：p = [1,2,3], q = [1,2,3]
//输出：true
// 
//
// 示例 2： 
// 
// 
//输入：p = [1,2], q = [1,null,2]
//输出：false
// 
//
// 示例 3： 
// 
// 
//输入：p = [1,2,1], q = [1,1,2]
//输出：false
// 
//
// 
//
// 提示： 
//
// 
// 两棵树上的节点数目都在范围 [0, 100] 内 
// -10⁴ <= Node.val <= 10⁴ 
// 
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 1254 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    isSame:=true
	compareTree(p,q,&isSame)
	return isSame
}

func compareTree(p *TreeNode,q *TreeNode, isSame *bool){
	if *isSame == false{
		return
	}
	if p==nil{
		if q != nil{
			*isSame=false
			return
		}
		return
	}
	if q==nil{
		if p != nil{
			*isSame=false
			return
		}
		return
	}
	if p.Val != q.Val{
		*isSame=false
		return
	}
	compareTree(p.Left,q.Left,isSame)
	compareTree(p.Right,q.Right,isSame)
}


//leetcode submit region end(Prohibit modification and deletion)
