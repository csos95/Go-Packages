package calculator

import (
	"fmt"
)

func ExampleCalculator_Interpret() {
	var calculator = NewCalculator(3)
	fmt.Println(calculator.Interpret("+ 1 2"))
	calculator.Clear()
	fmt.Println(calculator.Interpret("+ + 1 2 3"))
	calculator.Clear()
	fmt.Println(calculator.Interpret("+ + + 1 2 3 4"))
	calculator.Clear()
	fmt.Println(calculator.Interpret("+ + + + 1 2 3 4 5"))
	calculator.Clear()
	//Output:
	// <nil>
	// <nil>
	// Error: numbers stack is full, create a calculator with a larger stack size
	// Error: operators stack is full, create a calculator with a larger stack size
}

func ExampleCalculator_Process() {
	var calculator = NewCalculator(20)
	calculator.Interpret("+ 1 2")
	fmt.Println(calculator.Process())
	calculator.Clear()
	calculator.Interpret("- 2 3")
	fmt.Println(calculator.Process())
	calculator.Clear()
	calculator.Interpret("* 2 3")
	fmt.Println(calculator.Process())
	calculator.Clear()
	calculator.Interpret("/ 9 3")
	fmt.Println(calculator.Process())
	calculator.Clear()
	calculator.Interpret("- + / * 3 5 3 2 1")
	fmt.Println(calculator.Process())
	calculator.Clear()
	//Output:
	// 3 <nil>
	// -1 <nil>
	// 6 <nil>
	// 3 <nil>
	// 6 <nil>
}

func ExampleCalculator_Calculate() {
	var calculator = NewCalculator(20)
	fmt.Println(calculator.Calculate("- + / * 3 5 3 2 1"))
	fmt.Println(calculator.Calculate("* * * * * * * * * * * 1 2 3 4 5 6 7 8 9 10 11 12"))
	//Output:
	// 6 <nil>
	// 4.790016e+08 <nil>
}
