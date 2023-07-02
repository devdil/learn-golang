package main

import (
	"errors"
	"fmt"
)

var SomeCustomEvent = errors.New("some Custom Event")

func main() {
	err := returnError()
	switch {
	case errors.Is(err, SomeCustomEvent):
		fmt.Println(" A custom event error")
	default:
		fmt.Println("Something default...")
	}
}

func returnError() error {
	return SomeCustomEvent
}
