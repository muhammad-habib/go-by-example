package main

import "fmt"

type User struct {
	ID      int
	Age     int
	Balance float64
	Active  bool
	Data    [8192]byte
}

func ByValue(id int) User {
	return User{
		ID:      id,
		Age:     30,
		Balance: 1000.0,
		Active:  true,
	}
}

func ByPointer(id int) *User {
	return &User{
		ID:      id,
		Age:     30,
		Balance: 1000.0,
		Active:  true,
	}
}

func main() {
	fmt.Println("Go Escape Analysis - Stack vs Heap")
	fmt.Println("Run: go test -bench=. -benchmem")
	fmt.Println()

	v := ByValue(1)
	p := ByPointer(2)

	fmt.Printf("Value: ID=%d (stack)\n", v.ID)
	fmt.Printf("Pointer: ID=%d (heap)\n", p.ID)
}
