/**
给定一个非负索引值 num ，请返回杨辉三角中从上到下第 num 层。索引值从 0 开始。 杨辉三角中，每个数是左上方和右上方的数之和。 
 
 数据范围： 
 
 例如当输入3时，对应的输出为[1,3,3,1]， 杨辉三角的第3行（从0开始算起）部分如下图蓝色部分所示： 
 
 Related Topics 动态规划 数组 
示例:
输入:0
输出:[1] 

*/
package nowcoder.editor.cn;  //根据实际修改
//nowcoder submit region begin(Prohibit modification and deletion)
package main
//import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 
 * @param num int整型 
 * @return int整型一维数组
*/
func getRow( num int ) []int {

    // write code here
    line1:=[]int{1}
    line2:=[]int{1,1}

    if num==0{
        return line1
    }
    if num==1{
        return line2
    }
    dp:=make([][]int,num+1)
    dp[0]=line1
    dp[1]=line2
    if num<=1{
        return dp[num]
    }
    for i:=2;i<=num;i++{
        data:=make([]int,i+1)
        for j:=range data{
            if j==0||j==i{
                data[j]=1
            }else{
                data[j]=dp[i-1][j-1]+dp[i-1][j]
            }
        }
        dp[i] = data
    }
    return dp[num]
}
//nowcoder submit region end(Prohibit modification and deletion)
