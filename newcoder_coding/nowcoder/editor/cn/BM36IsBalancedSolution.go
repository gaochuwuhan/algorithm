/**
输入一棵节点数为 n 二叉树，判断该二叉树是否是平衡二叉树。 在这里，我们只需要考虑其平衡性，不需要考虑其是不是排序二叉树 平衡二叉树（Balanced 
Binary Tree），具有以下性质：它是一棵空树或它的左右两个子树的高度差的绝对值不超过1，并且左右两个子树都是一棵平衡二叉树。
 样例解释： 样例二叉树如图，为一颗平衡二叉树
 注：我们约定空树是平衡二叉树。 
 数据范围：,树上节点的val值满足 要求：空间复杂度，时间复杂度 
 Related Topics 树 dfs 
示例:
输入:{1,2,3,4,5,6,7}
输出:true

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
 * @param pRoot TreeNode类 
 * @return bool布尔型
*/
func IsBalanced_Solution( pRoot *TreeNode ) bool {
    // write code here
    if pRoot==nil{
        return true
    }
    return getDepth(pRoot)!=-1
}

func getDepth(tree *TreeNode) int{
    if tree == nil{
        return 0
    }

    ld:=getDepth(tree.Left)
    if ld==-1{
        return -1
    }
    rd:=getDepth(tree.Right)
    if rd==-1{
        return -1
    }
    if int(math.Abs(float64(ld)-float64(rd)))>1{
        return -1
    }
    //返回下面节点的高度加上当前root节点1就是总高度
    return 1+int(math.Max(float64(ld),float64(rd)))
}

//nowcoder submit region end(Prohibit modification and deletion)
