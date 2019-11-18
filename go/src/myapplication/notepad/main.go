package main

import (
	"fmt"
	"myapplication/notepad/app_code"
	"myapplication/notepad/dbconnection"
	//"os"
)

func main() {
	var UserChoice string
	message, dateandtime := app_code.Start()
	fmt.Println("Would you like to save your message ? (Yes or No)")
	fmt.Scanln(&UserChoice)
	if UserChoice == "Yes" {
		fmt.Println("Saving your message...")
		uc := dbconnection.NewUserContoller(dbconnection.GetSession())
		uc.InsertNote(message, dateandtime)
	} else {
		fmt.Println("Your message wasn't saved as per your choice. Have a good one!")
	}
}
