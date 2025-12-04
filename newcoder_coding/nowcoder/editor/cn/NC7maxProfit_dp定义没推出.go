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
package nowcoder.editor.cn;  //根据实际修改
//nowcoder submit region begin(Prohibit modification and deletion)
package main
//import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 
 * @param prices int整型一维数组 
 * @return int整型
*/

func Max(x,y int)int{
    if x>y{
        return x
    }
    return y
}

func maxProfit1( prices []int ) int {
    // write code here
    //无收益情况
    if len(prices)<2{
        return 0
    }
    maxSum:=0
    //使用二维数组记录买入i和卖出j 对应的收益情况，其中j>i
    //dp:=make([][]int,len(prices))
    //for i:=range dp{
    //    dp[i]=make([]int,len(prices))
    //}
    for i:=0;i<len(prices);i++{
        for j:=i+1;j<len(prices);j++{
            money:=prices[j]-prices[i]
            maxSum = Max(money,maxSum)
        }
    }
    return maxSum
}
//dp解法：变的东西是天数、买卖单价、和当天收益。递推公式需要找到当前一次的收益和上一次收益的产生关联，可以想到 第i天的收益取决于第i-1天的收益和第i天卖还是买
//但是买和卖只能进行一次，所以 递推公式卖和买不作为状态转移，将第i天持股、不持股作为dp一个状态0代表持股，1代表不持股
//dp[i][0]=max(dp[i-1][0],-price[i]) 维持昨天持股，或者今天第一次买 因为只能买一次，所以 后面是-price[i]
//dp[i][1]=max(dp[i-1][1],dp[i-1][0]+price[i]) 维持昨天不持股 或者今天卖掉
//base case: 第0天买了dp[0][0]=-price[0]; 不买dp[0][1]=0
//遍历顺序：由于后面的天数依赖前一天的，所以是正序遍历，已经确定了第0天的base case值，遍历从1开始
func maxProfit( prices []int ) int {
    // write code here
    //无收益情况
    if len(prices)<2{
        return 0
    }
    //使用二维数组记录买入i和卖出j 对应的收益情况，其中j>i
    dp:=make([][]int,len(prices))
    for i:=range dp{
       dp[i]=make([]int,2)
    }
    dp[0][0]=-prices[0]
    dp[0][1]=0
    for i:=1;i<len(prices);i++{
        dp[i][0]=Max(dp[i-1][0],-prices[i])
        dp[i][1]=Max(dp[i-1][1],dp[i-1][0]+prices[i])
    }
    maxSum:=Max(dp[len(prices)-1][0],dp[len(prices)-1][1])

    return maxSum
}
//nowcoder submit region end(Prohibit modification and deletion)
