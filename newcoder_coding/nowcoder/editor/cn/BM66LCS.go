/**
给定两个字符串str1和str2,输出两个字符串的最长公共子串 题目保证str1和str2的最长公共子串存在且唯一。 
 数据范围： 
 要求： 空间复杂度 ，时间复杂度 
 
 Related Topics 动态规划 
示例:
输入:"1AB2345CD","12345EF"
输出:"2345" 

*/
package nowcoder.editor.cn;  //根据实际修改
//nowcoder submit region begin(Prohibit modification and deletion)
package main
//import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * longest common substring
 * @param str1 string字符串 the string
 * @param str2 string字符串 the string
 * @return string字符串
*/
//第二次画图做上来了
func LCS( str1 string ,  str2 string ) string {
    // write code here
    if str1 == "" || str2 == ""{
        return ""
    }

    dp:=make([][]int, len(str1)+1)
    for i:=range dp{
        dp[i] = make([]int,len(str2)+1)
    }

    //记录最大长度和最后一个index
    maxL:=0
    var lastI int

    for i:=1;i<=len(str1);i++{
        for j:=1;j<=len(str2);j++{
            if str1[i-1] == str2[j-1]{
                dp[i][j] = dp[i-1][j-1]+1
                if dp[i][j] > maxL{
                    maxL = dp[i][j]
                    lastI = i-1
                }
            }
        }
    }
    if maxL==0{
        return ""
    }
    startI:=lastI-maxL+1
    return str1[startI:lastI+1]
}
//nowcoder submit region end(Prohibit modification and deletion)