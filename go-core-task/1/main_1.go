package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
	"strconv"
)

func createVariables() (int, int, int, float64, string, bool, complex64) {
	decimalInt := 42
	octalInt := 052
	hexInt := 0x2A

	floatVal := 3.14159
	stringVal := "Hello, Go!"
	boolVal := true
	complexVal := complex64(3 + 4i)

	return decimalInt, octalInt, hexInt, floatVal, stringVal, boolVal, complexVal
}

func printVarTypes(vars ...any) {
	fmt.Println("Task 1 and 2: Variables types and values")
	for i, v := range vars {
		fmt.Printf("Var %d: value = %v, type = %s\n", i+1, v, reflect.TypeOf(v))
	}
	fmt.Println()
}

func convertVarsToStringAndJoin(vars ...any) string {
	var result string

	for _, v := range vars {
		var str string

		switch val := v.(type) {
		case int:
			str = strconv.Itoa(val)
		case float64:
			str = strconv.FormatFloat(val, 'f', -1, 64)
		case string:
			str = val
		case bool:
			str = strconv.FormatBool(val)
		case complex64:
			str = fmt.Sprintf("%v", val)
		default:
			str = fmt.Sprintf("%v", val)
		}
		result += str
	}

	return result
}

func stringToRuneSlice(str string) []rune {
	return []rune(str)
}

func hashRuneSliceWithSalt(runes []rune, salt string) string {
	saltRunes := []rune(salt)

	midIndex := len(runes) / 2

	saltedRunes := make([]rune, 0, len(runes)+len(saltRunes))
	saltedRunes = append(saltedRunes, runes[:midIndex]...)
	saltedRunes = append(saltedRunes, saltRunes...)
	saltedRunes = append(saltedRunes, runes[midIndex:]...)

	byteSlice := []byte(string(saltedRunes))
	hash := sha256.Sum256(byteSlice)

	return hex.EncodeToString(hash[:])
}

func showRunes(runes []rune) {
	fmt.Println("Task 4: Rune slice")
	fmt.Printf("Rune slice: %v\n", runes)
	fmt.Printf("Slice length: %d\n", len(runes))
	fmt.Println("Runes and their corresponding symbols:")
	for i, r := range runes {
		fmt.Printf("  [%d] code: %d, symbol: %c\n", i, r, r)
	}
	fmt.Println()
}

func main() {
	decimalInt, octalInt, hexInt, floatVal, stringVal, boolVal, complexVal := createVariables()
	printVarTypes(decimalInt, octalInt, hexInt, floatVal, stringVal, boolVal, complexVal)

	combinedString := convertVarsToStringAndJoin(
		decimalInt, octalInt, hexInt, floatVal, stringVal, boolVal, complexVal,
	)
	fmt.Println("Task 3: Vars joined to string")
	fmt.Printf("Joined string: %s\n\n", combinedString)

	runeSlice := stringToRuneSlice(combinedString)
	showRunes(runeSlice)

	salt := "go-2024"
	hashResult := hashRuneSliceWithSalt(runeSlice, salt)
	fmt.Printf("Task 5: Salt rune slice with '%s' and hash it\n", salt)
	fmt.Printf("Resulting SHA256 hash: %s\n", hashResult)
}
