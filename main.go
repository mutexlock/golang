package main

import (
	"fmt"

	"github.com/mutexlock/golang/algsort"
)

func main() {
	h := []int{1, 3, 4, 34, 9}
	algsort.HeapSort(h)
	fmt.Println(h)
	fmt.Println("good bye")
}
