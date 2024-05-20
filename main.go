package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanToInt = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var intToRoman = []struct {
	Value  int
	Symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Ошибка:", r)
		}
	}()

	result := calculate(input)
	fmt.Println("Output:", result)
}

func calculate(input string) string {
	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		panic("строка не является математической операцией")
	}

	aStr, op, bStr := parts[0], parts[1], parts[2]

	a, isRomanA := romanToInt[aStr]
	b, isRomanB := romanToInt[bStr]

	if isRomanA && isRomanB {
		if a < 1 || a > 10 || b < 1 || b > 10 {
			panic("римские числа должны быть в диапазоне от I до X")
		}
		return calculateRoman(a, b, op)
	} else if isRomanA || isRomanB {
		panic("римские числа должны быть в диапазоне от I до X")
	}

	aInt, errA := strconv.Atoi(aStr)
	bInt, errB := strconv.Atoi(bStr)

	if errA == nil && errB == nil {
		if aInt < 1 || aInt > 10 || bInt < 1 || bInt > 10 {
			panic("арабские числа должны быть в диапазоне от 1 до 10")
		}
		return calculateArabic(aInt, bInt, op)
	}

	panic("используются одновременно разные системы счисления или неверный формат чисел")
}

func calculateRoman(a, b int, op string) string {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			panic("деление на ноль")
		}
		result = a / b
	default:
		panic("неизвестная операция")
	}

	if result < 1 {
		panic("в римской системе нет отрицательных чисел или нуля")
	}

	return intToRomanExtended(result)
}

func calculateArabic(a, b int, op string) string {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			panic("деление на ноль")
		}
		result = a / b
	default:
		panic("неизвестная операция")
	}

	return strconv.Itoa(result)
}

func intToRomanExtended(num int) string {
	var result strings.Builder
	for _, entry := range intToRoman {
		for num >= entry.Value {
			result.WriteString(entry.Symbol)
			num -= entry.Value
		}
	}
	return result.String()
}
