/**
给出一个仅包含字符仅由括号字符 ;、、、、、 ;的括号序列字符串 （），你需要判断给出的括号序列字符串 ; ;是否是有效的括号序列。 有效括号序列的定义如下： 
空序列是有效括号序列； 如果 ; ;是有效括号序列，则 、 和 都是有效括号序列； 如果 ; ;和 ; ;都是有效括号序列，则它们的拼接 ; ;也是有效括号序列
。 如果括号序列字符串 ; ;是有效的括号序列，返回一个布尔值 ;；否则返回一个布尔值 ;。 
 Related Topics 栈 字符串 
示例:
输入:"["
输出:false 

*/
package nowcoder.editor.cn;  //根据实际修改
//nowcoder submit region begin(Prohibit modification and deletion)
package main
//import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 
 * @param s string字符串 
 * @return bool布尔型
*/

var leftPairs = map[string] string{
    "{":"}",
    "[":"]",
    "(": ")",
}
func isValid( s string ) bool {
    // write code here
    if len(s) == 1{
        return false
    }
    leftStack:=[]string{}
    for _,v:=range s{
        vs:=string(v)
        if len(leftStack) == 0{
            _,has:=leftPairs[vs]
            if !has{
                return false
            }
            //放到left栈里
            leftStack = append(leftStack,vs)
            continue
        }
        //判断是左括号就放到栈里，否则进行pair比对
        _,has:=leftPairs[vs]
        if !has{
            //右括号判断
            paired:=isPair(&leftStack,vs)
            if !paired{
                return false
            }
        }else{
            leftStack = append(leftStack,vs)
        }
    }
    //记得最后判断栈里的是否都比对过了，有可能是"((" 残留
    if len(leftStack) > 0{
        return false
    }
    return true

}

func isPair(leftStack *[]string, rightStr string) bool{
    //判断栈顶元素和rightStr是否配对
    topNode:=(*leftStack)[len(*leftStack)-1]
    *leftStack = (*leftStack)[0:len(*leftStack)-1]
    pairVal,_:=leftPairs[topNode]

    if pairVal == rightStr{
        return true
    }
    return false


}
//nowcoder submit region end(Prohibit modification and deletion)
