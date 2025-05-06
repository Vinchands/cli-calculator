package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

type Operation struct {
	action string
	value1 float32
	value2 float32
}

// Console specific functions
func clearConsole() {
	var cmd *exec.Cmd
	
	switch runtime.GOOS {
		case "windows": cmd = exec.Command("cmd", "/c", "cls")
		case "linux", "darwin": cmd = exec.Command("clear")
		default:
			fmt.Println("Unsupported OS")
			return
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

// App specific functions
func performOperation(operation Operation) float32 {
	
	var (
		x = operation.value1
		y = operation.value2
	)
	
	switch operation.action {
		case "sum": return x + y
		case "sub": return x - y
		case "mul": return x*y
		case "div": return x/y
		default:
			fmt.Println("Error!")
			time.Sleep(2 * time.Second)
			return 0
	}
}

func showOutput(operation Operation) {
	clearConsole()
	
	fmt.Printf("Result: %f\n\n", performOperation(operation))
	
	fmt.Print("Enter to continue...")
	fmt.Scanln()
	main()
}

func handleInputs(action string) {
	var (
		value1 float32
		value2 float32
	)
	
	fmt.Print("Value 1: ")
	fmt.Scanln(&value1)
	fmt.Print("Value 2: ")
	fmt.Scanln(&value2)
	
	operation := Operation{
		action: action,
		value1: value1,
		value2: value2,
	}
	
	showOutput(operation)
}

func handleSelection(selection string) {
	switch selection {
		case "1": handleInputs("sum")
		case "2": handleInputs("sub")
		case "3": handleInputs("mul")
		case "4": handleInputs("div")
		case "5":
			fmt.Println("Bye!")
			time.Sleep(2 * time.Second)
			clearConsole()
			os.Exit(0)
		default:
			fmt.Println("Invalid option!")
			time.Sleep(1 * time.Second)
			main()
	}
}

func showMenu() {
	var selection string
	
	clearConsole()
	fmt.Println("CALCULATOR APP - MENU")
	fmt.Println("1. (+) Addition      ")
	fmt.Println("2. (-) Substraction  ")
	fmt.Println("3. (*) Multiplication")
	fmt.Println("4. (/) Division      ")
	fmt.Println("5. (x) Exit          ")
	fmt.Println("---------------------")
	fmt.Print("Select: ")
	fmt.Scanln(&selection)
	
	handleSelection(selection)
}

func main() {
	showMenu()
}
