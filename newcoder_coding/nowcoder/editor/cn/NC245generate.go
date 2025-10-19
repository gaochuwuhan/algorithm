/**
给定一个非负整数 num ，生成杨辉三角的前 num 行。 杨辉三角中，每个数是左上方和右上方的数之和。 
 数据范围： 
 
 例如当输入为4时，对应的返回值为[[1],[1,1],[1,2,1],[1,3,3,1]]，打印结果如下图所示： 
 
 Related Topics 动态规划 数组 
示例:
输入:1
输出:[[1]] 

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
 * @return int整型二维数组
*/
func generate( num int ) [][]int {

    // write code here
    dp:=make([][]int,num)
    for i:=range dp{
        dp[i] = make([]int,i+1)
    }
    line1:=[]int{1}

    line2:=[]int{1,1}
    dp[0] = line1
    if num<=1{
        return dp
    }
    dp[1] = line2
    if num<3{
        return dp
    }

    for i:=3;i<=num;i++{
        data:=make([]int,i)
        for j:=range data{
            if j==0 || j==i-1{
                data[j] = 1
            }else{
                data[j] = dp[i-2][j-1]+dp[i-2][j]
            }
        }
        dp[i-1]=data
    }
    return dp
}
//nowcoder submit region end(Prohibit modification and deletion)