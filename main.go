package main

import "fmt"

func main() {
	// Here we declared array we need to declare the array length [3] and also the type that is going to be in it "int" and then equal sign and again
	// Declare the the length and type with array elements in the curly braces {} once the array length is delcared it cannot be modified in golang

	// var ages [3]int = [3]int{20,25,30}/

	// /Short hand method to declare the array
	var ages = [3]int{20, 25, 30}

	name := [4]string{"sagar", "sachin", "chinchorkar", "kumar"}
	name[1] = "swapna"
	fmt.Println(ages, len(ages))
	fmt.Println(name, len(name))

	// Slices (use arrays under the hood)

	var scores = []int{100, 20, 52}
	scores[2] = 25
	scores = append(scores, 88)
	fmt.Println(scores, len(scores))

	// slice ranges
	// /Here the rangeOne will include elements from position 1 to position 2 of the array
	rangeOne := name[1:3]
	// Here the rangeTwo will include elements from element 2 onwards
	rangeTwo := name[2:]
	//  Here the rangeThree will include the elements from 0 position till the 2nd position of the array
	rangeThree := name[:3]
	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "sagar")
	fmt.Println(rangeOne)
}
