/**
给定一个二叉树，确定他是否是一个完全二叉树。 
 完全二叉树的定义：若二叉树的深度为 h，除第 h 层外，其它各层的结点数都达到最大个数，第 h 层所有的叶子结点都连续集中在最左边，这就是完全二叉树。（第 
h 层可能包含 [1~2h] 个节点）
 
 数据范围：节点数满足 
 样例图1： 样例图2： 样例图3： 
 
 Related Topics 树 dfs 广度优先搜索(BFS) 
示例:
输入:{1,2,3,4,5,6}
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
 * @param root TreeNode类 
 * @return bool布尔型
*/
func isCompleteTree( root *TreeNode ) bool {
    // write code here
    q:=NewQueue()
    lack:=false//标记之前是否缺少节点
    q.Enqueue(root)
    for !q.Isempty(){
        node:=q.Dequeue()
        if node==nil{
            lack=true
            continue
        }
        if lack{
            return false //当有节点的时候，发现之前缺过节点必然不是完全二叉树
        }
        q.Enqueue(node.Left)
        q.Enqueue(node.Right)
    }
    return true

}

type Queue struct {
    queue []*TreeNode
}

func NewQueue() *Queue{
    return &Queue{
        queue: make([]*TreeNode, 0),
    }
}

func (q *Queue) Enqueue(tree *TreeNode) {
    q.queue = append(q.queue,tree)
}

func (q *Queue) Dequeue() *TreeNode{
    if !q.Isempty(){
        tree:=q.queue[0]
        l:=len(q.queue)
        cur:=q.queue[1:l]
        q.queue = cur
        return tree
    }
    return nil
}

func (q *Queue) Isempty() bool{
    return len(q.queue)==0
}

//nowcoder submit region end(Prohibit modification and deletion)
