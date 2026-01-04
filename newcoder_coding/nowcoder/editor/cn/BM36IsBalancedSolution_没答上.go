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
    //判断二叉树任意一个节点中，他的左子树和右子树的高度差是否>1，有则不是平衡二叉树
    // 核心函数：输入：root，计算左子树高度和右子树高度，如果相减>1,返回
    res:=getDep(pRoot)
    return res!=-1
}

func getDep(root *TreeNode) int{
    if root==nil{
        //没有节点的时候高度是0
        return 0
    }
    leftH:=getDep(root.Left)
    if leftH==-1{
        return -1 //如过高度差大于1 则返回值就是-1，不用再计算直接返回
    }
    rightH:=getDep(root.Right)
    if rightH==-1{
        return -1
    }

    //查看高度差是否为平衡
    if int(math.Abs(float64(rightH)-float64(leftH)))>1{
        return -1 //高度差大于1
    }else{
        high:=1+max(leftH,rightH)
        return high
    }

}

func max(x,y int)int{
    if x>y{
        return x
    }
    return y
}


//nowcoder submit region end(Prohibit modification and deletion)
