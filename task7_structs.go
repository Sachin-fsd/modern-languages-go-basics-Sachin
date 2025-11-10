package main

import "fmt"

// --- DEFINING A STRUCT ---
// 
// A 'struct' is a custom data type that groups related fields.
type Employee struct {
	FirstName string
	LastName  string
	ID        int
	IsActive  bool
}

// --- CREATING A METHOD FOR A STRUCT ---
// 
// A method is a function with a 'receiver' argument.
// This 'Display' method is associated with the 'Employee' struct.
func (e Employee) Display() {
	fmt.Printf("Employee: %s %s\n", e.FirstName, e.LastName)
	fmt.Printf("ID: %d\n", e.ID)
	fmt.Printf("Active: %t\n", e.IsActive)
}

// This method 'Deactivate' uses a pointer receiver (*Employee)
// so it can modify the original struct's data.
func (e *Employee) Deactivate() {
	e.IsActive = false
	fmt.Printf("%s %s has been deactivated.\n", e.FirstName, e.LastName)
}

// --- DEMONSTRATING ENCAPSULATION ---
// 
// Go's encapsulation is based on packages. Fields or methods
// starting with a capital letter are 'exported' (public).
// Fields starting with a lowercase letter are 'unexported' (private)
// and can only be accessed from within the same package.

// We can demonstrate encapsulation principles by using methods
// to control access to data, even if fields are public.
// The 'Deactivate' method is a good example of controlling
// *how* a field (IsActive) is changed.

func main() {
	// --- USING THE STRUCT ---
	// 
	// Create an instance of the Employee struct
	emp1 := Employee{
		FirstName: "Deepa",
		LastName:  "Sharma",
		ID:        101,
		IsActive:  true,
	}

	// Another way to initialize
	emp2 := Employee{FirstName: "Rohan", LastName: "Verma", ID: 102, IsActive: true}

	// --- CALLING METHODS ---
	fmt.Println("--- Employee 1 Info ---")
	emp1.Display()

	fmt.Println("\n--- Employee 2 Info ---")
	emp2.Display()

	// Call the method that modifies the struct
	fmt.Println("\n--- Modifying Employee 1 ---")
	emp1.Deactivate() // This changes emp1.IsActive to false
	
	fmt.Println("\n--- Employee 1 Info (After Update) ---")
	emp1.Display() // Now IsActive will be false
}