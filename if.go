package main

import "fmt"

func main() {

	//case 1--------------------------------------------------------------------------------------------------
	// number1 := 17 //commented away for preceeding
	// number2 := 24

	// resultMessage := "the numbers are the same"

	// Insert your code here
	// Hint: You may wish to make use of strconv.Itoa to convert int to string

	// if number1, number2 := 17, 24; number1 < number2 {
	// 	fmt.Printf("%v is bigger than %v\n", number2, number1)
	// } else if number1 > number2 {
	// 	resultMessage = "number1 is bigger than number2" //local variable for resultMessage change does not affected global variable value
	// 	fmt.Println(resultMessage)
	// } else {
	// 	fmt.Println(resultMessage)
	// }

	// working code for case 1

	//--------------------------------------------------case 2---------------------------------------------------------

	// if number1, number2 := 17, 24; number1 < number2 {
	// 	fmt.Printf("%v is bigger than %v\n", number2, number1)
	// } else if number1 > number2 {
	// 	fmt.Printf("%v is bigger than %v\n", number1, number2)
	// } else {
	// 	fmt.Printf("%v is equal to %v", number1, number2)
	// }

	// working code for case 2

	//-------------------------------------------------case 3----------------------------------------------------------

	var inputNumber int
	fmt.Printf("Input any desired number: ")
	fmt.Scan(&inputNumber)

	if inputNumber%2 == 0 {
		fmt.Printf("%v is an even number\n", inputNumber)
	} else {
		fmt.Printf("%v is an odd number\n", inputNumber)
	}

	if number1, number2 := 17, 24; inputNumber > number1 && inputNumber > number2 {
		fmt.Printf("%v is the biggest number in the group\n", inputNumber)
		number3 := inputNumber - number2
		fmt.Printf("%v is bigger than the 2nd biggest number by %v", inputNumber, number3)

	} else if inputNumber < number2 && inputNumber > number1 {
		fmt.Printf("%v is the 2nd biggest number in the group\n", inputNumber)
		number3 := number2 - inputNumber
		fmt.Printf("%v is smaller than the biggest number by %v", inputNumber, number3)

	} else if inputNumber == number2 {
		fmt.Printf("the number is same as the biggest number in the group, %v", number2)

	} else if inputNumber == number1 {
		fmt.Printf("the number is same as the smallest number in the group, %v", number1)

	} else {
		fmt.Printf("%v is the smallest number in the group\n", inputNumber)
		number3 := number1 - inputNumber
		fmt.Printf("%v is smaller than the 2nd smallest number by %v", inputNumber, number3)

	}

}
