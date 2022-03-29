package main

import "fmt"

func main() {
	// number1 := 17 //commented away for preceeding
	// number2 := 24
	// resultMessage := "No outcome."

	//Insert your code here
	//Hint: You may wish to make use of strconv.Itoa to convert int to string

	// if number1, number2 := 17, 24; number1 < number2 {
	// 	fmt.Printf("%v is bigger than %v\n", number2, number1)
	// } else if number1 > number2 {
	// 	fmt.Printf("%v is bigger than %v\n", number1, number2)
	// } else {
	// 	fmt.Println("the numbers are the same")
	// }
	//working code for part 1

	var inputNumber int
	fmt.Printf("Input any desired number: ")
	fmt.Scan(&inputNumber)

	if number1, number2 := 17, 24; inputNumber > number1 && inputNumber > number2 {
		fmt.Printf("%v is the biggest number in the group\n", inputNumber)
	} else if inputNumber > number1 {
		fmt.Printf("%v is the 2nd biggest number in the group\n", inputNumber)
	} else {
		fmt.Printf("%v is the smallest number in the group", inputNumber)
	}

}
