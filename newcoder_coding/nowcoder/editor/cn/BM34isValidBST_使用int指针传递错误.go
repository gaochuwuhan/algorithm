/**
给定一个二叉树根节点，请你判断这棵树是不是二叉搜索树。 
 二叉搜索树满足每个节点的左子树上的所有节点均小于当前节点且右子树上的所有节点均大于当前节点。 
 例： 图1 
 图2 
 数据范围：节点数量满足 ，节点上的值满足 
 
 Related Topics 树 
示例:
输入:{1,2,3}
输出:false 

*/
package nowcoder.editor.cn; //根据实际修改


import "math"


//nowcoder submit region begin(Prohibit modification and deletion)
package main
//import "fmt"
import . "nc_tools"
import "math"

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
 * @return bool布尔型
*/
func isValidBST( root *TreeNode ) bool {
    // write code here
    if root==nil{
        return true
    }
    prev:=math.MinInt
    var midOrd func (treeRoot *TreeNode) bool
    midOrd= func (treeRoot *TreeNode) bool{
        if treeRoot==nil{
            return true
        }
        if !midOrd(treeRoot.Left){
            return false
        }
        if treeRoot.Val<=prev{
            return false
        }
        prev=treeRoot.Val
        return midOrd(treeRoot.Right)
    }
    return midOrd(root)

}


//nowcoder submit region end(Prohibit modification and deletion)

