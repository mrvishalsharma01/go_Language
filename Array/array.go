package main

import (
	"fmt"
)

func main() {

	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}
	arr3 := []int{1, 2, 3}

	arr4 := [3]int{}
	arr5 := [5]int{1: 20, 2: 50}

	arr6 := [...]string{"apple", "banana", "cherry"}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arr5)
	fmt.Println(arr6)
	fmt.Println("Length of arr1:", len(arr6))
}
