package main

import (
	"fmt"

	"github.com/mutexlock/golang/algsort"
)

func main() {
	fmt.Println(algsort.Heap([]int{1, 3, 4}))
	fmt.Println("good bye")
}
