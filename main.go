package main

import (
	"bufio"
	"errors"
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
	{10, "X"},
	{9, "IX"},
	{8, "VIII"},
	{7, "VII"},
	{6, "VI"},
	{5, "V"},
	{4, "IV"},
	{3, "III"},
	{2, "II"},
	{1, "I"},
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	result, err := calculate(input)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("Output:", result)
}

func calculate(input string) (string, error) {
	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		return "", errors.New("строка не является математической операцией")
	}

	aStr, op, bStr := parts[0], parts[1], parts[2]

	a, isRomanA := romanToInt[aStr]
	b, isRomanB := romanToInt[bStr]

	if isRomanA && isRomanB {
		if a < 1 || a > 10 || b < 1 || b > 10 {
			return "", errors.New("римские числа должны быть в диапазоне от I до X")
		}
		return calculateRoman(a, b, op)
	}

	aInt, errA := strconv.Atoi(aStr)
	bInt, errB := strconv.Atoi(bStr)

	if errA == nil && errB == nil {
		if aInt < 1 || aInt > 10 || bInt < 1 || bInt > 10 {
			return "", errors.New("арабские числа должны быть в диапазоне от 1 до 10")
		}
		return calculateArabic(aInt, bInt, op)
	}

	return "", errors.New("используются одновременно разные системы счисления или неверный формат чисел")
}

func calculateRoman(a, b int, op string) (string, error) {
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
			return "", errors.New("деление на ноль")
		}
		result = a / b
	default:
		return "", errors.New("неизвестная операция")
	}

	if result < 1 {
		return "", errors.New("в римской системе нет отрицательных чисел или нуля")
	}

	return intToRomanExtended(result), nil
}

func calculateArabic(a, b int, op string) (string, error) {
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
			return "", errors.New("деление на ноль")
		}
		result = a / b
	default:
		return "", errors.New("неизвестная операция")
	}

	return strconv.Itoa(result), nil
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
