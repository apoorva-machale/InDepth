package main

import (
	"fmt"
	"time"
)

//Interface is fullfiled when a type has all methods required by the interface
type message interface {
	getMessage() string
}

func sendMessage(msg message) (string, int) {
	m := msg.getMessage() // call interface method
	fmt.Println(m)        // print message
	return m, len(m) 
}

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

func main() {
	// Test birthdayMessage
	bm := birthdayMessage{
		birthdayTime:  time.Now(),
		recipientName: "Alice",
	}
	sendMessage(bm)

	// Test sendingReport
	sr := sendingReport{
		reportName:    "Weekly Newsletter",
		numberOfSends: 42,
	}
	sendMessage(sr)
}