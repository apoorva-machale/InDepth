package main
import (
	"fmt"
)
type expense interface{
	cost() int
}

type formatter interface{
	format() string
}

type email struct{
	isSubscribed bool
	body string
}

func (e email) cost() int{
	if e.isSubscribed == true{
		return 2*len(e.body) 
	}
	return 5*len(e.body)
}

func (e email) format() string{
	status:=""
	if e.isSubscribed{
		status = "Subscribed"
	}else{
		status = "Not Subscribed"
	}
	
	return fmt.Sprintf("'%v' | %v",e.body, status)
}

func main(){
	e := email {
		isSubscribed: false,
		body: "Hello, I am new here",
	}
	fmt.Println("Formatted:", e.format())
	fmt.Println("Cost:", e.cost())
}