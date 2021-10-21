package main

import (
	go_say_hello "github.com/ProgrammerZamanNow/go-say-hello"
	say_hello "github.com/arrosyad/moduleSayHello"
)
import "fmt"

func main() {
	fmt.Println(go_say_hello.SayHello())
	fmt.Println(say_hello.SayHello())
}
