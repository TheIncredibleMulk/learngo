package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	people := [2]string{"nico", "flynn"}
	for _, person := range people {
		go isSexy(person, channel)
	}
	fmt.Println("Waiting for messages")
	resultOne := <-channel
	resultTwo := <-channel
	fmt.Println("Recieved this message: ", resultOne)
	fmt.Println("Recieved this message: ", resultTwo)
}

func isSexy(person string, channel chan string) {
	time.Sleep(time.Second * 10)
	channel <- person + " is sexy"
}

/*
func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
*/
