package main

import (
	"encoding/json"
	"fmt"
)

func arrays_() {
	// Declaration of arrays

	var ar_name [3]int = [3]int{20, 25, 30}
	fmt.Println("\n This is a simple array: ", ar_name)

	var arr_name2 = [3]int{1, 2, 3}
	// this is a shorthand declaration of array
	// go automatically infers type and we don't have to explicitly specify it
	arr_name2[1] = 4

	fmt.Printf("\n This is another array initialized through shorthand %v and of length %v \n", arr_name2, len(arr_name2))

	/*
		Just like arrays in C  the arrays of go have fixed length and one cannot append more characters than length of the arrays

		But unlike arrays slices are similar to arrays of Python as in they have dynamic length
	*/
	var scores = []int{100, 50, 60, 40, 80}
	scores = append(scores, 70)
	// this will append the value at the end of the array

	fmt.Println(scores, len(scores))
	// Slices are similar to arrays but they have dynamic length

	// We can also declare a slice using the following syntax
	var scores2 []int = make([]int, 0, 10)
	fmt.Println(scores2)
	/* Slice Ranges */
}

func print_encoding(arr []int) {
	arrj1, _ := json.Marshal(arr)
	fmt.Println(string(arrj1))
}

func slice_() {
	var s []int
	t := []int{}
	u := make([]int, 5)
	v := make([]int, 0, 5)

	fmt.Printf("\n %d, %d, %T, %5t, %#[3]v\n", len(s), cap(s), s, s == nil)
	fmt.Printf("\n %d, %d, %T, %5t, %#[3]v\n", len(t), cap(t), t, t == nil)
	fmt.Printf("\n %d, %d, %T, %5t, %#[3]v\n", len(u), cap(u), u, u == nil)
	fmt.Printf("\n %d, %d, %T, %5t, %#[3]v\n", len(v), cap(v), v, v == nil)

	print_encoding(s)
	print_encoding(t)
	print_encoding(u)
	print_encoding(v)
	/*
		s has 0 length, 0 capacity and null pointer |0|0|0| ----> null slice
		t has 0 length, 0 capacity and special pointer |0|0|struct{}|  ----> empty slice, it has a pointer to something under the hood
																		  it is a non null pointer but it is a kind of sentinel value
		u has 5 length, 5 capacity and pointer to contiguous mem of length 5 , the memories are filled with 0
		v has 0 length, 5 capacity and pointer to contiguous mem of length 5 , the memories are filled with garbage value or has space for values but no values
	*/

	/*
		Why does it matter?
		SLices (and maps) encode differently when they are null vs when they are empty
		s would be encoded as null where as t  would be encoded as empty.
		s ---> null
		t ---> []
		u ---> [0,0,0,0,0]
		v ---> []

		Also when checking for empty slices do  not use arr == null checker for very this reason
		a more proper choice would be to check if len(arr)	== 0 as it encompasses all the cases.

	*/

	/*
		why would append not work for u, where only length is described?

		when we declare a slice using make( []int, 5), the slice created has the following property
		the length is 5
		the capacity is 5
		the ptr is pointing to a contiguous memory location
		and the memories are filled with 0

		so if we now append to this array the capacity is overcome so a new array with updated memory address will
		be created and the value will be appended at the end that is at 6th position

	*/

	/*
		length v/s capacity

	*/
	a := [5]int{1, 2, 3,4,5}
	b := a[:1]
	c := b[0:2]
	b = append(b, 6)
	fmt.Println(b)
	fmt.Println(a)
	fmt.Printf("\n a has length %v and capacity %v", len(a), cap(a))
	fmt.Printf("\n b has length %v and capacity %v", len(b), cap(b))
	fmt.Printf("\n c has length %v and capacity %v\n", len(c), cap(c))

	/*
		Unintuitive operations in go
	*/
	// suppose a array

	arr1 := [5]int{1,2,3,4,5}

	//Now if i slice, say 2 ele form arr1 
	slice1 := arr1[:2]
	
	// What you are doing is defining a logical view into arr1 that is bounded by the rules specified in slice1
	// What it means for us is when you access slice1 you are viewing into chars of arr1 through slice1
	// so if you ask for the length and capacity of slice 1 
	// intuitively you would except it to be 2 and 2 , but it is actually 2, 5
	fmt.Printf("\nThe length of slice1 is %v and its capacity is %v", len(slice1), cap(slice1))

	// What it results in is that whe n you append to slice1 you are actually mutating arr1
	fmt.Println("\nThis is arr1 before slice1 uses append", arr1)
	slice1 = append(slice1, 6)
	fmt.Println("This is arr1 and slice1 after slice1 append operation", arr1, slice1)

	// This is an unintuitive oration allowed by go , for a more expected behavior describing the capacity of the 
	// slice during its declaration would be helpful
	arr2 := [5]int{1,2,3,4,5}
	slice2 := arr2[:2:2]
	fmt.Printf("\nThe length of slice2 is %v and its capacity is %v", len(slice2), cap(slice2))

	// Now when you append to slice2 you are exceeding its defined capacity and thus forces go to reallocate the memory for slice2
	fmt.Println("\nThis is arr2 before slice2 uses append", arr2)
	slice2 = append(slice2, 6)
	fmt.Println("This is arr2 and slice1 after slice2 append operation", arr2, slice2)

	// This is an unintuitive operation allowed by go , for a more expected behavior describing the capacity of the 
	// slice during its declaration would be helpful
}

func main() {

	fmt.Println("This is a simple script to learn concepts related to array and slices in GoLang")
	slice_()
}
