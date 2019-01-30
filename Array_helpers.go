package Array

import "fmt"
import "testing"

// Helper for format strings
func format(str string, args ...Any) string {
	return fmt.Sprintf(str, args...)
}

// Helper for equivalance two arrays
func isEq(f Array, s Array) bool {
	if len(f) != len(s) {
		return false
	}

	for idx, item := range f {
		if item != s[idx] {
			return false
		}
	}

	return true
}

// Helper-Value for get stringified type
const StringType string = "Array.Array"

// Helper for get stringified type
func getType(arg Any) string {
	return format("%T", arg)
}

// Helper for log something
func log(args ...Any) {
	fmt.Println(args...)
}

// Helper for testing
func throwErr(t *testing.T, message string, expected Any, actual Any) {
	edge := "\n\n----\n\n"

	errorMessage := edge + message

	errorMessage += "\nExpected: " + format("%v", expected)
	errorMessage += "\nActual: " + format("%v", actual)

	errorMessage += edge

	t.Errorf(errorMessage)
}

// Helper for methods Every and Some
func isEven(item Any, _ int, _ Array) bool {
	newItem := toInt(item)

	return newItem % 2 == 1
}

// Helper for convert Any to int
func toInt(item Any) int {
	var newItem int

	val, ok := item.(int)
	if ok { newItem = val }

	return newItem
}
