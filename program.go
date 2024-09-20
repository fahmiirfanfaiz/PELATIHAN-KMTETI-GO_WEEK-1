package main

import "fmt"

func main() {
	var celcius, fahrenheit, kelvin, reamur float32
	var choice int
	
	fmt.Print("Enter the temperature value in Celsius : ")
	fmt.Scanf("%g", &celcius)  

	fmt.Scanln()

	fmt.Println("Choose conversion option below:")
	fmt.Println("1. Convert Celsius to Fahrenheit")
	fmt.Println("2. Convert Celsius to Kelvin")
	fmt.Println("3. Convert Celsius to Reamur")
	fmt.Print("Enter your choice: ")
	fmt.Scanf("%d", &choice)  

	
	switch choice {
	case 1:
		fahrenheit = (celcius*9/5) + 32
		fmt.Printf("Fahrenheit: %.2f °F\n", fahrenheit)
	case 2:
		kelvin = celcius + 273.15
		fmt.Printf("Kelvin: %.2f K\n", kelvin)
	case 3:
		reamur = celcius*4/5
		fmt.Printf("Reamur: %.2f °R\n", reamur)
	default:
		fmt.Println("Invalid choice, please choose 1, 2, or 3.")
	}
}
