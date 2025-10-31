package main

import "fmt"

// Example 4: Program Permutasi dengan Golang
// Implementasi berbagai cara untuk menghasilkan permutasi

// Fungsi untuk menghasilkan permutasi dari slice integer
func permutationsInt(arr []int) [][]int {
	var result [][]int

	var permute func([]int, int)
	permute = func(arr []int, n int) {
		if n == 1 {
			temp := make([]int, len(arr))
			copy(temp, arr)
			result = append(result, temp)
			return
		}

		for i := 0; i < n; i++ {
			permute(arr, n-1)
			if n%2 == 1 {
				arr[0], arr[n-1] = arr[n-1], arr[0]
			} else {
				arr[i], arr[n-1] = arr[n-1], arr[i]
			}
		}
	}

	permute(arr, len(arr))
	return result
}

// Fungsi untuk menghasilkan permutasi dari string
func permutationsString(str string) []string {
	var result []string

	var permute func([]rune, int)
	permute = func(chars []rune, n int) {
		if n == 1 {
			result = append(result, string(chars))
			return
		}

		for i := 0; i < n; i++ {
			permute(chars, n-1)
			if n%2 == 1 {
				chars[0], chars[n-1] = chars[n-1], chars[0]
			} else {
				chars[i], chars[n-1] = chars[n-1], chars[i]
			}
		}
	}

	chars := []rune(str)
	permute(chars, len(chars))
	return result
}

// Fungsi permutasi dengan backtracking (alternatif)
func permutationsBacktrack(arr []int) [][]int {
	var result [][]int
	var current []int
	used := make([]bool, len(arr))

	var backtrack func()
	backtrack = func() {
		if len(current) == len(arr) {
			temp := make([]int, len(current))
			copy(temp, current)
			result = append(result, temp)
			return
		}

		for i := 0; i < len(arr); i++ {
			if used[i] {
				continue
			}

			current = append(current, arr[i])
			used[i] = true
			backtrack()
			current = current[:len(current)-1]
			used[i] = false
		}
	}

	backtrack()
	return result
}

// Fungsi untuk menghitung faktorial (jumlah permutasi)
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// Fungsi untuk permutasi dengan ukuran tertentu (nPr)
func permutationsOfSize(arr []int, r int) [][]int {
	var result [][]int
	var current []int
	used := make([]bool, len(arr))

	var backtrack func()
	backtrack = func() {
		if len(current) == r {
			temp := make([]int, len(current))
			copy(temp, current)
			result = append(result, temp)
			return
		}

		for i := 0; i < len(arr); i++ {
			if used[i] {
				continue
			}

			current = append(current, arr[i])
			used[i] = true
			backtrack()
			current = current[:len(current)-1]
			used[i] = false
		}
	}

	backtrack()
	return result
}

func permutationExample() {
	fmt.Println("=== Example 4: Permutation Program ===")

	// Contoh 1: Permutasi dari array integer
	fmt.Println("\n1. Permutasi dari [1, 2, 3]:")
	nums := []int{1, 2, 3}
	perms := permutationsInt(nums)
	fmt.Printf("Jumlah permutasi: %d (faktorial dari %d = %d)\n",
		len(perms), len(nums), factorial(len(nums)))
	for i, perm := range perms {
		fmt.Printf("Permutasi %d: %v\n", i+1, perm)
	}

	// Contoh 2: Permutasi dari string
	fmt.Println("\n2. Permutasi dari string 'ABC':")
	str := "ABC"
	strPerms := permutationsString(str)
	fmt.Printf("Jumlah permutasi: %d\n", len(strPerms))
	for i, perm := range strPerms {
		fmt.Printf("Permutasi %d: %s\n", i+1, perm)
	}

	// Contoh 3: Permutasi dengan backtracking
	fmt.Println("\n3. Permutasi menggunakan backtracking [1, 2, 3, 4]:")
	nums2 := []int{1, 2, 3, 4}
	perms2 := permutationsBacktrack(nums2)
	fmt.Printf("Jumlah permutasi: %d (faktorial dari %d = %d)\n",
		len(perms2), len(nums2), factorial(len(nums2)))
	// Tampilkan hanya 10 permutasi pertama untuk menghemat output
	fmt.Println("Menampilkan 10 permutasi pertama:")
	for i := 0; i < 10 && i < len(perms2); i++ {
		fmt.Printf("Permutasi %d: %v\n", i+1, perms2[i])
	}

	// Contoh 4: Permutasi dengan ukuran tertentu (nPr)
	fmt.Println("\n4. Permutasi 2 elemen dari [1, 2, 3, 4] (4P2):")
	nums3 := []int{1, 2, 3, 4}
	perms3 := permutationsOfSize(nums3, 2)
	fmt.Printf("Jumlah permutasi: %d\n", len(perms3))
	for i, perm := range perms3 {
		fmt.Printf("Permutasi %d: %v\n", i+1, perm)
	}

	fmt.Println()
}
