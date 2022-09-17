package main

import (
	"fmt"
)

func philosopherEat(toLeft chan string, toRight chan string, fromLeft chan string, fromRight chan string) {
	for {
		fromLeft <- "I want to eat"
		<-toLeft

		fromRight <- "I want to eat"
		<-toRight

		fromLeft <- "I'm done eating"
		fromRight <- "I'm done eating"
	}
}

func checkFork(toLeft chan string, toRight chan string, fromLeft chan string, fromRight chan string) {
	select {
	case msgLeft := <-fromLeft:
		fmt.Println(msgLeft)
		toLeft <- "Take me now!"

		<-fromLeft
	case msgRight := <-fromRight:
		fmt.Println(msgRight)
		toRight <- "Take me now!"

		<-fromRight
	}

}

func main() {
	f0toLeft := make(chan string)
	f0toRight := make(chan string)
	f0fromLeft := make(chan string)
	f0fromRight := make(chan string)

	f1toLeft := make(chan string)
	f1toRight := make(chan string)
	f1fromLeft := make(chan string)
	f1fromRight := make(chan string)

	f2toLeft := make(chan string)
	f2toRight := make(chan string)
	f2fromLeft := make(chan string)
	f2fromRight := make(chan string)

	f3toLeft := make(chan string)
	f3toRight := make(chan string)
	f3fromLeft := make(chan string)
	f3fromRight := make(chan string)

	f4toLeft := make(chan string)
	f4toRight := make(chan string)
	f4fromLeft := make(chan string)
	f4fromRight := make(chan string)

	go philosopherEat(f0toLeft, f4toRight, f0fromLeft, f4fromRight)
	go philosopherEat(f1toLeft, f0toRight, f1fromLeft, f0fromRight)
	go philosopherEat(f1toLeft, f2toRight, f1fromLeft, f2fromRight)
	go philosopherEat(f2toLeft, f3toRight, f3fromLeft, f2fromRight)
	go philosopherEat(f4toLeft, f3toRight, f4fromLeft, f3fromRight)

	go checkFork(f0toLeft, f0toRight, f0fromLeft, f0fromRight)
	go checkFork(f1toLeft, f1toRight, f1fromLeft, f1fromRight)
	go checkFork(f2toLeft, f2toRight, f2fromLeft, f2fromRight)
	go checkFork(f3toLeft, f3toRight, f3fromLeft, f3fromRight)
	go checkFork(f4toLeft, f4toRight, f4fromLeft, f4fromRight)

	infiniteTime()
}

func infiniteTime() {
	for {
		var i int = 0
		i++
	}
}
