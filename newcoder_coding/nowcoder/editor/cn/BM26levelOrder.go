/**
给定一个二叉树，返回该二叉树层序遍历的结果，（从左到右，一层一层地遍历）
 例如：
 给定的二叉树是{3,9,20,#,#,15,7},
 
 该二叉树层序遍历的结果是
 [
 [3],
 [9,20],
 [15,7] ] 
 
 提示: 0 <= 二叉树的结点数 <= 1500 
 
 
 Related Topics 树 广度优先搜索(BFS) 
示例:
输入:{1,2}
输出:[[1],[2]] 

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
 * @return int整型二维数组
*/
func levelOrder( root *TreeNode ) [][]int {
    // write code here
    if root==nil{
        return nil
    }
    queue:=NewQueue()
    cur:=root
    queue.Enqueue(cur)
    ret:=make([][]int,0)
    for !queue.Isempty(){
        //出队列
        layer:=make([]int,0)//仅将一层的放进队列
        enq:=make([]*TreeNode,0) //暂存作为一下次入队列的一层
        //这里for循环就是一层的遍历，都放到了队列里
        for {
            node:=queue.Dequeue()
            layer = append(layer,node.Val)
            if node.Left!=nil{
                enq = append(enq,node.Left)
            }
            if node.Right != nil{
                enq = append(enq,node.Right)
            }
            if queue.Isempty(){
                if len(layer) >0{
                    ret =append(ret,layer)
                }
                break
            }
        }
        if len(enq) >0{
            for _,node:=range enq{
                queue.Enqueue(node)
            }
        }

    }
    return ret

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
