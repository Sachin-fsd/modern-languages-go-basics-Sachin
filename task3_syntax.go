package main

import (
	"fmt"
	"strconv"
)

func main() {
	// --- VARIABLE DECLARATIONS ---

	// 'var' declares a variable. It can be package-level or function-level.
	// It's initialized to its "zero value" (e.g., 0 for int, "" for string).
	var age int
	fmt.Println("Default 'age' (zero value):", age)
	age = 30
	fmt.Println("Assigned 'age':", age)

	// You can declare and initialize in one line.
	var name string = "Alice"
	fmt.Println("Initialized 'name':", name)

	// Go can infer the type if an initializer is present.
	var isStudent = true
	fmt.Println("Inferred 'isStudent' type (bool):", isStudent)

	// Inside functions, the short assignment operator ':=' is common
	// for declaring and initializing.
	location := "Wonderland"
	fmt.Println("Short assignment 'location':", location)

	// --- CONSTANTS ---
	// 
	// Constants are declared with 'const' and cannot be changed.
	const pi = 3.14159
	// pi = 3.14 // This would cause a compile error

	// Constants can be numbers, strings, or booleans.
	const truth = true
	fmt.Println("Constant 'pi':", pi)

	// --- BASIC DATA TYPES ---
	// 

	var myInt int = 100
	var myFloat float64 = 45.12
	var myString string = "Hello GO"
	var myBool bool = false
	var myComplex complex128 = 5 + 5i

	fmt.Printf("Type: %T, Value: %v\n", myInt, myInt)
	fmt.Printf("Type: %T, Value: %v\n", myFloat, myFloat)
	fmt.Printf("Type: %T, Value: %v\n", myString, myString)
	fmt.Printf("Type: %T, Value: %v\n", myBool, myBool)
	fmt.Printf("Type: %T, Value: %v\n", myComplex, myComplex)

	// --- TYPE CONVERSIONS ---
	// 
	// Go requires explicit type conversions; there are no implicit conversions.
	var i int = 42
	var f float64 = float64(i) // Explicit conversion from int to float64
	var u uint = uint(f)       // Explicit conversion from float64 to uint

	fmt.Printf("Int %d converted to Float64: %f\n", i, f)
	fmt.Printf("Float64 %f converted to Uint: %d\n", f, u)

	// You can also convert between string and numbers using the 'strconv' package.
	numString := "123"
	numInt, err := strconv.Atoi(numString) // Atoi (string to int)
	if err != nil {
		fmt.Println("Error converting string to int")
	} else {
		fmt.Printf("String '%s' converted to Int: %d\n", numString, numInt)
	}

	intToConvert := 456
	stringFromInt := strconv.Itoa(intToConvert) // Itoa (int to string)
	fmt.Printf("Int %d converted to String: '%s'\n", intToConvert, stringFromInt)
}