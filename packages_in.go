package main
// A go program is composed of many packages
// if a script belong to main pacakge then it isd compiled into executable when compiling 

import "fmt"
//import statement for formatted input/output

func main(){

	// All funcitons start with func keyword
	// func main is the entry point of the program i.e it is fired automatically 

	/*
	
		In Go, the limitations and rules for the main function are as follows:

		**Single main package per executable:**

			--An entire Go project that compiles into a single executable can only have one main package. The main package is special in Go because it signifies the entry point of the application.
			
		**Single main function in the main package:**

			Within the main package, there must be exactly one main function. This function serves as the starting point when the program is run.
			
		**Multiple files in the main package:**

			You can have multiple files in the main package, but they must collectively define a single main package. Only one of these files should contain the main function, though they can contain other functions, types, and variables that the main function can use.
			
		**Different packages:**

			Other packages in your project should not have a main function. They can be named anything other than main and can be used to organize your code into reusable modules. These packages can be imported and used by the main package or by other packages.
		
	*/

	fmt.Println("Hello! World")

}