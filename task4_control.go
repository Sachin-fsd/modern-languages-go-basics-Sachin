package main

import "fmt"

func main() {

	// --- IF / ELSE ---
	// 
	// Basic conditional branching.
	age := 20
	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}

	// A common pattern in Go is to include a short statement
	// before the condition. 'num' is only scoped to this if-else block.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// --- SWITCH ---
	// 
	// 'switch' is a powerful multi-way conditional.
	day := "Wednesday"
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's a weekday.")
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("That's not a valid day.")
	}

	// Switch without an expression is an alternate way to write if/else chains.
	score := 85
	switch {
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80:
		fmt.Println("Grade: B")
	case score >= 70:
		fmt.Println("Grade: C")
	default:
		fmt.Println("Grade: F")
	}

	// --- FOR (Go's only loop) ---
	// 

	// 1. The C-style 'for' loop
	fmt.Println("C-style loop:")
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	// 2. 'for' as a 'while' loop
	fmt.Println("'while' style loop:")
	n := 1
	for n < 4 {
		fmt.Println(n)
		n *= 2
	}

	// 3. 'for' as an infinite loop (with a break)
	// for {
	// 	fmt.Println("Loop forever (until a break)")
	// 	break
	// }

	// --- RANGE ---
	// 
	// 'range' iterates over elements in various data structures.
	// Practical use: Iterating over a slice.
	fmt.Println("'range' over a slice:")
	nums := []int{10, 20, 30}
	sum := 0
	// 'range' returns both the index and the value.
	for index, value := range nums {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
		sum += value
	}
	fmt.Println("Total sum:", sum)

	// If you don't need the index, use the blank identifier '_'.
	for _, value := range nums {
		if value == 20 {
			fmt.Println("Found 20!")
		}
	}
}