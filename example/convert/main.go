package main

import (
	"fmt"
	"github.com/rung/go-jsonstrconv"
	"log"
)

func main() {
	msg := "{\"id\":12345,\"name\":\"John\",\"tel\":12345}"
	converted, err := jsonstrconv.ToString([]byte(msg))
	if err != nil {
		log.Fatalln("it failed to convert. error=", err.Error())
	}
	fmt.Println(string(converted))
}
