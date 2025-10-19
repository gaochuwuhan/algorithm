
//IMPORTANT!! Submit Code Region Begin(Do not remove this line)
func isPalindrome(x int) bool {
	text:=strconv.Itoa(x)
	idxEnd:=len(text)-1
 	for i:=0;i<=idxEnd/2;i++{
		 if text[i] != text[idxEnd-i]{
			 return false
		}
	}
	return true
}
//IMPORTANT!! Submit Code Region End(Do not remove this line)