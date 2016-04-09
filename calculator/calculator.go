/*
   Implements a calculator that uses two stacks to process
   equations written in polish notation.

   todo: add sqrt, pow, and maybe log
*/

package calculator

import (
	"fmt"
	"stack"
	"strconv"
	"strings"
)

//Calculator is a calculator made up of a stack for the Operators and the Numbers.
//To use it, pass a equation in polish notation to Calculator.Interpret
//and then call Calculator.Process to get a float64 result
type Calculator struct {
	Operators stack.Stack
	Numbers   stack.Stack
}

//NewCalculator returns a new calculator with stacks of the specified size
func NewCalculator(size int) Calculator {
	return Calculator{stack.NewStack(size), stack.NewStack(size)}
}

//Calculate takes an equation in polish notation and calls the intpret and process
//functions then returns the result.
func (c *Calculator) Calculate(equation string) (float64, error) {
	c.Clear()
	ok := c.Interpret(equation)
	if ok != nil {
		return 0, ok
	}
	return c.Process()
}

//Process processes the two stacks and returns the float64 result
func (c *Calculator) Process() (float64, error) {
	var num1, num2, result float64
	var inum1, inum2, ioperator interface{}
	var operator string
	if c.Operators.IsEmpty() || c.Numbers.IsEmpty() {
		return 0, fmt.Errorf("Error: one of the stacks is empty\nOperators %v\nNumbers %v\n", c.Operators.Items, c.Numbers.Items)
	} else if c.Operators.Top != c.Numbers.Top-1 {
		return 0, fmt.Errorf("Error: operator should have one less than numbers\nOperators %v\nNumbers %v\n", c.Operators.Items, c.Numbers.Items)
	}
	for {
		inum1, _ = c.Numbers.Pop()
		num1 = inum1.(float64)
		inum2, _ = c.Numbers.Pop()
		num2 = inum2.(float64)
		ioperator, _ = c.Operators.Pop()
		operator = ioperator.(string)

		switch operator {
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		case "*":
			result = num1 * num2
		case "/":
			result = num1 / num2
		}
		if c.Numbers.IsEmpty() {
			return result, nil
		}
		c.Numbers.Push(result)
	}
}

//Interpret breaks a string equation down into Operators and equations and places
//them in thier respective stacks
func (c *Calculator) Interpret(equation string) error {
	tokens := strings.FieldsFunc(equation, func(r rune) bool {
		return r == ' ' || r == '\n'
	})

	for _, token := range tokens {
		number, ok := strconv.Atoi(token)
		if ok == nil {
			ok := c.Numbers.Push(float64(number))
			if ok != nil {
				return fmt.Errorf("Error: numbers stack is full, create a calculator with a larger stack size")
			}
		} else if token == "+" || token == "-" || token == "*" || token == "/" {
			ok := c.Operators.Push(token)
			if ok != nil {
				return fmt.Errorf("Error: operators stack is full, create a calculator with a larger stack size")
			}
		} else {
			return fmt.Errorf("Error: unknown input entered: %v\n", token)
		}
	}
	c.Numbers.Reverse()
	return nil
}

//Clear clears the calculator stacks
func (c *Calculator) Clear() {
	c.Operators.Clear()
	c.Numbers.Clear()
}
