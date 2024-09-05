package main

import (
	"fmt"
)

// Quadrilateral interface
type Quadrilateral interface {
	Area() int
	Perimeter() int
}

// Rectangle type
type Rectangle struct {
	Length, Breadth int
}

func (r Rectangle) Area() int {
	return r.Length * r.Breadth
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.Length + r.Breadth)
}

// Square type
type Square struct {
	Side int
}

func (s Square) Area() int {
	return s.Side * s.Side
}

func (s Square) Perimeter() int {
	return 4 * s.Side
}

func Print(q Quadrilateral) {
	fmt.Printf("Area : %d\n", q.Area())
	fmt.Printf("Perimeter : %d\n", q.Perimeter())
}

func main() {
	var choice int
	fmt.Println("Enter 1 for Rectangle or 2 for Square:")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		rect := Rectangle{Length: 10, Breadth: 20}
		Print(rect)
	case 2:
		square := Square{Side: 15}
		Print(square)
	default:
		fmt.Println("Invalid choice")
	}
}
