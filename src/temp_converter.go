package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println("Temperature Converter")

	// Get user input for temperature type
	fmt.Print("Enter current temperature type (Celsius or Fahrenheit) or (c or f): ")
	tempType := getUserInput()

	//Get user input for value
	fmt.Print("Enter temperature value: ")
	tempValue := getUserInput()

	// Convert temperature based on user input
	convertedTemp, err := convertTemperature(tempType, tempValue)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Display the conversion result
	fmt.Printf("Converted temperature: %s\n", convertedTemp)
}


func getUserInput() string {
	//function to get input
	var input string
	fmt.Scanln(&input)
	return strings.TrimSpace(input)
}

func convertTemperature(tempType, tempValue string) (string, error) {
	//convert temp function
	switch strings.ToLower(tempType) {
	case "celsius":
		// Convert Celsius to Fahrenheit
		return celsiusToFahrenheit(tempValue), nil
	case "fahrenheit":
		// Convert Fahrenheit to Celsius
		return fahrenheitToCelsius(tempValue), nil
	case "c":
		// Convert Celsius to Fahrenheit
		return celsiusToFahrenheit(tempValue), nil
	case "f":
		// Convert Fahrenheit to Celsius
		return fahrenheitToCelsius(tempValue), nil
	default:
		return "", fmt.Errorf("invalid temperature type. Please enter Celsius or Fahrenheit")
	}
}

func celsiusToFahrenheit(celsius string) string {
	// Convert Celsius to Fahrenheit: (C * 9/5) + 32
	return fmt.Sprintf("%.2f°F", (toFloat(celsius) * 9/5) + 32)
}

func fahrenheitToCelsius(fahrenheit string) string {
	// Convert Fahrenheit to Celsius: (F - 32) * 5/9
	return fmt.Sprintf("%.2f°C", (toFloat(fahrenheit) - 32) * 5/9)
}

func toFloat(value string) float64 {
	// Convert string to float64
	var result float64
	fmt.Sscanf(value, "%f", &result)
	return result
}