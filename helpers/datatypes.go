package helpers

import "fmt"

type Points struct {
	X, Y int
}

func ArrayTest() {
	var array1 = [5]int{2, 4}
	var array2 = array1[:]
	array2[1] = 500
	fmt.Println(array1, array2)

}

//This selects a half-open range which includes the first element, but excludes the last one.

func SliceTest() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	d := s[:]
	s = s[2:]
	printSlice(s)

	s = s[:]
	s[1] = 8
	printSlice(s)
	printSlice(d)
	// new slice

	newSlice := []int{}
	printSlice(newSlice)
	fmt.Println(newSlice)
	newSlice = append(newSlice, 400, 500)

	// append function slice test
	newSlice = []int{1, 3, 4, 4}
	newSliceCopy := newSlice[:]
	//append will return  a slice withr eference to new array
	newSlice = append(newSlice, 400, 500)
	newSliceCopy[0] = 100
	fmt.Println(newSlice, newSliceCopy)
	//make slice with make function
	a := make([]int, 5)
	printSlice(a)

	// if we ask slice for differnt capacity levels
	newSlice = []int{1, 2, 3, 4, 5, 6}
	newSliceCopy = newSlice[:]
	//capacity is less than what x slice needs
	x := newSliceCopy[:5]
	fmt.Println(x)
}

func RangeFunction() {
	id := 10
	fmt.Println("id - %d", id)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
