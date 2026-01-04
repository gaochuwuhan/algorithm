/**
给定一棵二叉树，判断其是否是自身的镜像（即：是否对称）
 例如： 下面这棵二叉树是对称的
 
 下面这棵二叉树不对称。
 
 数据范围：节点数满足 ，节点上的值满足 要求：空间复杂度 ，时间复杂度 备注： 你可以用递归和迭代两种方法解决这个问题 
 Related Topics 树 
示例:
输入:{1,2,2,3,4,4,3}
输出:true 

*/
package nowcoder.editor.cn;  //根据实际修改
//nowcoder submit region begin(Prohibit modification and deletion)
package main
//import "fmt"
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
 * @param pRoot TreeNode类 
 * @return bool布尔型
*/
func isSymmetrical( pRoot *TreeNode ) bool {
    // write code here
    if pRoot==nil{
        return true
    }
    leftTree:=pRoot.Left
    rightTree:=pRoot.Right
    isDuiChen:=true
    preOrderTree(leftTree,rightTree,&isDuiChen)
    return isDuiChen
}

func preOrderTree(rootL *TreeNode,rootR *TreeNode,isDuiChen *bool){
    if *isDuiChen==false{
        return
    }
    if rootR==nil && rootL!=nil{
        *isDuiChen=false
        return
    }
    if rootL==nil && rootR!=nil{
        *isDuiChen=false
        return
    }

    if rootL!=nil && rootR!=nil{

        if rootR.Val!=rootL.Val{
            *isDuiChen=false //对称还有一个判断条件是值得相等，不能只判断形状
            return
        }
        preOrderTree(rootL.Left,rootR.Right,isDuiChen)
        preOrderTree(rootL.Right,rootR.Left,isDuiChen)

    }
}

//nowcoder submit region end(Prohibit modification and deletion)
