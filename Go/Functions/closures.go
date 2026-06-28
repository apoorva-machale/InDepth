package main
import (
	"fmt"
)
func adder() func(int) int {
	sum := 0
	return func(a int) int{
		sum+=a
		return sum
	}
}

func main() {
	add := adder()

	fmt.Println(add(5))  // 5
	fmt.Println(add(3))  // 8
	fmt.Println(add(10)) // 18
}