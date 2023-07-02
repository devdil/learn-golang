package main

import (
	"fmt"
)

func printMessage(message string, chn chan string) {
	chn <- message
}

func main() {
	times := 10
	ch := make(chan string, 10)
	for num := 0; num < times; num++ {
		go printMessage(fmt.Sprint(num), ch)
	}
	//time.Sleep(10 * time.Second)
	//close(ch)
	for elem := range ch {
		fmt.Println(elem)
	}

}
