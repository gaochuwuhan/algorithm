/**
 * @nc app=nowcoder id=508378c0823c423baa723ce448cbfd0c topic=295 question=634 lang=Go
 * 2025-09-26 01:04:07
 * https://www.nowcoder.com/practice/508378c0823c423baa723ce448cbfd0c?tpId=295&tqId=634
 * [BM29] 二叉树中和为某一值的路径(一)
 */

/** @nc code=start */

package main
// import "fmt"
import . "nc_tools"
/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 
 * @param root TreeNode类 
 * @param sum int整型 
 * @return bool布尔型
*/
func walk(tree *TreeNode, curSum int, sum int, has *bool) {
	if *has {
		return
	}
	if tree.Left == nil && tree.Right == nil {
		if curSum+tree.Val == sum {
			*has = true
		}
		return
	}
	curSum = curSum + tree.Val
	if tree.Left != nil {
		walk(tree.Left, curSum, sum, has)
	}
	if tree.Right !=nil{
		walk(tree.Right, curSum, sum, has)
	}
}
func hasPathSum( root *TreeNode ,  sum int ) bool {
    // write code here
	if root == nil {
		return false
	}
	curSum := 0
	has := false
	walk(root, curSum, sum, &has)
	return has
}

/** @nc code=end */
