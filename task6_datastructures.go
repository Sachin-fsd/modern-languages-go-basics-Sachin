package main

import "fmt"

func main() {

	// --- ARRAYS ---
	// 
	// Arrays have a fixed size, which is part of their type.
	var fruits [3]string
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fmt.Println("Array (fixed size):", fruits)
	fmt.Println("Array length:", len(fruits))

	// Declare and initialize
	primes := [5]int{2, 3, 5, 7, 11}
	fmt.Println("Array (initialized):", primes)

	// --- SLICES ---
	// 
	// Slices are a more powerful, flexible view of an underlying array.
	// They do not have a fixed size.
	
	// Creating a slice (this creates an underlying array)
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", numbers)

	// Slicing a slice (or array) 
	// 'subSlice' points to the same underlying data as 'numbers'
	subSlice := numbers[1:4] // Includes elements 1, 2, 3 (index 4 is excluded)
	fmt.Println("Sub-slice [1:4]:", subSlice)

	// Append: Adds elements to a slice. May return a new slice
	// if the underlying array needs to be re-allocated. 
	numbers = append(numbers, 6, 7)
	fmt.Println("Slice after 'append':", numbers)

	// Copy: Safely copies elements from one slice to another. 
	src := []int{10, 20, 30}
	dest := make([]int, len(src)) // Create a slice of the same length
	copiedCount := copy(dest, src)
	fmt.Println("Copied 'dest' slice:", dest)
	fmt.Println("Number of items copied:", copiedCount)
	
	// Modifying 'dest' does not affect 'src'
	dest[0] = 99
	fmt.Println("Modified 'dest':", dest)
	fmt.Println("Original 'src':", src)

	// --- MAPS ---
	// 
	// Maps are Go's built-in hash table type (key-value pairs).

	// Create a map using 'make'
	// 'make(map[key-type]value-type)'
	studentGrades := make(map[string]int)

	// Add key-value pairs
	studentGrades["Alice"] = 92
	studentGrades["Bob"] = 78
	studentGrades["Charlie"] = 88
	fmt.Println("Map (studentGrades):", studentGrades)
	fmt.Println("Bob's grade:", studentGrades["Bob"])

	// Delete a key-value pair
	delete(studentGrades, "Bob")
	fmt.Println("Map after deleting Bob:", studentGrades)

	// Check if a key exists
	// Accessing a key returns two values: the value and a boolean 'ok'
	grade, ok := studentGrades["Charlie"]
	fmt.Printf("Charlie's grade: %d, Exists: %t\n", grade, ok)

	grade, ok = studentGrades["Bob"] // Bob was deleted
	fmt.Printf("Bob's grade: %d, Exists: %t\n", grade, ok)
}