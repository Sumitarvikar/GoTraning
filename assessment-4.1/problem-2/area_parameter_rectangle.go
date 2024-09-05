package main

import (
	"fmt"
)

type Rectangle struct {
	Length, Breadth, area int
}

func (r *Rectangle) Area() {
	r.area = r.Length * r.Breadth
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.Length + r.Breadth)
}

func main() {
	var length, breadth int
	fmt.Println("Enter length and breadth of the rectangle:")
	fmt.Scan(&length, &breadth)

	rect := Rectangle{Length: length, Breadth: breadth}
	rect.Area()
	fmt.Printf("Area of Rectangle: %d\n", rect.area)
	fmt.Printf("Perimeter of Rectangle: %d\n", rect.Perimeter())
}
