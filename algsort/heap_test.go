package algsort

import "testing"

//第一个非叶子结点 arr.length/2-1=5/2-1=1
func TestHeap(t *testing.T) {
	sortRet := Heap([]int{1, 2, 3})
	t.Log(sortRet)
}
