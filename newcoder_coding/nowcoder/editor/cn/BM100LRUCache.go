/**
设计LRU(最近最少使用)缓存结构，该结构在构造时确定大小，假设大小为 capacity ，操作次数是 n ，并有如下功能:
 1. Solution(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
 2. get(key)：如果关键字 key 存在于缓存中，则返回key对应的value值，否则返回 -1 。
 3. set(key, value)：将记录(key, value)插入该结构，如果关键字 key 已经存在，则变更其数据值 value，如果不存在，则向缓存
中插入该组 key-value ，如果key-value的数量超过capacity，弹出最久未使用的key-value
 
 提示:
 1.某个key的set或get操作一旦发生，则认为这个key的记录成了最常使用的，然后都会刷新缓存。
 2.当缓存的大小超过capacity时，移除最不经常使用的记录。
 3.返回的value都以字符串形式表达，如果是set，则会输出"null"来表示(不需要用户返回，系统会自动输出)，方便观察
 4.函数set和get必须以O(1)的方式运行
 5.为了方便区分缓存里key与value，下面说明的缓存里key用""号包裹 数据范围: 
 
 
 
 
 Related Topics 链表 哈希 模拟 
示例:
输入:["set","set","get","set","get","set","get","get","get"],[[1,1],[2,2],[1],[3,3
],[2],[4,4],[1],[3],[4]],2
输出:["null","null","1","null","-1","null","-1","3","4"]

*/
package nowcoder.editor.cn; //根据实际修改




import "strconv"




//nowcoder submit region begin(Prohibit modification and deletion)
package main
import "strconv"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 
 * @param operators string字符串一维数组 
 * @param param int整型二维数组 
 * @param capacity int整型 
 * @return string字符串一维数组
*/
func LRUCache( operators []string ,  param [][]int ,  capacity int ) []string {
    // write code here
    //lru 使用了双链表来存储具体元素，get的时候将该节点放到头部，set的时候需判断当前双链表长度，如果未超出放到当前为尾部的Next，并指向head；
    //超出了就将尾部的节点删除
    lru:=Constructor(capacity)
    res:=[]string{}
    for i:=0;i<len(operators);i++{
        op:=operators[i]
        pa:=param[i]
        if op=="set"{
            lru.set(pa[0],pa[1])
            res = append(res,"null")
        }else{
            val:=lru.get(pa[0])
            res = append(res,strconv.Itoa(val))
        }
    }
    return res


}

type Solution struct {
    // write code here
    Head *DoubleLinkNode
    Tail *DoubleLinkNode
    cap int
    length int
    store map[int]*DoubleLinkNode

}


func Constructor(capacity int) Solution {
    // write code here
    l:=Solution{
        Head:   nil,
        Tail:   nil,
        cap:    capacity,
        length: 0,
        store:  make(map[int]*DoubleLinkNode),
    }
    return l
}


func (this *Solution) get(key int) int {
    // write code here
    if this.length==1{
        return this.Head.Val
    }
    node,ok:=this.store[key]
    if ok{
        //node 前后两个节点相连
        nodeNext:=node.Next
        nodePre:=node.Pre
        nodeNext.Pre = nodePre
        nodePre.Next = nodeNext
        //将node放到head
        head:=this.Head
        node.Next = head
        head.Pre=node
        this.Head=node
        return node.Val

    }
    return -1
}


func (this *Solution) set(key int, value int) {
    // write code here
    node:=&DoubleLinkNode{Val: value}

    if this.length<this.cap{
        if this.Head==nil{
            this.Head = node

        }else{
            //head不为空，放到最后
            tail:=this.Tail
            if tail==nil{
                this.Tail=node
                this.Head.Next=this.Tail
                this.Tail.Pre=this.Head
            }else{
                node.Pre = tail
                tail.Next=node
                this.Tail=node
            }

        }
        this.length++
    }else{
        //满了
        if this.cap==1{
            delete(this.store,key)
            this.Head=node
        }else{
            tail:=this.Tail
            delete(this.store,key)
            tailPre:=tail.Pre
            tailPre.Next=node
            node.Pre = tailPre
        }
    }
    // 放map里
    this.store[key] = node
}
type DoubleLinkNode struct {
    Pre *DoubleLinkNode
    Next *DoubleLinkNode
    Val int

}


//nowcoder submit region end(Prohibit modification and deletion)
