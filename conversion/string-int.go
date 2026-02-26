package main

import (
	"fmt"
	"strconv"
)

func intToStr() {

	num := 42
	str := fmt.Sprintf("%d", num)
	str2 := strconv.FormatInt(int64(num), 10)

	fmt.Printf("Sprintf - %s\n", str) // "42"

	fmt.Printf("FormatInt - %s\n", str2) // "42"

	num3 := 65
	char := string(rune(num3))
	fmt.Println(char)
	fmt.Printf("string() conversion - %s\n", char) // "A" (ASCII 65)

	num2 := 8364
	char2 := string(rune(num2))
	fmt.Printf("string() conversion € (Unicode 8364) - %s\n", char2) // "€" (Unicode 8364)

	num4 := 42
	str3 := strconv.Itoa(num4)
	fmt.Printf("strconv Itoa - %s\n", str3)     // "42"
	fmt.Printf("strconv Itoa Type: %T\n", str3) // Type: string
}

func strToIntAtoi() {
	str := "42"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Atoi - : %d", num)       // 42
		fmt.Printf("Atoi - Type: %T\n", num) // Type: int
	}
}
func strToInt() {
	str := "42"
	num, err := strconv.ParseInt(str, 10, 64) // base 10, 64-bit
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("ParseInt - strToInt: %d\n", num) // 42 (but as int64)
	}
}

func example() {
	// Int to String (number as text)
	fmt.Println("=== Int to String ===")
	num := 42

	// Method 1: strconv.Itoa
	str1 := strconv.Itoa(num)
	fmt.Printf("strconv.Itoa(%d) = %s\n", num, str1)

	// Method 2: fmt.Sprintf
	str2 := fmt.Sprintf("%d", num)
	fmt.Printf("fmt.Sprintf(%%d, %d) = %s\n", num, str2)

	// Method 3: Different bases
	str3 := strconv.FormatInt(int64(num), 16)
	fmt.Printf("Hex: %s\n", str3)

	// String to Int
	fmt.Println("\n=== String to Int ===")

	// Valid conversion
	validStr := "123"
	if result, err := strconv.Atoi(validStr); err == nil {
		fmt.Printf("strconv.Atoi(%s) = %d\n", validStr, result)
	}

	// Invalid conversion (error handling)
	invalidStr := "abc"
	if result, err := strconv.Atoi(invalidStr); err != nil {
		fmt.Printf("Error converting '%s': %v\n", invalidStr, err)
	} else {
		fmt.Printf("Result: %d\n", result)
	}

	// Character conversion
	fmt.Println("\n=== Int to Character ===")
	ascii := 65
	char := string(rune(ascii))
	fmt.Printf("string(%d) = %s\n", ascii, char)

	unicode := 8364
	euroChar := string(rune(unicode))
	fmt.Printf("string(%d) = %s\n", unicode, euroChar)

}

func main_string_int() {
	fmt.Printf("----------\n")
	intToStr()
	fmt.Printf("----------\n")
	strToIntAtoi()
	fmt.Printf("----------\n")
	strToInt()
	fmt.Printf("----------\n")
	fmt.Printf("example\n")
	fmt.Printf("----------\n")
	example()
}
