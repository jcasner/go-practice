package reverseString

import (
	"fmt"
	"strings"
)

// Reverse takes a string as an input and reverses the order
// of the words. For example, "Hello World" becomes "World Hello"
func Reverse(input string) {
	originalArray := strings.Split(input, " ")

	var resultArray = make([]string, len(originalArray))

	for index, value := range originalArray {
		resultArray[len(originalArray)-index-1] = value
	}

	fmt.Println(strings.Join(resultArray, " "))
}
