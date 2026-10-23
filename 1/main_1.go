package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

const salt = "go-2024"

func TypeOf(v any) string {
	return fmt.Sprintf("%T", v)
}

func ToString(v any) string {
	switch val := v.(type) {
	case int:
		return strconv.Itoa(val)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	case string:
		return val
	case bool:
		return strconv.FormatBool(val)
	case complex64:
		return strconv.FormatComplex(complex128(val), 'f', -1, 64)
	default:
		return fmt.Sprint(val)
	}
}

func JoinValues(values ...any) string {
	var sb strings.Builder
	for _, v := range values {
		sb.WriteString(ToString(v))
	}
	return sb.String()
}

func AddSalt(runes []rune, salt string) []rune {
	mid := len(runes) / 2
	res := make([]rune, 0, len(runes)+len(salt))
	res = append(res, runes[:mid]...)
	res = append(res, []rune(salt)...)
	res = append(res, runes[mid:]...)
	return res
}

func HashRunes(runes []rune, salt string) string {
	sum := sha256.Sum256([]byte(string(AddSalt(runes, salt))))
	return hex.EncodeToString(sum[:])
}

func main() {
	var numDecimal int = 42           // Десятичная система
	var numOctal int = 052            // Восьмеричная система
	var numHexadecimal int = 0x2A     // Шестнадцатиричная система
	var pi float64 = 3.14             // Тип float64
	var name string = "Golang"        // Тип string
	var isActive bool = true          // Тип bool
	var complexNum complex64 = 1 + 2i // Тип complex64

	values := []any{numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum}
	for _, v := range values {
		fmt.Printf("%v: %s\n", v, TypeOf(v))
	}

	joined := JoinValues(values...)
	fmt.Println("Строка:", joined)

	runes := []rune(joined)
	fmt.Println("Руны:", runes)

	fmt.Println("Строка с солью:", string(AddSalt(runes, salt)))
	fmt.Println("SHA256:", HashRunes(runes, salt))
}
