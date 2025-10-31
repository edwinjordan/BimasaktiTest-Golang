package main

import "fmt"

// Example 2: Function Closure di Golang
// Closure adalah fungsi yang dapat mengakses variabel dari scope luar fungsi tersebut

func functionClosureExample() {
	fmt.Println("=== Example 2: Function Closure ===")

	// Contoh 1: Counter menggunakan closure
	counter := func() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}

	increment := counter()
	fmt.Println("Counter 1:", increment()) // Output: 1
	fmt.Println("Counter 2:", increment()) // Output: 2
	fmt.Println("Counter 3:", increment()) // Output: 3

	// Closure baru akan memiliki state sendiri
	newIncrement := counter()
	fmt.Println("New Counter 1:", newIncrement()) // Output: 1

	// Contoh 2: Closure dengan parameter
	multiplier := func(factor int) func(int) int {
		return func(num int) int {
			return num * factor
		}
	}

	double := multiplier(2)
	triple := multiplier(3)

	fmt.Printf("\nDouble 5: %d\n", double(5)) // Output: 10
	fmt.Printf("Triple 5: %d\n", triple(5))   // Output: 15

	// Contoh 3: Closure untuk konfigurasi
	greetingGenerator := func(greeting string) func(string) string {
		return func(name string) string {
			return fmt.Sprintf("%s, %s!", greeting, name)
		}
	}

	sayHello := greetingGenerator("Hello")
	sayHi := greetingGenerator("Hi")

	fmt.Println("\n" + sayHello("Alice")) // Output: Hello, Alice!
	fmt.Println(sayHi("Bob"))             // Output: Hi, Bob!

	// Contoh 4: Closure dengan multiple variables
	fibonacci := func() func() int {
		a, b := 0, 1
		return func() int {
			result := a
			a, b = b, a+b
			return result
		}
	}

	fib := fibonacci()
	fmt.Println("\nFibonacci sequence:")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", fib())
	}
	fmt.Println()
	fmt.Println()
}
