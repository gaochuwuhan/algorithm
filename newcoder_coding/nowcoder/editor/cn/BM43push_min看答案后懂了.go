/**
定义栈的数据结构，请在该类型中实现一个能够得到栈中所含最小元素的 min 函数，输入操作时保证 pop、top 和 min 函数操作时，栈中一定有元素。 
 此栈包含的方法有： push(value):将value压入栈中 pop():弹出栈顶元素 top():获取栈顶元素 min():获取栈中最小元素 
 数据范围：操作数量满足 ，输入的元素满足 
 进阶：栈的各个操作的时间复杂度是 ，空间复杂度是 
 
 示例: 输入: ["PSH-1","PSH2","MIN","TOP","POP","PSH1","TOP","MIN"] 输出: -1,2,1,-1 解析:
 "PSH-1"表示将-1压入栈中，栈中元素为-1 "PSH2"表示将2压入栈中，栈中元素为2，-1 “MIN”表示获取此时栈中最小元素==>返回-1 
"TOP"表示获取栈顶元素==>返回2 "POP"表示弹出栈顶元素，弹出2，栈中元素为-1 "PSH1"表示将1压入栈中，栈中元素为1，-1
 "TOP"表示获取栈顶元素==>返回1 “MIN”表示获取此时栈中最小元素==>返回-1
 
 
 Related Topics 栈 
示例:
输入: ["PSH-1","PSH2","MIN","TOP","POP","PSH1","TOP","MIN"]
输出:-1,2,1,-1

*/
package nowcoder.editor.cn;  //根据实际修改
//nowcoder submit region begin(Prohibit modification and deletion)
package main
//使用辅助栈实现，主栈正常push，pop，push时，对比辅助栈的最顶元素，小的继续push放到辅助栈，pop两个栈都正常pop
var stack43 []int
var helpStack []int
func Push(node int) {
    // write code here
    stack43 = append(stack43,node)

    if len(helpStack)==0{
        helpStack = append(helpStack,node)
    }else{
        helpTop:=helpStack[len(helpStack)-1]
        minNode:=minVal(node,helpTop)
        helpStack = append(helpStack,minNode)
    }
}
func Pop() {
    // write code here
    stack43 = stack43[0:len(stack43)-1]
    helpStack = helpStack[0:len(helpStack)-1]
}
func Top() int {
    // write code here
    return stack43[len(stack43)-1]
}
func Min() int {
    // write code here
    return helpStack[len(helpStack)-1]
}

func minVal(x,y int)int{
    if x<y{
        return x
    }
    return y
}
//nowcoder submit region end(Prohibit modification and deletion)
