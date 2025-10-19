/**
对于长度为n的一个字符串A（仅包含数字，大小写英文字母），请设计一个高效算法，计算其中最长回文子串的长度。 
 数据范围： 要求：空间复杂度 ，时间复杂度 进阶: 空间复杂度 ，时间复杂度 
 Related Topics 字符串 动态规划 
示例:
输入:"ababc"
输出:3 

*/
package nowcoder.editor.cn; //根据实际修改
//nowcoder submit region begin(Prohibit modification and deletion)
package main
//import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 
 * @param A string字符串 
 * @return int整型
*/

func Max(x,y int) int{
    if x>y{
        return x
    }
    return y
}

func isDuichen(data string) bool{
    mid:=len(data)/2
    for i:=0;i<=mid;i++{
        if data[i] != data[len(data)-1-i]{
            return false
        }
    }
    return true
}

func getLongestPalindrome( A string ) int {

    // write code here

    //多层遍历，暴力穷举法 把所有搭配的子串进行对称判断，如果对称则更新当前maxL值
    if len(A) == 0{
        return 0
    }
    maxL:=1
    //for i:=0;i<len(A);i++{
    //
    //    for j:=i+1;j<len(A);j++{
    //        subStr:=A[i:j+1]
    //        isD:=isDuichen(subStr)
    //        if isD{
    //            maxL=Max(maxL,len(subStr))
    //        }
    //    }
    //}
    //动态规划法: 使用二维数组记录开始结束idx对应子串是否为回文，dp[start][end]=1|0[true|false]
    //递推公式：dp[i][j]=1,当i=j
    //dp[i][j]=0 (A[i]!=A[j])
    //当子串长度 ≤ 3（即 j - i <= 2）时：
    //
    //只要首尾相等 A[i] == A[j]
    //
    //那它 一定是回文
    //
    //不用再去看 dp[i+1][j-1]
    // dp[i][j]=dp[i+1][j-1] (A[i]==A[j]) 如果是1，更新maxL 其中 j-1>i+1,即 j-i>2

    dp:=make([][]int,len(A))
    for i := range dp {
        dp[i] = make([]int, len(A))
    }
    n:=len(A)
    // 反向遍历 i，确保 dp[i+1][j-1] 已经计算
    for i := n - 1; i >= 0; i-- {
        for j := i; j < n; j++ {
            if A[i] == A[j] {
                if j-i <= 2 {
                    dp[i][j] = 1
                } else {
                    dp[i][j] = dp[i+1][j-1]
                }

                if dp[i][j] == 1 {
                    maxL = Max(maxL, j-i+1)
                }
            }
            //不相等时，数组已经默认初始化值是0了，不继续写
        }
    }

    return maxL
}
//nowcoder submit region end(Prohibit modification and deletion)
