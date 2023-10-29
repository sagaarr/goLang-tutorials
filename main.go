package main // this will tell go compiler that code should be compiled into executable program in the end (standalone executable file).

import "fmt" // package from go standard library. this one for formating strings and printing messages in console

// main is the default entry file name in golang when program executes it will automatically look for function main to start with execution.
// Each project will have a single entry point file with main.go and with main function in it. we can write other functions in it or import them	.

func main() {
	// Here you can see the first letter of the imported function is in capital which is a standard practice in go for library functions
	fmt.Println("Hello, Ninja!!")
}
