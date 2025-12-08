/**
给定一个无重复元素的整数数组nums，请你找出其中没有出现的最小的正整数 进阶： ;空间复杂度 ，时间复杂度 数据范围: 
 Related Topics 数组 哈希 
示例:
输入:[1,0,2]
输出:3 

*/
package nowcoder.editor.cn; //根据实际修改

//
//import ("sort"
//)


//nowcoder submit region begin(Prohibit modification and deletion)
package main

//import (
//
//    "sort"
//)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 
 * @param nums int整型一维数组 
 * @return int整型
*/
//方法一 排序后找到不连续的数字
//func minNumberDisappeared( nums []int ) int {
//    // write code here
//    sort.Ints(nums)
//    //先排序，然后找到第一个正数
//    if nums[len(nums)-1]<0{
//        return 1
//    }
//    pre:=0
//    for i:=0;i<len(nums);i++{
//        if nums[i]<=0{
//            continue
//        }
//        if pre == 0&&nums[i]!=1{ //第一个数不为1，那最小没出现的必为1
//            return 1
//        }
//        //对于 1，2，3，5 缺的就是4
//        if nums[i] != pre+1{
//            return nums[i]-1 //num=5,pre=3
//        }
//        pre = nums[i]
//    }
//
//    return pre+1
//
//}


func minNumberDisappeared( nums []int ) int {
    // write code here
    //缺的可能是：
    //1：当里面没有1时
    //len(num)，当正数从1开始顺序递增时
    //1~len(num)的数字：当有1 -x，但是中间少了一个连续得数
    //可以把所有在1-n 放到0-n-1的索引上，放完重新遍历一下，num[i]如果不==i+1，那么缺失的值就是i+1

    n := len(nums)

    // 1. 原地哈希：把值 v 放到索引 v-1
    for i := 0; i < n; i++ {
        for {
            v := nums[i]
            // 只处理 1..n 范围内的值
            if v <= 0 || v > n || nums[v-1] == v {
                break
            }
            // 交换 nums[i] 与 nums[v-1]
            nums[i], nums[v-1] = nums[v-1], nums[i]
        }
    }

    // 2. 第一处 nums[i]≠i+1 即为缺失的最小正整数
    for i := 0; i < n; i++ {
        if nums[i] != i+1 {
            return i + 1
        }
    }

    // 3. 1..n 都出现，返回 n+1
    return n + 1

}
func Max(x,y int)int{
    if x>y{
        return x
    }
    return y
}
//nowcoder submit region end(Prohibit modification and deletion)
