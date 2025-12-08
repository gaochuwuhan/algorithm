package nowcoder.editor.cn;
type ListNode struct {
	Val  int
	Next *ListNode
}

func Max(x,y int)int{
	if x>y{
		return x
	}
	return y
}