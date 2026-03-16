/**
用两个栈来实现一个队列，使用n个元素来完成 n 次在队列尾部插入整数(push)和n次在队列头部删除整数(pop)的功能。 队列中的元素为int类型。保证操作合
法，即保证pop操作时队列内已有元素。 
 数据范围： 要求：存储n个元素的空间复杂度为 ，插入与删除的时间复杂度都是 
 Related Topics 栈 
示例:
输入:["PSH1","PSH2","POP","POP"]
输出:1,2

*/
package nowcoder.editor.cn;  //根据实际修改
//nowcoder submit region begin(Prohibit modification and deletion)
package main

var stack1 [] int //入队列，
var stack2 [] int //出队列，如果空了，stack1全部出栈到stack2中

func Push(node int) {
    stack1 = append(stack1,node)
}

func Pop() int{
    //
    if len(stack2) == 0{
        for i:=len(stack1)-1;i>=0;i--{
            stack2 = append(stack2,stack1[i])
        }
        stack1 = []int{}
    }
    data:=stack2[len(stack2)-1]
    stack2 = stack2[0:len(stack2)-1]
    return data
}
//nowcoder submit region end(Prohibit modification and deletion)
