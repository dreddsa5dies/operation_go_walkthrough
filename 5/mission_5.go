<<<<<<< HEAD
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//начало изменяемого кода
	guesses := make([]string, 0)

	// Agent Getter:
	// Brute force attack code
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			xString := strconv.Itoa(x)
			yString := strconv.Itoa(y)
			xPadding := strings.Repeat("0", 2-len(xString))
			yPadding := strings.Repeat("0", 2-len(yString))
			code := xPadding + xString + "-" + yPadding + yString
			if isCorrect(code) {
				guesses = append(guesses, code) //проверяем код до добавления его в массив
			}
			// someone is coming
		}
	}
	//конец изменяемого кода

	response := validateCode(guesses)
	fmt.Println(guesses)
}

func validateCode(codes []string) string {
	for _, c := range codes {
		if countAttempts() > 3 {
			return "3 Wrong Guesses - LOCKED!"
		}
		if isCorrect(c) {
			return "Code: " + c + ". Welcome Epoch."
		}
	}
	return "Incorrect code."
}
=======
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//начало изменяемого кода
	guesses := make([]string, 0)

	// Agent Getter:
	// Brute force attack code
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			xString := strconv.Itoa(x)
			yString := strconv.Itoa(y)
			xPadding := strings.Repeat("0", 2-len(xString))
			yPadding := strings.Repeat("0", 2-len(yString))
			code := xPadding + xString + "-" + yPadding + yString
			if isCorrect(code) {
				guesses = append(guesses, code) //проверяем код до добавления его в массив
			}
			// someone is coming
		}
	}
	//конец изменяемого кода

	response := validateCode(guesses)
	fmt.Println(guesses)
}

func validateCode(codes []string) string {
	for _, c := range codes {
		if countAttempts() > 3 {
			return "3 Wrong Guesses - LOCKED!"
		}
		if isCorrect(c) {
			return "Code: " + c + ". Welcome Epoch."
		}
	}
	return "Incorrect code."
}
>>>>>>> f125375fc8a280defe0b2f9398366c4c9e9246bc
