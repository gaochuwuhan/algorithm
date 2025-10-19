/**
一只青蛙一次可以跳上1级台阶，也可以跳上2级……它也可以跳上n级。求该青蛙跳上一个n级的台阶(n为正整数)总共有多少种跳法。 
 数据范围：
 进阶：空间复杂度 ， 时间复杂度 
 
 Related Topics 动态规划 递归 记忆化搜索 
示例:
输入:3
输出:4

*/
package nowcoder.editor.cn;  //根据实际修改
//nowcoder submit region begin(Prohibit modification and deletion)
package main
//import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 
 * @param number int整型 
 * @return int整型
*/
//不同的是从最后一条可以条1、2变成了可以条1-n，所以递推公式变成
//dp[i]=dp[i-1]+dp[i-2]+... 而到达n只和到达n-1的有关所以如果不用dp数组的话 要用一个lastCount变量把n-1的走法保存
//而sum就是将所有计算过的lastCount求和作为当前n的总和
func jumpFloorII( number int ) int {

    // write code here
    sum:=1
    lastCount:=1
    if number == 1{
        return sum
    }
    for i:=2;i<=number;i++{
        sum = sum+lastCount
        lastCount = sum
    }
    return sum
}
//nowcoder submit region end(Prohibit modification and deletion)
