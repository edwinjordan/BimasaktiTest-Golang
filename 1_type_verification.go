package main

import (
	"fmt"
	"reflect"
)

// Example 1: Pengecekan atau Verifikasi Tipe Variable di Golang
// Contoh berbagai cara untuk memeriksa tipe variabel di Go

func typeVerificationExample() {
	fmt.Println("=== Example 1: Type Verification ===")

	// Cara 1: Menggunakan Type Assertion
	var i interface{} = "Hello"

	// Type assertion dengan pengecekan
	str, ok := i.(string)
	if ok {
		fmt.Printf("Value adalah string: %s\n", str)
	} else {
		fmt.Println("Value bukan string")
	}

	// Cara 2: Menggunakan Type Switch
	checkType := func(val interface{}) {
		switch v := val.(type) {
		case int:
			fmt.Printf("Tipe: int, Nilai: %d\n", v)
		case string:
			fmt.Printf("Tipe: string, Nilai: %s\n", v)
		case float64:
			fmt.Printf("Tipe: float64, Nilai: %f\n", v)
		case bool:
			fmt.Printf("Tipe: bool, Nilai: %t\n", v)
		default:
			fmt.Printf("Tipe tidak dikenali: %T\n", v)
		}
	}

	checkType(42)
	checkType("Golang")
	checkType(3.14)
	checkType(true)

	// Cara 3: Menggunakan reflect package
	var num int = 100
	var text string = "Test"
	var flag bool = false

	fmt.Printf("\nMenggunakan reflect.TypeOf():\n")
	fmt.Printf("Tipe dari num: %s\n", reflect.TypeOf(num))
	fmt.Printf("Tipe dari text: %s\n", reflect.TypeOf(text))
	fmt.Printf("Tipe dari flag: %s\n", reflect.TypeOf(flag))

	// Cara 4: Menggunakan fmt.Printf dengan %T
	fmt.Printf("\nMenggunakan fmt.Printf dengan %%T:\n")
	fmt.Printf("Tipe num: %T, nilai: %v\n", num, num)
	fmt.Printf("Tipe text: %T, nilai: %v\n", text, text)
	fmt.Printf("Tipe flag: %T, nilai: %v\n", flag, flag)

	// Contoh dengan struct
	type Person struct {
		Name string
		Age  int
	}

	person := Person{Name: "John", Age: 30}
	fmt.Printf("\nTipe person: %T\n", person)
	fmt.Printf("Tipe person (reflect): %s\n", reflect.TypeOf(person))

	fmt.Println()
}
