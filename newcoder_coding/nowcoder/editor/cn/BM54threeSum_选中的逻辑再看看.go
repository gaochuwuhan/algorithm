/**
给出一个有n个元素的数组S，S中是否有元素a,b,c满足a+b+c=0？找出数组S中所有满足条件的三元组。 
 数据范围：，数组中各个元素值满足 空间复杂度：，时间复杂度 
 注意：
 三元组（a、b、c）中的元素必须按非降序排列。（即a≤b≤c） 解集中不能包含重复的三元组。 例如，给定的数组 S = {-10 0 10 20 -10 -4
0},解集为(-10, -10, 20),(-10, 0, 10) 
 
 Related Topics 数组 双指针 排序 
示例:
输入:[0]
输出:[]

*/
package nowcoder

import "sort"

.editor.cn; //根据实际修改


import (
"sort"
"strconv")

//nowcoder submit region begin(Prohibit modification and deletion)
package main
import (
    "sort"
    "strconv")

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 
 * @param num int整型一维数组 
 * @return int整型二维数组
*/
func threeSum( num []int ) [][]int {
    // write code here

    if len(num) <3{
        return nil
    }

    sort.Ints(num)


    result:=make([][]int,0)
    m:=map[string]struct{}{}
    for i:=0;i<len(num)-2;i++{
        l:=i+1
        r:=len(num)-1
        for l<r{
            if num[i]+num[l]+num[r]<0{
                l++
                continue
            }
            if num[i]+num[l]+num[r]==0{
                //去重
                temp:=[]int{num[i],num[l],num[r]}
                sort.Ints(temp)
                str:=strconv.Itoa(temp[0])+"_"+strconv.Itoa(temp[1])+"_"+strconv.Itoa(temp[2])
                _,ok:=m[str]
                if !ok{
                    m[str] = struct{}{}
                    result = append(result, []int{num[i],num[l],num[r]})
                }
                //这里注意
                l++
                r--
            }
            if num[i]+num[l]+num[r] > 0{
                r--
            }
        }
    }
    return result

}
//nowcoder submit region end(Prohibit modification and deletion)
