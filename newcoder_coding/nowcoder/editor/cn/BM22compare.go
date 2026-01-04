/**
牛客项目发布项目版本时会有版本号，比如1.02.11，2.14.4等等 现在给你2个版本号version1和version2，请你比较他们的大小 版本号是由修订
号组成，修订号与修订号之间由一个"."连接。1个修订号可能有多位数字组成，修订号可能包含前导0，且是合法的。例如，1.02.11，0.1，0.2都是合法的版本号
 每个版本号至少包含1个修订号。 修订号从左到右编号，下标从0开始，最左边的修订号下标为0，下一个修订号下标为1，以此类推。 
 比较规则： 一. 比较版本号时，请按从左到右的顺序依次比较它们的修订号。比较修订号时，只需比较忽略任何前导零后的整数值。比如"0.1"和"0.01"的版本号是
相等的 二. 如果版本号没有指定某个下标处的修订号，则该修订号视为0。例如，"1.1"的版本号小于"1.1.1"。因为"1.1"的版本号相当于"1.1.0"，第
3位修订号的下标为0，小于1 三. version1 > version2 返回1，如果 version1 < version2 返回-1，不然返回0. 
 数据范围： version1 和 version2 的修订号不会超过int的表达范围，即不超过 32 位整数 的范围 
 进阶： 空间复杂度 ， 时间复杂度 
 
 Related Topics 字符串 双指针 
示例:
输入:"1.1","2.1"
输出:-1 

*/
package nowcoder

import (
    "strconv"
    "strings"
)

.editor.cn; //根据实际修改
//nowcoder submit region begin(Prohibit modification and deletion)
package main
//import "fmt"
import (
"strconv"
"strings"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 比较版本号
 * @param version1 string字符串 
 * @param version2 string字符串 
 * @return int整型
*/
func compare( version1 string ,  version2 string ) int {
    // write code here
    v1I,v2I:=0,0
    v1Nums:=strings.Split(version1,".")
    v2Nums:=strings.Split(version2,".")
    for v1I<len(v1Nums) && v2I<len(v2Nums){
        numStr1:=v1Nums[v1I]
        numStr2:=v2Nums[v2I]
        version1Num:=rmPrezero(numStr1)
        version2Num:=rmPrezero(numStr2)
        if version1Num>version2Num{
            return 1
        }
        if version1Num<version2Num{
            return -1
        }
        v1I++
        v2I++
    }
    if v1I<len(v1Nums){
        for v1I<len(v1Nums){
            numStr1:=v1Nums[v1I]
            version1Num:=rmPrezero(numStr1)
            if version1Num!=0{
                return 1
            }
            v1I++
        }
    }
    if v2I<len(v2Nums){
        for v2I<len(v2Nums){
            numStr2:=v2Nums[v2I]
            version2Num:=rmPrezero(numStr2)
            if version2Num!=0{
                return 1
            }
            v2I++
        }
    }

    return 0
}

func rmPrezero(str string) int{
    if str=="0"{
        return 0
    }
    idx:=0
    for i,v:=range str{
        if v!='0'{
            idx=i
            break
        }
    }
    retStr:=str[idx:]
    ret,_:=strconv.Atoi(retStr)
    return ret

}
//nowcoder submit region end(Prohibit modification and deletion)
