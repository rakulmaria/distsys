package main

import (
	"fmt"
	"time"
)

func fork(Lin chan string, Lout chan string, Rin chan string, Rout chan string) {
	for {
		select {
		case <-Lin:
			//fmt.Println("Left phil asked")
			Lout <- "im available"
			<-Lin
		case <-Rin:
			//fmt.Println("Right phil asked")
			Rout <- "im availabe"
			<-Rin
		}
	}

}

func phil(Lin chan string, Lout chan string, Rin chan string, Rout chan string) {
	for i := 0; i < 10; i++ {

		Lin <- "are u availabe"
		<-Lout
		//fmt.Println("Picked up left fork")
		Rin <- "are u availabe"
		<-Rout
		fmt.Println("eating")
		//delay()

		Lin <- "Done"
		Rin <- "Done"

	}
	fmt.Println("-----------Finished phil--------------")
}

func delay() {
	time.Sleep(time.Duration(3 * time.Second))
}

func main() {

	ch0 := make(chan string)
	//fork 0 channel in left
	ch1 := make(chan string)
	// fork 0 channel out left
	ch2 := make(chan string)
	//fork 0 channel in right
	ch3 := make(chan string)
	//fork 0 channel out right

	ch4 := make(chan string)
	//fork 1 channel in left
	ch5 := make(chan string)
	// fork 1 channel out left
	ch6 := make(chan string)
	//fork 1 channel in right
	ch7 := make(chan string)
	//fork 1 channel out right

	ch8 := make(chan string)
	//fork 2 channel in left
	ch9 := make(chan string)
	// fork 2 channel out left
	ch10 := make(chan string)
	//fork 2 channel in right
	ch11 := make(chan string)
	//fork 2 channel out right

	ch12 := make(chan string)
	//fork 3 channel in left
	ch13 := make(chan string)
	// fork 3 channel out left
	ch14 := make(chan string)
	//fork 3 channel in right
	ch15 := make(chan string)
	//fork 3 channel out right

	ch16 := make(chan string)
	//fork 4 channel in left
	ch17 := make(chan string)
	// fork 4 channel out left
	ch18 := make(chan string)
	//fork 4 channel in right
	ch19 := make(chan string)
	//fork 4 channel out right

	go phil(ch4, ch5, ch2, ch3)
	go phil(ch6, ch7, ch8, ch9)
	go phil(ch12, ch13, ch10, ch11)
	go phil(ch14, ch15, ch16, ch17)
	go phil(ch0, ch1, ch18, ch19)

	go fork(ch0, ch1, ch2, ch3)
	go fork(ch4, ch5, ch6, ch7)
	go fork(ch8, ch9, ch10, ch11)
	go fork(ch12, ch13, ch14, ch15)
	go fork(ch16, ch17, ch18, ch19)

	//time.Sleep(10 * time.Second)
	for {

	}

}
