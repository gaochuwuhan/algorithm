
//IMPORTANT!! Submit Code Region Begin(Do not remove this line)
func longestCommonPrefix(strs []string) string {
commPrefix := ""
firstS := strs[0]
//firstSlen:=len(firstS)
hasRuneF := true
for fi, runeF := range firstS {
//遍历其余字符串的字符
if !hasRuneF {
break
}
for i := 0; i < len(strs); i++ {
strI := strs[i]
if fi <= len(strI)-1 {
if string(strI[fi]) != string(runeF) {
hasRuneF = false
break
}
}else{
hasRuneF = false
break
}
//判断是最后一个字符串，说明都有前缀
if i == len(strs)-1 {
commPrefix = commPrefix + string(runeF)
}
}
}
return commPrefix
}
//IMPORTANT!! Submit Code Region End(Do not remove this line)