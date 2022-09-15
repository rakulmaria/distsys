package main

import (
	"fmt"
)

func main() {
	fmt.Println("starting main")
	forks := make([]fork, 0)
	philosophers := make([]philosopher, 0)

	f0 := fork{0, true, make(chan bool), make(chan bool), make(chan bool), make(chan bool)}
	f1 := fork{1, true, make(chan bool), make(chan bool), make(chan bool), make(chan bool)}
	f2 := fork{2, true, make(chan bool), make(chan bool), make(chan bool), make(chan bool)}
	f3 := fork{3, true, make(chan bool), make(chan bool), make(chan bool), make(chan bool)}
	f4 := fork{4, true, make(chan bool), make(chan bool), make(chan bool), make(chan bool)}

	p0 := philosopher{0, f0, f4, 0}
	p1 := philosopher{1, f1, f0, 0}
	p2 := philosopher{2, f2, f1, 0}
	p3 := philosopher{3, f3, f2, 0}
	p4 := philosopher{4, f4, f0, 0}

	philosophers = append(philosophers, p0, p1, p2, p3, p4)
	forks = append(forks, f0, f1, f2, f3, f4)

	fmt.Println("philosophers: ", philosophers)
	fmt.Println("forks: ", forks)

	go triesToEat(philosophers[0])
	go triesToEat(philosophers[1])
	go triesToEat(philosophers[2])
	go triesToEat(philosophers[3])
	go triesToEat(philosophers[4])

	go lifeOfAFork(forks[0])
	go lifeOfAFork(forks[1])
	go lifeOfAFork(forks[2])
	go lifeOfAFork(forks[3])
	go lifeOfAFork(forks[4])

	clearTheTable(philosophers)
}

func think(p philosopher) {
	fmt.Print(p.name, " is thinking")
}

func lifeOfAFork(f fork) {
	fmt.Println("life of a fork running...")

	for {
		f.rightOut <- f.isAvailable
		fmt.Println(f.number, " is in between messages in lifeOfAFork")
		f.leftOut <- f.isAvailable

		select {
		case msgLeftIn := <-f.leftIn:
			fmt.Println("Message left in is recieved")
			if msgLeftIn {
				fmt.Println(f.number, " recieved ", msgLeftIn)
				f.isAvailable = false
			} else {
				f.isAvailable = true
				f.leftOut <- f.isAvailable
				f.rightOut <- f.isAvailable
			}
		case msgRightIn := <-f.rightIn:
			fmt.Println("Message right in is recieved")
			if msgRightIn {
				fmt.Println(f.number, " recieved ", msgRightIn)
				f.isAvailable = false
			} else {
				f.isAvailable = true
				f.leftOut <- f.isAvailable
				f.rightOut <- f.isAvailable
			}
		}

	}
	fmt.Println("forks are done")
	// for true {

	// 	msg2 := <-f.rightIn
	// 	msg1 := <-f.leftIn

	// 	if msg1 || msg2 {
	// 		f.isAvailable = false
	// 		if msg1 {
	// 			f.rightIn <- f.isAvailable
	// 			<-f.leftIn
	// 			f.isAvailable = true
	// 		}
	// 		if msg2 {
	// 			f.leftIn <- f.isAvailable
	// 			<-f.rightIn
	// 			f.isAvailable = true

	// 		}
	// 	}
	// 	lifeOfAFork(f)
	// }

}

func triesToEat(p philosopher) {
	fmt.Println(p.name, " is inside tries to eat now...")

	// Philosopher skal spørge fork om de er ledige - KIG PÅ DET HER
	p.left.rightIn <- false

	ans1 := <-p.left.rightOut
	if ans1 {
		p.left.rightIn <- true
		fmt.Println("ans1 is ", ans1)
		ans2 := <-p.right.leftOut
		if ans2 {
			fmt.Println("ans2 is ", ans2)
			p.right.leftIn <- true
			fmt.Println(p.name, " is going to eat now ")
			eat(p)
		} else {
			p.left.rightIn <- false
		}
	} else {
		<-p.right.leftOut
		think(p)
	}

}

func eat(p philosopher) {

	p.counter++

	fmt.Println(p.name, " is eating... nom nom... for the ", p.counter, "th time")

	p.left.rightIn <- false
	p.right.leftIn <- false

	think(p)

	fmt.Println(p.name, " has finished eaten")
}

func clearTheTable(p []philosopher) {
	doneEating := make([]bool, 5, 5)
	var allDone bool

	for !allDone {
		// fmt.Print(p[0].counter)

		for i := 0; i < len(p); i++ {
			if p[i].counter >= 3 {
				doneEating[i] = true
			}
		}

		if doneEating[0] && doneEating[1] && doneEating[2] && doneEating[3] && doneEating[4] {
			allDone = true
		}

	}
	fmt.Println("Table is cleared")

}

type philosopher struct {
	name  int
	left  fork
	right fork
	// hasLeftFork  chan bool
	// hasRightFork chan bool
	counter int
}

type fork struct {
	number      int
	isAvailable bool
	leftIn      chan bool
	rightIn     chan bool
	leftOut     chan bool
	rightOut    chan bool
}
