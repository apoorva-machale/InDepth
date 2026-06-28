//defer allows function to be executed automatically just before its enclosing func returns.
package main

import(
	"fmt"
)

func bootup(){
	defer fmt.Println("TEXTIO BOOTUP DONE")
	ok := connectToDB()
	if !ok{
		return
	}
	ok = shouldConnectToPaymentProvider()
	if !ok{
		return
	}
	fmt.println("all systems ready!")
}

var shouldconnectToDB = true

func connectToDB() bool {
	fmt.Println("Connecting to database..")
	if shouldconnectToDB {
		fmt.Println("connected")
		return true
	}
	fmt.Println("connection failed")
	return false
}

var shouldConnectToPaymentProvider = true

func shouldConnectToPaymentProvider() bool{
	fmt.println("connection to payment provider..")
	if shouldConnectToPaymentProvider {
		fmt.println("connected")
		return true
	}
	fmt.println("connection failed")
	return false
}

func test(dbSuccess, paymentSuccess bool) {
	shouldConnectToDB = dbSuccess
	shouldConnectToPaymentProvider = paymentSuccess
	bootup()
	fmt.Println("====================================")
}

func main() {
	test(true, true)
	test(false, true)
	test(true, false)
	test(false, false)
}