package main

//import "fmt"
import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("Starting Textio server...")
	// initialize variables here
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string
	smsSendingLimit=0
	costPerSMS=0
	hasPermission=false
	username=""
	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)

	//using walrus operator
	messageStart:="Happy birthday! You are now"
	age:=21
	messageEnd:="years old!"
	//fmt.Println - prints full message on single line separated by spaces
	fmt.Println(messageStart, age, messageEnd)

	//type cast
	tempFloat:= 54.45
	tempInt := int64(tempFloat)
	fmt.Printf(tempInt)

	//runes and string encoding
	const name = "🐻"
	fmt.Printf("constant 'name' byte length: %d\n", len(name))
	fmt.Printf("constant 'name' rune length: %d\n", utf8.RuneCountInString(name))
	fmt.Println("=====================================")
	fmt.Printf("Hi %s, so good to have you back in the arcanum\n", name)
}
