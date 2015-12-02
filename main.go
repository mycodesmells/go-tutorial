package main

import (
	"fmt"
	"time"

	"github.com/mycodesmells/go-tutorial/database"
)

func main() {
	fmt.Println("Hello world!")

	go func() {
		dbResult := database.MakeQuery()
		fmt.Println(dbResult)
	}()

	someRepetitiveTask(100)
}

func someRepetitiveTask(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("Hello, Bored Person #%d!\n", i)
		time.Sleep(200 * time.Millisecond)
	}
}