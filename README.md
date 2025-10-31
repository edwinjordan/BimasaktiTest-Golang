# BimasaktiTest-Golang

Kumpulan contoh-contoh kode Golang untuk pembelajaran dan referensi.

## Daftar Contoh

1. **Type Verification** - Pengecekan dan verifikasi tipe variabel
2. **Function Closure** - Implementasi closure dalam fungsi
3. **Class Inheritance** - Simulasi inheritance menggunakan struct embedding
4. **Permutation** - Program untuk menghasilkan permutasi
5. **Stop Goroutine** - Berbagai cara menghentikan goroutine

## Cara Menjalankan

### Menjalankan semua contoh:
```bash
go run .
```

### Menjalankan contoh tertentu:
```bash
go run . 1    # Type Verification
go run . 2    # Function Closure
go run . 3    # Class Inheritance
go run . 4    # Permutation
go run . 5    # Stop Goroutine
```

### Build dan jalankan:
```bash
go build
./BimasaktiTest-Golang      # Jalankan semua
./BimasaktiTest-Golang 1    # Jalankan contoh tertentu
```

## Penjelasan Setiap Contoh

### 1. Type Verification (`1_type_verification.go`)
Mendemonstrasikan berbagai cara untuk memeriksa tipe variabel di Go:
- Type assertion
- Type switch
- Menggunakan `reflect.TypeOf()`
- Menggunakan `fmt.Printf` dengan `%T`

### 2. Function Closure (`2_function_closure.go`)
Menunjukkan penggunaan closure dalam Go:
- Counter dengan state
- Closure dengan parameter
- Generator functions
- Fibonacci sequence dengan closure

### 3. Class Inheritance (`3_class_inheritance.go`)
Simulasi inheritance menggunakan struct embedding:
- Composition dan embedding
- Method overriding
- Polymorphism dengan interface
- Multi-level embedding

### 4. Permutation (`4_permutation.go`)
Implementasi algoritma permutasi:
- Permutasi array integer
- Permutasi string
- Permutasi dengan backtracking
- Permutasi dengan ukuran tertentu (nPr)

### 5. Stop Goroutine (`5_stop_goroutine.go`)
Berbagai metode untuk menghentikan goroutine:
- Menggunakan channel signal
- Menggunakan context
- Context dengan timeout
- Multiple goroutines dengan WaitGroup
- Quit channel pattern

## Requirements

- Go 1.16 atau lebih baru

## Lisensi

Kode ini disediakan untuk tujuan pembelajaran.