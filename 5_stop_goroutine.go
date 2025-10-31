package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Example 5: Cara Menghentikan Goroutine di Golang
// Berbagai metode untuk menghentikan goroutine secara graceful

// Cara 1: Menggunakan channel sebagai signal
func stopWithChannel() {
	fmt.Println("\n1. Menghentikan Goroutine dengan Channel Signal:")

	stop := make(chan bool)
	done := make(chan bool)

	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-stop:
				fmt.Println("   Goroutine menerima signal stop")
				done <- true
				return
			case <-ticker.C:
				fmt.Println("   Goroutine sedang bekerja...")
			}
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("   Mengirim signal stop...")
	stop <- true
	<-done
	fmt.Println("   Goroutine telah berhenti")
}

// Cara 2: Menggunakan context.Context
func stopWithContext() {
	fmt.Println("\n2. Menghentikan Goroutine dengan Context:")

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan bool)

	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("   Goroutine menerima context cancellation")
				done <- true
				return
			case <-ticker.C:
				fmt.Println("   Goroutine sedang bekerja...")
			}
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("   Membatalkan context...")
	cancel()
	<-done
	fmt.Println("   Goroutine telah berhenti")
}

// Cara 3: Menggunakan context.WithTimeout
func stopWithTimeout() {
	fmt.Println("\n3. Menghentikan Goroutine dengan Timeout:")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	done := make(chan bool)

	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				fmt.Printf("   Goroutine berhenti karena: %v\n", ctx.Err())
				done <- true
				return
			case <-ticker.C:
				fmt.Println("   Goroutine sedang bekerja...")
			}
		}
	}()

	<-done
	fmt.Println("   Goroutine telah berhenti")
}

// Cara 4: Menggunakan sync.WaitGroup untuk multiple goroutines
func stopMultipleGoroutines() {
	fmt.Println("\n4. Menghentikan Multiple Goroutines dengan WaitGroup:")

	var wg sync.WaitGroup
	stop := make(chan bool)

	// Start multiple goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			ticker := time.NewTicker(500 * time.Millisecond)
			defer ticker.Stop()

			for {
				select {
				case <-stop:
					fmt.Printf("   Goroutine %d berhenti\n", id)
					return
				case <-ticker.C:
					fmt.Printf("   Goroutine %d sedang bekerja...\n", id)
				}
			}
		}(i)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("   Mengirim signal stop ke semua goroutines...")
	close(stop) // Menutup channel akan memberi signal ke semua goroutines
	wg.Wait()
	fmt.Println("   Semua goroutines telah berhenti")
}

// Cara 5: Menggunakan context dengan multiple goroutines
func stopMultipleWithContext() {
	fmt.Println("\n5. Menghentikan Multiple Goroutines dengan Context:")

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(ctx, &wg, i)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("   Membatalkan context...")
	cancel()
	wg.Wait()
	fmt.Println("   Semua goroutines telah berhenti")
}

func worker(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("   Worker %d menerima cancellation\n", id)
			return
		case <-ticker.C:
			fmt.Printf("   Worker %d sedang bekerja...\n", id)
		}
	}
}

// Cara 6: Menggunakan quit channel pattern
func stopWithQuitChannel() {
	fmt.Println("\n6. Menghentikan Goroutine dengan Quit Channel Pattern:")

	quit := make(chan struct{})
	finished := make(chan struct{})

	go func() {
		defer close(finished)
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-quit:
				fmt.Println("   Goroutine menerima quit signal")
				return
			case <-ticker.C:
				fmt.Println("   Goroutine sedang bekerja...")
			}
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("   Menutup quit channel...")
	close(quit)
	<-finished
	fmt.Println("   Goroutine telah berhenti")
}

func stopGoroutineExample() {
	fmt.Println("=== Example 5: Cara Menghentikan Goroutine ===")

	// Jalankan semua contoh
	stopWithChannel()
	time.Sleep(500 * time.Millisecond)

	stopWithContext()
	time.Sleep(500 * time.Millisecond)

	stopWithTimeout()
	time.Sleep(500 * time.Millisecond)

	stopMultipleGoroutines()
	time.Sleep(500 * time.Millisecond)

	stopMultipleWithContext()
	time.Sleep(500 * time.Millisecond)

	stopWithQuitChannel()

	fmt.Println()
}
