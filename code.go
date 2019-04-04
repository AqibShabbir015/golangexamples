package golangexamples

import "github.com/ehteshamz/greetings"
import "fmt"

func ConcatSlice(sliceToConcat []byte) string {
	appendString := ""

	for i := 0; i < len(sliceToConcat); i++ {
		appendString = appendString + string(sliceToConcat[i]) + "-"
	}
	return appendString
}

func Encrypt(sliceToConcat []byte, ceaserCount int) {
	encryptString := ""
	for i := 0; i < len(sliceToConcat); i++ {
		charAsci := int(sliceToConcat[i]) + ceaserCount
		encryptString = encryptString + string(charAsci)
	}
	fmt.Println("Encrypted String is ", encryptString)
}

func EZGreetings(name string) {
	print(greetings.PrintGreetings(name))
}
