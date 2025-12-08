/**
给你一个大小为 n ;的字符串数组 strs ，其中包含n个字符串 , 编写一个函数来查找字符串数组中的最长公共前缀，返回这个公共前缀。 数据范围： ， 进阶：
空间复杂度 ，时间复杂度 
 Related Topics 字符串 
示例:
输入:["abca","abc","abca","abc","abcc"]
输出:"abc"

*/
package nowcoder


.editor.cn; //根据实际修改
//nowcoder submit region begin(Prohibit modification and deletion)
package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 
 * @param strs string字符串一维数组 
 * @return string字符串
*/
func longestCommonPrefix( strs []string ) string {
    // write code here
    if len(strs)==0{
        return ""

    }
    if len(strs) == 1{
        return strs[0]
    }
    comp:=strs[0]
    pref:=""
    for i:=1;i<len(strs);i++{
        tgt:=strs[i]
        var j int
        if len(comp)< len(tgt){
            j=len(comp)
        }else{
            j=len(tgt)
        }
        pre:=make([]byte,0)
        for k:=0;k<j;k++{
            if comp[k] != tgt[k]{
                break
            }
            pre = append(pre,comp[k])
        }
        //一开始写错了，直接把被比较的comp替换成公共子串就行
        pref = string(pre)
        comp = pref

    }
    return pref

}
//nowcoder submit region end(Prohibit modification and deletion)
