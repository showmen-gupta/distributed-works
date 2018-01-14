// +build !solution

package lab1

// Task 1: Fibonacci numbers
//
// fibonacci(n) returns the n-th Fibonacci number, and is defined by the
// recurrence relation F_n = F_n-1 + F_n-2, with seed values F_0=0 and F_1=1.
func fibonacci(n uint) uint {
	//seed value checkings
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2) //recursively returning the values
	}

}
