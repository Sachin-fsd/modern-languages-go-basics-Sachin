package main

import (
	"fmt"
	// Import your custom package using its module path.
	// You may need to run 'go mod init task5_project'
	// and 'go mod tidy' in the 'task5_project' directory
	// for this to work correctly.
	"task5_project/custommath"
)

// This is a custom function within the 'main' package.
// It takes two integers as parameters and returns one integer.
func multiply(x, y int) int {
	return x * y
}

func main() {
	// --- Using a custom function from the same package ---
	product := multiply(5, 5)
	fmt.Printf("Result of local 'multiply' function: %d\n", product)

	// --- Using an imported function from a custom package ---
	// 
	sum := custommath.Add(10, 20)
	fmt.Printf("Result of imported 'custommath.Add' function: %d\n", sum)

	// This line would cause a compile error because 'subtract' is not exported.
	// diff := custommath.subtract(10, 5)
}