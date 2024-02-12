package iteration

import "strings"

func Repeat(character string, repeatedCount int) string {
	// var repeated string
	// for i := 0; i < repeatedCount; i++ {
	// 	repeated += character
	// }
	// return repeated
	return strings.Repeat(character, repeatedCount)
}
