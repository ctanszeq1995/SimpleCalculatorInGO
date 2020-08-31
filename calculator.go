package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

type calculation struct {
	value int
}

func main() {
	initValue := calculation {0}
	fmt.Println("This is a simple calculator!")
	fmt.Println("The initial value is 0.")
	fmt.Println("First, enter \"+\", \"-\", \"*\", or \"/\"")
	fmt.Println("Then enter the value for calculation. Only integer values are allowed.")
	fmt.Println("To exit, enter \"exit\".")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		symbol := ""
		fmt.Println("Please enter the symbol: ")
		if scanner.Scan() {
			symbol = scanner.Text()
		}
		switch symbol {
			case "+", "-", "*", "/":
				valid, exit := initValue.calculate(symbol)
				if exit {
					fmt.Println(fmt.Sprintf("The final value is: %d", initValue.value))
					os.Exit(0)
				}
				
				if !valid{
					fmt.Println("Value entered is not an integer!")
					fmt.Println(fmt.Sprintf("The final value is: %d", initValue.value))
					os.Exit(1)
				}
				
				fmt.Println(fmt.Sprintf("Current value is: %d", initValue.value))
			case "exit":
				fmt.Println(fmt.Sprintf("The final value is: %d", initValue.value))
				os.Exit(0)
			default:
				fmt.Println(fmt.Sprintf("\"%s\" is not a valid symbol!", symbol))
				os.Exit(1)
		}
	}
}

func (v *calculation)calculate(symbol string) (validValue bool, exit bool) {
	inputValueString := ""
	validValue = false
	fmt.Println("Please enter the value: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		inputValueString = scanner.Text()
	}
	
	if inputValueString == "exit" {
		return validValue, true
	}
	
	inputValue, err := strconv.Atoi(inputValueString)
	
	// If the input value cannot be converted to int, then is invalid input
	if err != nil {
		validValue = false
		return validValue, false
	}
	
	switch symbol {
		case "+":
			v.value = v.value + inputValue
			validValue = true
		case "-":
			v.value = v.value - inputValue
			validValue = true
		case "*":
			v.value = v.value * inputValue
			validValue = true
		case "/":
			v.value = v.value / inputValue
			validValue = true
		default:
			return validValue, false
	}
	
	return validValue, false
}