/**
假设你有一个数组prices，长度为n，其中prices[i]是股票在第i天的价格，请根据这个价格数组，返回买卖股票能获得的最大收益 1.你可以买入一次股票和卖
出一次股票，并非每天都可以买入或卖出一次，总共只能买入和卖出一次，且买入必须在卖出的前面的某一天 2.如果不能获取到任何利润，请返回0 3.假设买入卖出均无手续
费
 
 数据范围： 
 要求：空间复杂度 ，时间复杂度 
 
 Related Topics 动态规划 贪心 
示例:
输入:[8,9,2,5,4,7,1]
输出:5 

*/
package nowcoder


//nowcoder submit region begin(Prohibit modification and deletion)
package main


/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 
 * @param prices int整型一维数组 
 * @return int整型
*/

//dp[i][0]和dp[i][1]分别代表第i天不持有股票和持有股票
// 表达式：dp[i][0]=max(dp[i-1][0],dp[i-1][1]+price[i]); dp[i][1]=max(dp[i-1][1],-price[i])
// base case：dp[0][0]=0 dp[0][1]=-price[0]
// 维护一个max值，计算每次持有和持有的最大值
func maxProfit( prices []int ) int {
    // write code here
    var maxGet=-1
    dp:=make([][]int,len(prices))
    for i:=range dp{
        dp[i] = make([]int,2)
    }
    dp[0][0]= 0
    dp[0][1]=-prices[0]
    for i:=1;i<len(prices);i++{
        dp[i][0] = Max(dp[i-1][0],dp[i-1][1]+prices[i])
        dp[i][1] = Max(dp[i-1][1],-prices[i])
        //maxGet=Max(maxGet,dp[i][1])
    }
    maxGet=Max(dp[len(prices)-1][0],dp[len(prices)-1][1])

    return maxGet

}

func Max(x,y int) int{
    if x>y{
        return x
    }
    return y
}
//nowcoder submit region end(Prohibit modification and deletion)

