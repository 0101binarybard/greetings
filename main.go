package main

import (
	"fmt"

	"github.com/0101binarybard/greetings/service"
	"github.com/0101binarybard/greetings/utils"
)

func main() {
	message := service.SayHello("Naresh")
	fmt.Println(message)

	message2 := utils.SayHello()
	fmt.Println(message2)
}
