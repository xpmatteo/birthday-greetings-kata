package main

import (
	"fmt"
	"log"
	"os"

	"github.com/birthday-greetings-kata/pkg/birthday_greetings"
)

func main() {
	service := birthday_greetings.NewBirthdayService()

	// Create a new XDate with the current date
	xDate := birthday_greetings.NewXDate()

	err := service.SendGreetings("employee_data.txt", xDate, "localhost", 25)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error sending greetings: %v\n", err)
		log.Fatal(err)
	}
}
