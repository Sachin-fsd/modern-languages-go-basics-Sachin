// Package 'custommath' provides simple math functions.
package custommath

// Add returns the sum of two integers.
// Note: The function name 'Add' is capitalized,
// making it 'exported' and usable by other packages.
func Add(a, b int) int {
	return a + b
}

// 'subtract' is not capitalized, so it is 'unexported'
// and can only be used within the 'custommath' package.
func subtract(a, b int) int {
	return a - b
}