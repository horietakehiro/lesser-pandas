package helper

// import (
// 	"fmt"
// )

// PadString pads the given string with given char.
func PadString(str, char string, length int) string {
	if length <= len(str) {
		return str
	}
	end := length - len(str)
	for i := 0; i < end; i++ {
		str += char
	}
	return str
}
