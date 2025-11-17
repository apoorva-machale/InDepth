package main
import "fmt"

func getMonthlyPrice(tier string) int {
	pennies:= 0
	switch  tier{
		case "basic":
			pennies = int(100.00) * 100
		case "premium":
			pennies = int(150.00) * 100
		case "enterprise":
			pennies= int(500.00) * 100
		default:
			pennies= 0
	}
	fmt.Println(pennies)
	return pennies	
}

//Passing variables by value
func increment(x int) {
    x++
}

func getIncrement() {
    x := 5
    increment(x)

    fmt.Println(x)
    // still prints 5,
    // because the increment function
    // received a copy of x
}


//Ignoring return values
func getPoint()(x int, y int){
	return 3,4
}

func returnPoint() int {
	x, _ := getPoint()
	return x
}

func conversions(f func(int) int, a, b, c int) (int, int, int) {
	return f(a), f(b), f(c)
}

func main() {
	getMonthlyPrice("basic")
	getIncrement()
	fmt.Println(returnPoint())

	// using an anonymous function
	newX, newY, newZ := conversions(func(a int) int {
	    return a + a
	}, 1, 2, 3)
	fmt.Println(newX,newY,newZ)
}
