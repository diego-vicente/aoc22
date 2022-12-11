package dsa

// Compute the Greatest Common Divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Compute the Least Common Multiple using GCD recursively
func LCM(a, b int, args ...int) int {
	result := a * b / GCD(a, b)

	// Compute it recursively for the rest of arguments
	for i := 0; i < len(args); i++ {
		result = LCM(result, args[i])
	}

	return result
}
