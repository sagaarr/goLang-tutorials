package main

func main() {

	// // strings (strings are in double quotes not in single quotes in go)
	// var nameOne string = "sagar"
	// var nameTwo = "chinchorkar"
	// var nameThree string
	// fmt.Println(nameOne, nameTwo, nameThree)

	// // we cannot change the type after the variables are declared
	// nameOne = "Peach"
	// nameTwo = "browser"
	// fmt.Println(nameOne, nameTwo)

	// // this is a short hand for declaring a variable.  (Cannot use this short hand version outside the function)
	// nameour := "sagar"
	// fmt.Println(nameour)

	// ints

	// var ageOne int = 20
	// var ageTwo = 30
	// ageThree := 40
	// fmt.Println(ageOne, ageTwo, ageThree)

	// bits and memory
	// Numeric type range for each bit size ie; 8, 16, 32, 64 and so on can be found on
	//  https://go.dev/ref/spec#Numeric_types
	// var numOne int8 = 25
	// var numTwo int8 = -128
	// var numThree uint = 25 // cannot declare negative int.
	// var numFour uint8 = 255 // uint8 ranges from 0-255

	var scoreOne float32 = 25.90
	var scoreTwo float64 = 156465465 // have a higher precision than float32 as well.
	scoreThree := 1.5                // in shorthand as well the default float considered is float64.

}
