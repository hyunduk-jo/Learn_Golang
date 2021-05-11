package main

import "fmt"

func main() {
	x := 10
	y := 5
	if z := x + y; z > 20 {
		fmt.Println("x + y is ", z)
	} else {
		fmt.Println("x and y should be bigger")
	}

	const (
		red = iota
		orange
		yellow
		greeting
		blue
		navy
		purple
	)
	color := purple
	switch color {
	case red:
		fmt.Print(1)
	case orange:
		fmt.Print(2)
	case yellow:
		fmt.Print(3)
	case greeting:
		fmt.Print(4)
	case blue:
		fmt.Print(5)
	case navy:
		fmt.Print(6)
	case purple:
		fmt.Print(7)
	}
}
