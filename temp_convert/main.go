package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	options := []string{
		"1. Convert Celsius to Kelvin",
		"2. Convert Kelvin to Celsius",
		"3. Convert Celsius to Fahrenheit",
		"4. Convert Fahrenheit to Celsius",
		"q. Quit",
	}

	firstTimeIteration := true

	// Open a CSV file for writing
	file, err := os.Create("calculations.csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the header row to the CSV
	writer.Write([]string{"Input Temperature", "Conversion Type", "Result Temperature"})

	for {
		var an string

		if firstTimeIteration {
			an = "an"
			time.Sleep(1 * time.Second)
			firstTimeIteration = false
		} else {
			an = "another"
			time.Sleep(2 * time.Second)
		}

		fmt.Printf("Choose %s option:\n", an)
		for _, option := range options {
			fmt.Println(option)
		}

		var choice string
		fmt.Print("Enter your choice (1, 2, 3, 4, or q): ")
		fmt.Scanln(&choice)

		if strings.ToLower(choice) == "q" {
			fmt.Println("Exiting program. Goodbye!")
			break
		}

		var temp float64
		switch choice {
		case "1":
			fmt.Print("Enter the temperature in Celsius: ")
			fmt.Scanln(&temp)
			kelvin := temp + 273.15
			fmt.Printf("The temperature in Kelvin is: %.2f K\n", kelvin)
			writer.Write([]string{fmt.Sprintf("%.2f", temp), "Celsius to Kelvin", fmt.Sprintf("%.2f", kelvin)})
		case "2":
			fmt.Print("Enter the temperature in Kelvin: ")
			fmt.Scanln(&temp)
			celsius := temp - 273.15
			fmt.Printf("The temperature in Celsius is: %.2f °C\n", celsius)
			writer.Write([]string{fmt.Sprintf("%.2f", temp), "Kelvin to Celsius", fmt.Sprintf("%.2f", celsius)})
		case "3":
			fmt.Print("Enter the temperature in Celsius: ")
			fmt.Scanln(&temp)
			fahrenheit := (temp * 9 / 5) + 32
			fmt.Printf("The temperature in Fahrenheit is: %.2f °F\n", fahrenheit)
			writer.Write([]string{fmt.Sprintf("%.2f", temp), "Celsius to Fahrenheit", fmt.Sprintf("%.2f", fahrenheit)})
		case "4":
			fmt.Print("Enter the temperature in Fahrenheit: ")
			fmt.Scanln(&temp)
			celsius := (temp - 32) * 5 / 9
			fmt.Printf("The temperature in Celsius is: %.2f °C\n", celsius)
			writer.Write([]string{fmt.Sprintf("%.2f", temp), "Fahrenheit to Celsius", fmt.Sprintf("%.2f", celsius)})
		default:
			fmt.Println("Invalid choice. Please enter 1, 2, 3, 4, or q.")
		}
	}
}
