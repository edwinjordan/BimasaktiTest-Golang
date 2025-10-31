package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("==========================================")
	fmt.Println("   GOLANG EXAMPLES - BimasaktiTest")
	fmt.Println("==========================================")
	fmt.Println()

	// Check if user wants to run specific example
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "1":
			typeVerificationExample()
		case "2":
			functionClosureExample()
		case "3":
			classInheritanceExample()
		case "4":
			permutationExample()
		case "5":
			stopGoroutineExample()
		default:
			fmt.Println("Usage: go run . [1|2|3|4|5]")
			fmt.Println("  1 - Type Verification")
			fmt.Println("  2 - Function Closure")
			fmt.Println("  3 - Class Inheritance")
			fmt.Println("  4 - Permutation")
			fmt.Println("  5 - Stop Goroutine")
			fmt.Println("\nRunning all examples...")
			fmt.Println()
			runAllExamples()
		}
	} else {
		runAllExamples()
	}

	fmt.Println("==========================================")
	fmt.Println("   All examples completed successfully!")
	fmt.Println("==========================================")
}

func runAllExamples() {
	typeVerificationExample()
	functionClosureExample()
	classInheritanceExample()
	permutationExample()
	stopGoroutineExample()
}
