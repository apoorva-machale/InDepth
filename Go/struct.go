package main
import (
	"fmt"
)
type messageToSend struct {
	message   string
	sender    user
	recipient user
}

//Anonymous struct - used when it is required only once
// messageToSend := struct {
// 	message   string
// 	sender    user
// 	recipient user
// }


type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.name !=""  && mToSend.sender.number !=0 && mToSend.recipient.name!="" && mToSend.recipient.number!=0{
		return true
	}	
	
	return false
		
}

type authenticationInfo struct {
	username string
	password string
}

//struct method
func (a authenticationInfo) getBasicAuth() string {
	value := fmt.Sprintf("Authorization: Basic %v:%v", a.username,a.password)
	return value
}

func main(){
	sender := user{
		name:   "Alice",
		number: 12345,
	}

	recipient := user{
		name:   "Bob",
		number: 98765,
	}

	msg := messageToSend{
		message:   "Hello!",
		sender:    sender,
		recipient: recipient,
	}

	ok := canSendMessage(msg)
	fmt.Println("Can send message:", ok)

	user := authenticationInfo{
		username: "Apoorva",
		password: "XYZ",
	}

	fmt.Println(user.getBasicAuth())

}