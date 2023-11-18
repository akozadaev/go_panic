package main

import (
	"fmt"
	"log"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()

	var dividend int = 4
	var divider int = 0

	fmt.Println(dividend, divider)
	fmt.Println(dividend / divider)

	fmt.Print("Я все еще работаю, если паника обработана")

}
