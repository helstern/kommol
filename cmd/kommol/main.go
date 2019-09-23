package main

import (
	"fmt"
	"github.com/helstern/kommol/internal/greetings"
)

func main() {
	myGreeting := greetings.Hello()
	fmt.Println(myGreeting)
}
