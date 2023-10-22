package main

import (
	"fmt"
	"time"
)

func printMessage(msg string) {
	time.Sleep(1 * time.Second)
	fmt.Println(msg)
}

func main() {
	printMessage("Message 1")
	printMessage("Message 2")
	printMessage("Message 3")
}

