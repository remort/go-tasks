package main

import (
	"crypto/sha256"
	"encoding/hex"
	"reflect"
	"testing"
)

func TestCreateVariables(t *testing.T) {
	decimalInt, octalInt, hexInt, floatVal, stringVal, boolVal, complexVal := createVariables()

	if decimalInt != 42 {
		t.Errorf("decimalInt = %d; expect 42", decimalInt)
	}
	if octalInt != 052 {
		t.Errorf("octalInt = %d; expect 42 (052 as hex)", octalInt)
	}
	if hexInt != 0x2A {
		t.Errorf("hexInt = %d; expect 42 (0x2A as octal)", hexInt)
	}

	if floatVal != 3.14159 {
		t.Errorf("floatVal = %f; expect 3.14159", floatVal)
	}

	if stringVal != "Hello, Go!" {
		t.Errorf("stringVal = %s; expect 'Hello, Go!'", stringVal)
	}

	if boolVal != true {
		t.Errorf("boolVal = %t; expect true", boolVal)
	}

	if complexVal != complex64(3+4i) {
		t.Errorf("complexVal = %v; expect (3+4i)", complexVal)
	}
}

func TestPrintTypes(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("printTypes panic: %v", r)
		}
	}()

	vars := []any{42, 3.14, "test", true}
	printVarTypes(vars...)
}

func TestConvertToStringAndJoin(t *testing.T) {
	tests := []struct {
		name     string
		vars     []any
		expected string
	}{
		{
			name:     "Different type vars",
			vars:     []any{42, 3.14, "hello", true, complex64(1 + 2i)},
			expected: "423.14hellotrue(1+2i)",
		},
		{
			name:     "Single element",
			vars:     []any{100},
			expected: "100",
		},
		{
			name:     "Empty slice",
			vars:     []any{},
			expected: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := convertVarsToStringAndJoin(test.vars...)
			if result != test.expected {
				t.Errorf("convertToStringAndJoin(%v) = %s; expect %s", test.vars, result, test.expected)
			}
		})
	}
}

func TestStringToRuneSlice(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []rune
	}{
		{
			name:     "Ordinary ASCII string",
			input:    "Hello",
			expected: []rune{'H', 'e', 'l', 'l', 'o'},
		},
		{
			name:     "Cyrillic string",
			input:    "Привет",
			expected: []rune{'П', 'р', 'и', 'в', 'е', 'т'},
		},
		{
			name:     "Empty string",
			input:    "",
			expected: []rune{},
		},
		{
			name:     "String with emoji",
			input:    "😀🎉",
			expected: []rune{'😀', '🎉'},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := stringToRuneSlice(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("stringToRuneSlice(%s) = %v; expect %v", test.input, result, test.expected)
			}
		})
	}
}

func TestHashRuneSliceWithSalt(t *testing.T) {
	tests := []struct {
		name     string
		runes    []rune
		salt     string
		expected string
	}{
		{
			name:  "Simple string",
			runes: []rune("test"),
			salt:  "go-2024",
		},
		{
			name:  "Empty rune slice",
			runes: []rune{},
			salt:  "go-2024",
		},
		{
			name:  "Cyrillic string",
			runes: []rune("Привет"),
			salt:  "go-2024",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			saltRunes := []rune(tt.salt)
			midIndex := len(tt.runes) / 2

			expectedSaltedRunes := make([]rune, 0, len(tt.runes)+len(saltRunes))
			expectedSaltedRunes = append(expectedSaltedRunes, tt.runes[:midIndex]...)
			expectedSaltedRunes = append(expectedSaltedRunes, saltRunes...)
			expectedSaltedRunes = append(expectedSaltedRunes, tt.runes[midIndex:]...)

			expectedByteSlice := []byte(string(expectedSaltedRunes))
			expectedHash := sha256.Sum256(expectedByteSlice)
			expectedHashStr := hex.EncodeToString(expectedHash[:])

			resultHash := hashRuneSliceWithSalt(tt.runes, tt.salt)

			if resultHash != expectedHashStr {
				t.Errorf("hashRuneSliceWithSalt(%s, %s)\n"+
					"  Expected rune slice: %v\n"+
					"  Expected string: %q\n"+
					"  Expected hash: %s\n"+
					"  Rsulting hash: %s",
					string(tt.runes), tt.salt,
					expectedSaltedRunes,
					string(expectedSaltedRunes),
					expectedHashStr, resultHash)
			}
		})
	}
}

func TestHashConsistency(t *testing.T) {
	runes := []rune("Hello, World!")
	salt := "go-2024"

	hash1 := hashRuneSliceWithSalt(runes, salt)
	hash2 := hashRuneSliceWithSalt(runes, salt)

	if hash1 != hash2 {
		t.Errorf(
			"Hashes mismatch: %s != %s. hashRuneSliceWithSalt() is not determenistic.",
			hash1, hash2,
		)
	}
}

func TestSaltPlacement(t *testing.T) {
	runes := []rune("abcdef")
	salt := "go-2024"
	saltRunes := []rune(salt)

	midIndex := len(runes) / 2

	expectedSaltedRunes := make([]rune, 0, len(runes)+len(saltRunes))
	expectedSaltedRunes = append(expectedSaltedRunes, runes[:midIndex]...) // "abc"
	expectedSaltedRunes = append(expectedSaltedRunes, saltRunes...)        // "go-2024"
	expectedSaltedRunes = append(expectedSaltedRunes, runes[midIndex:]...) // "def"

	resultHash := hashRuneSliceWithSalt(runes, salt)

	expectedHash := sha256.Sum256([]byte(string(expectedSaltedRunes)))
	expectedHashStr := hex.EncodeToString(expectedHash[:])

	if resultHash != expectedHashStr {
		t.Errorf("Salting error.\n"+
			"Expected rune slice: %v (%s)\n"+
			"Resulting hash: %s\n"+
			"Expecting hash: %s",
			expectedSaltedRunes, string(expectedSaltedRunes), resultHash, expectedHashStr)
	}
}
