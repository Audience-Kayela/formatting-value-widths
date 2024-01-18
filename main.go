package main

import "fmt"

func main() {
	fmt.Println("---------------------------------------")
	fmt.Printf("%12s | %s\n", "Product", "Cost in Cents")
	fmt.Println("---------------------------------------")
	fmt.Printf("%12s | %2d\n", "Stamps", 50)
	fmt.Printf("%12s | %2d\n", "Paper Clips", 5)
	fmt.Printf("%12s | %2d\n", "Tape", 99)
	fmt.Println("---------------------------------------")

	// Formatting  fractional numbers

	fmt.Println("---------------------------------------")
	fmt.Println("With padding")
	fmt.Println("---------------------------------------")
	fmt.Printf("%7.3f\n", 12.3456)
	fmt.Printf("%7.2f\n", 12.3456)
	fmt.Printf("%7.1f\n", 12.3456)
	fmt.Println("---------------------------------------")
	fmt.Println("Without padding")
	fmt.Println("---------------------------------------")
	fmt.Printf("%.1f\n", 12.3456)
	fmt.Printf("%.2f\n", 12.3456)
}
