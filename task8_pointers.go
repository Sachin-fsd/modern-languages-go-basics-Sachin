package main

import "fmt"

// --- GARBAGE COLLECTION IN GO (Memory Management) ---
// 
// Go handles memory management automatically using a garbage collector (GC).
// You don't need to manually 'free' memory like in C/C++.
//
// 1. Allocation: When you create variables (e.g., using 'make', 'new', or
//    struct literals), Go allocates memory for them, either on the
//    stack (for short-lived, known-size variables) or the heap (for
//    variables that need to live longer, like slices, maps, or
//    when their address is taken).
//
// 2. Garbage Collection: A background process (the GC) runs periodically.
//    It uses a "mark-and-sweep" algorithm.
//    - Mark: It finds all objects that are still reachable (in-use) by
//      following pointers starting from "roots" (like global variables
//      and active function stacks).
//    - Sweep: It scans the heap and reclaims all memory from objects
//      that were *not* marked (i.e., they are no longer reachable/used).
//
// This process prevents memory leaks and makes development simpler.

// --- FUNCTION WITH POINTER PARAMETER ---
// 
// This function takes a pointer to an integer (*int).
// By using a pointer, it can modify the *original* variable
// that was passed in, rather than a copy.
func updateValue(val *int) {
	fmt.Printf("  (Inside function) Address 'val' points to: %p\n", val)
	// Dereference the pointer ('*val') to get the value,
	// then assign a new value to that memory location.
	*val = 100
	fmt.Println("  (Inside function) New value set to:", *val)
}

func main() {
	// --- POINTER USAGE ---
	// 
	
	// 'x' is a standard integer variable.
	x := 42
	fmt.Println("--- Basic Pointers ---")
	fmt.Println("Original value of x:", x)

	// The '&' operator gets the memory address of 'x'.
	// 'p' is a "pointer to an int" (type *int).
	p := &x
	fmt.Println("Memory address of x (value of p):", p)

	// The '*' operator "dereferences" a pointer.
	// It reads the value at the memory address the pointer holds.
	fmt.Println("Value at address p (dereferencing p):", *p)

	// You can also change the value of 'x' *through* the pointer.
	*p = 50
	fmt.Println("Value of x after modification via pointer:", x)

	// --- PASSING POINTERS TO FUNCTIONS ---
	// 
	y := 10
	fmt.Println("\n--- Pointers in Functions ---")
	fmt.Println("Original value of y:", y)
	fmt.Println("Memory address of y:", &y)

	// We pass the memory address of 'y' (using '&y') to the function.
	updateValue(&y)
	
	// 'y' has been changed by the function because we passed a pointer.
	fmt.Println("Final value of y (after function):", y)
}