package main

import (
	"bufio"
	"fmt"
	"github.com/csos95/Go-Packages/calculator"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var calc = calculator.NewCalculator(20)
	var result float64
	var input string
	var err error

	fmt.Println("This is a polish notation calculator.\nTo quit, enter 'quit' as an equation")
	for {
		fmt.Print("Enter an equation: ")
		input, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		} else if input == "quit\n" {
			return
		} else {
			result, err = calc.Calculate(input)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Result: %f\n", result)
			}
		}
	}
}
