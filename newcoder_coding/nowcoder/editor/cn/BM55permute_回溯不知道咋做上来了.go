/**
给出一组数字，返回该组数字的所有排列 例如： [1,2,3]的所有排列如下
 [1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2], [3,2,1].
 （以数字在数组中的位置靠前为优先级，按字典序排列输出。） 
 数据范围：数字个数 要求：空间复杂度 ，时间复杂度 
 Related Topics 递归 
示例:
输入:[1,2,3]
输出:[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

*/
package nowcoder.editor.cn; //根据实际修改



//nowcoder submit region begin(Prohibit modification and deletion)
package main


/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 
 * @param num int整型一维数组 
 * @return int整型二维数组
*/
func permute( num []int ) [][]int {
    // write code here
    ret:=make([][]int,0)
    ans:=make([]int,0)
    find(ans,num,&ret)
    return ret
}

func find(ans []int,num []int,ret *[][]int){
    for i:=0;i<len(num);i++{
        if sliceContain(ans,num[i]){
            continue
        }
        ans = append(ans,num[i])
        if len(ans)==len(num){
            *ret = append(*ret,ans)
        }
        find(ans,num,ret)
        ans = ans[0:len(ans)-1]//回溯

    }
}

func sliceContain(num []int, data int) bool{
    for _,v:=range num{
        if v==data{
            return true
        }
    }
    return false
}
//nowcoder submit region end(Prohibit modification and deletion)
