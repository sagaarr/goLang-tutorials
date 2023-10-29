package main

import "fmt"

func main() {
	age := 28
	name := "sagar"
	//  -----------------Print-----------------
	// fmt.Print is similar to Println but it will not print on a new line to do that we need to specify that in the string it self
	fmt.Print("Hello, ")
	fmt.Print("world! \n") // <--- Here \n will tell the compiler that any Print statement after this one needs to be printed on new line.
	fmt.Print("new line! \n")

	// -----------------Println-----------------
	fmt.Println("Hello sagar!")
	fmt.Println("ByeBye sagar..")
	fmt.Println("my age is", age, "and my name is", name) // Prints the string with age and name in it (variables in it).

	// -----------------Printf (formatted strings)-----------------
	fmt.Printf("my age is %v and my name is %v \n", age, name) // prints everything we pass

	fmt.Printf("my name is %q and my age is %q \n", name, age) // prints only strings

	fmt.Printf("age is of type %T \n", age) // prints type of the variable

	fmt.Printf("you scored %f points! \n", 225.55)    // prints floating points
	fmt.Printf("you scored %0.1f points! \n", 225.55) // prints round-off floating points

	// -----------------Sprintf (save formatted strings)-----------------
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name) // Saves the string in variable rather than printing it directly
	println("saved string is ::", str)

}
