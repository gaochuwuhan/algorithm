/**
给定一棵二叉树，已知其中的节点没有重复值，请判断该二叉树是否为搜索二叉树和完全二叉树。 输出描述：分别输出是否为搜索二叉树、完全二叉树。 
 
 数据范围：二叉树节点数满足 ，二叉树上的值满足 要求：空间复杂度 ，时间复杂度 
 注意：空子树我们认为同时符合搜索二叉树和完全二叉树。 
 Related Topics 树 dfs 广度优先搜索(BFS) 
示例:
输入:{2,1,3}
输出:[true,true]

*/
package nowcoder.editor.cn;  //根据实际修改
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
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
 * @param root TreeNode类 the root
 * @return bool布尔型一维数组
*/


func dequeue(queue *[]*TreeNode) *TreeNode{
    data:=(*queue)[0]
    *queue = (*queue)[1:len(*queue)]
    return data
}

func enqueue(data *TreeNode,queue *[]*TreeNode){
    *queue = append(*queue,data)
}

func isCompTree(tree *TreeNode) bool{
    if tree == nil{
        return true
    }
    //lackR:=0
    mustLeaf:=false //一旦碰到缺左子节点或右子节点那么之后的就一定是叶子节点
    queue:=make([]*TreeNode,0)
    enqueue(tree,&queue)
    for len(queue) >0{
        curTree:=dequeue(&queue)
        if curTree.Left == nil&& curTree.Right!=nil{
            return false
        }

        if curTree.Left!=nil{
            if mustLeaf{
                return false
            }
            enqueue(curTree.Left,&queue)

        }else{
            mustLeaf=true
        }
        if curTree.Right != nil{
            if mustLeaf{
                return false
            }
            enqueue(curTree.Right,&queue)

        }else {
            mustLeaf=true
        }
    }
    return true

}

func isSearchTree(tree *TreeNode,prev *int,ok *bool){
    if !*ok{
        return
    }
    if tree == nil{
        *ok=true
        return
    }
    isSearchTree(tree.Left,prev,ok)
    if tree.Val<=*prev{
        *ok=false
        return
    }
    *prev = tree.Val
    isSearchTree(tree.Right,prev,ok)
}
//{3,2,5,1,4}
func judgeIt( root *TreeNode ) []bool {
    // write code here
    if root == nil{
        return []bool{true,true}
    }
    comFlag:=isCompTree(root)
    prev:=-1
    searchFlag:=true
    isSearchTree(root,&prev,&searchFlag)
    return []bool{searchFlag,comFlag}
}
//nowcoder submit region end(Prohibit modification and deletion)
