package main

import "fmt"



/*
	** Go Datatype **
	- Basics:
		-- bool
		-- Numeric
		-- string
	
	- Numeric:

		-- int
		-- int8
		-- int16
		-- int32 (rune) # built-in alias
		-- int64

		-- uint
		-- uint8(byte) # built-in alias
		-- uint16
		-- uint32
		-- uint64

		-- uintptr

		-- float
		-- float32
		-- float64

		-- complex64
		-- complex128

*/
func initialization_demo(){

	// **Initialization of variables**


	// strings
	var name string = "Sanjit Khanal"
	// var variable_name variable_type = value ..... this is one way of declaration
	
	/*
		Errors:
		
		go is statically typed language and is type safe so the following will not work
		
		var name = "<NAME>"
		var name string = 12
		var name string = '<NAME>'
		
		also any unused variable will throw an error
	*/
	
	/* 
		Note to self :

		In Go, the **println** function is a built-in function that provides a simple way to print to the standard output. 
		It's part of the Go runtime and is available without importing any packages. However, it is intended primarily 
		for debugging and should not be used in production code.
	
	*/
	
	var age = 20 
	// var var_name = value ... this is another type of declaration which is called a short hand declaration
	
	var city string
	// var var_name ... this way we can initialize the variable to a default value
	
	mother_name := "Kamala"
	// This is a short-hand declaration of variables that is limited to inside a function

	fmt.Println("Hey I am", name, "and I am ", age," years old. I live in ",city)
	fmt.Println(mother_name)
}

func modification_demo(){

	// **Modification of variables**

	// strings
	var name string = "<NAME>"
	var age string 

	name = "Sanjit"
	// this is ok as the ney value being assigned to the variable confirms with its type constraints
	// var name string = 12 will throw an error as 12 is not a string

	age = "22"
	// Here comes the use of declared vars, which can be declared first and then initialized later

	fmt.Println("Hello I am ", name, ". I am ", age ," years old")
}

func numbers_demo(){
	var ageone int  = 20
	var agetwo = 30
	agethree := 40 

	fmt.Println(ageone, agetwo, agethree)

	/* 
		types of numbers in Go
		int 
		int64

	*/
}

func main() {

	initialization_demo()
	modification_demo()
	numbers_demo()
}