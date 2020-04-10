package algsort

import "testing"

func TestHeap(t *testing.T) {
	sortRet := Heap([]int{1, 2, 3})
	t.Log(sortRet)
}
