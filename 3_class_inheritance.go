package main

import "fmt"

// Example 3: Class Inheritance di Golang
// Go tidak memiliki inheritance klasik, tetapi menggunakan composition dan embedding

// Base "class" - Animal
type Animal struct {
	Name    string
	Age     int
	Species string
}

// Method untuk Animal
func (a *Animal) Eat() {
	fmt.Printf("%s is eating...\n", a.Name)
}

func (a *Animal) Sleep() {
	fmt.Printf("%s is sleeping...\n", a.Name)
}

func (a *Animal) GetInfo() string {
	return fmt.Sprintf("Name: %s, Age: %d, Species: %s", a.Name, a.Age, a.Species)
}

// "Derived class" - Dog (embedding Animal)
type Dog struct {
	Animal // Embedded struct - inheritance-like behavior
	Breed  string
}

// Method khusus untuk Dog
func (d *Dog) Bark() {
	fmt.Printf("%s says: Woof! Woof!\n", d.Name)
}

// Override method GetInfo
func (d *Dog) GetInfo() string {
	return fmt.Sprintf("%s, Breed: %s", d.Animal.GetInfo(), d.Breed)
}

// "Derived class" - Cat (embedding Animal)
type Cat struct {
	Animal
	Color string
}

// Method khusus untuk Cat
func (c *Cat) Meow() {
	fmt.Printf("%s says: Meow! Meow!\n", c.Name)
}

func (c *Cat) GetInfo() string {
	return fmt.Sprintf("%s, Color: %s", c.Animal.GetInfo(), c.Color)
}

// Interface untuk polymorphism
type Speaker interface {
	Speak()
}

// Implementasi interface untuk Dog
func (d *Dog) Speak() {
	d.Bark()
}

// Implementasi interface untuk Cat
func (c *Cat) Speak() {
	c.Meow()
}

// Contoh dengan multi-level embedding
type Vehicle struct {
	Brand string
	Model string
	Year  int
}

func (v *Vehicle) Start() {
	fmt.Printf("%s %s is starting...\n", v.Brand, v.Model)
}

type Car struct {
	Vehicle
	NumDoors int
}

func (c *Car) Honk() {
	fmt.Println("Beep! Beep!")
}

type ElectricCar struct {
	Car
	BatteryCapacity int
}

func (e *ElectricCar) Charge() {
	fmt.Printf("Charging battery (%d kWh capacity)...\n", e.BatteryCapacity)
}

func classInheritanceExample() {
	fmt.Println("=== Example 3: Class Inheritance (Composition & Embedding) ===")

	// Membuat instance Dog
	dog := Dog{
		Animal: Animal{
			Name:    "Buddy",
			Age:     3,
			Species: "Canine",
		},
		Breed: "Golden Retriever",
	}

	// Memanggil method dari "parent" dan "child"
	fmt.Println("\nDog Example:")
	fmt.Println(dog.GetInfo())
	dog.Eat()   // Method dari Animal
	dog.Sleep() // Method dari Animal
	dog.Bark()  // Method dari Dog

	// Membuat instance Cat
	cat := Cat{
		Animal: Animal{
			Name:    "Whiskers",
			Age:     2,
			Species: "Feline",
		},
		Color: "Orange",
	}

	fmt.Println("\nCat Example:")
	fmt.Println(cat.GetInfo())
	cat.Eat()   // Method dari Animal
	cat.Sleep() // Method dari Animal
	cat.Meow()  // Method dari Cat

	// Polymorphism dengan interface
	fmt.Println("\nPolymorphism Example:")
	animals := []Speaker{&dog, &cat}
	for _, animal := range animals {
		animal.Speak()
	}

	// Multi-level embedding
	fmt.Println("\nMulti-level Embedding Example:")
	tesla := ElectricCar{
		Car: Car{
			Vehicle: Vehicle{
				Brand: "Tesla",
				Model: "Model S",
				Year:  2024,
			},
			NumDoors: 4,
		},
		BatteryCapacity: 100,
	}

	tesla.Start()  // Method dari Vehicle
	tesla.Honk()   // Method dari Car
	tesla.Charge() // Method dari ElectricCar
	fmt.Printf("Tesla has %d doors\n", tesla.NumDoors)

	fmt.Println()
}
