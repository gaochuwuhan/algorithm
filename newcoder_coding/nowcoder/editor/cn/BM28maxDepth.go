/**
求给定二叉树的最大深度， 深度是指树的根节点到任一叶子节点路径上节点的数量。 最大深度是所有叶子节点的深度的最大值。 （注：叶子节点是指没有子节点的节点。）
 
 
 数据范围：，树上每个节点的val满足 
 要求： 空间复杂度 ,时间复杂度 
 
 Related Topics 树 dfs 
示例:
输入:{1,2}
输出:2 

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
 * @param root TreeNode类 
 * @return int整型
*/
func maxDepth( root *TreeNode ) int {
    // write code here
    if root == nil{
        return 0
    }
    maxD:=0
    depth:=1
    walk(depth,&maxD,root)
    return maxD
}

func walk(depth int,maxD *int, root *TreeNode) {
    if root.Left==nil && root.Right==nil{
        *maxD = max(*maxD,depth)
        return
    }
    if root.Left!=nil{
        walk(depth+1,maxD,root.Left)
    }

    if root.Right !=nil{
        walk(depth+1,maxD,root.Right)
    }
}

func max(x,y int) int{
    if x>y{
        return x
    }
    return y
}

//nowcoder submit region end(Prohibit modification and deletion)
