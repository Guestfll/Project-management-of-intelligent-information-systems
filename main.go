package main

import (
	"bufio"
	"fmt"
	"os"
)

// ToUpper converts a lowercase Latin character into an uppercase character.
// If the character is not a lowercase letter of the Latin alphabet, the original character is returned.
func ToUpper(c byte) byte {
	// Check if the character is a lowercase letter of the Latin alphabet
	if c >= 'a' && c <= 'z' {
		// The difference between lowercase and uppercase letters in ASCII is 32
		return c - 32
	}
	// Return the original character if it is not a lowercase letter
	return c
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите строку символов: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка при чтении ввода:", err)
		return
	}

	if len(input) > 0 && input[len(input)-1] == '\n' {
		input = input[:len(input)-1]
	}

	// Convert the string into a byte slice for processing
	symbols := []byte(input)
	// Create a slice to store the converted characters
	upperSymbols := make([]byte, len(symbols))

	// Process each character
	for i, c := range symbols {
		upperSymbols[i] = ToUpper(c)
	}

	// Convert the byte slice back to a string
	upperString := string(upperSymbols)

	// Output the result
	fmt.Printf("Исходная строка: %s\n", input)
	fmt.Printf("Преобразованная строка: %s\n", upperString)
}
