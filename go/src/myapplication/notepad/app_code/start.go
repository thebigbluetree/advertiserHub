package app_code

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func Start() (string, string){
	fmt.Println("Welcome to Notepad app!")
	fmt.Println("Enter the message you'd like to store")
	in := bufio.NewReader(os.Stdin)
	UserMessage, _ := in.ReadString('\n')
	return UserMessage, time.Now().Format("2006.01.02 15:04:05")
}