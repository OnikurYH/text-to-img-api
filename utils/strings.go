package utils

import (
	"bytes"
)

// StringsInsertRuneStep ...
func StringsInsertRuneStep(s string, step int, sep string) string {
	buffer := bytes.Buffer{}
	before := step - 1
	last := len(s) - 1
	for i, char := range s {
		buffer.WriteRune(char)
		if i%step == before && i != last {
			buffer.WriteString(sep)
		}
	}
	return buffer.String()
}
