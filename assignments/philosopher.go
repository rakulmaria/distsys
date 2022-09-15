package main

import (
	"fmt"
	// "time"
)

func eat(p philosopher) {
	p.left.fromLeftPhil <- p.name + ": I'm hungry, let me grab my left fork"
	<-p.left.toLeftPhil

	p.right.fromRightPhil <- p.name + ": I'm hungry, let me grab my right fork"
	<-p.right.toRightPhil
	fmt.Println(p.name, " is eating")

	p.left.fromLeftPhil <- "I'm done"
	p.right.fromRightPhil <- "I'm done"

	p.counter++
	fmt.Println(p.counter)

	fmt.Println(p.name, " is thinking")
	// time.Sleep(time.Duration(3 * time.Second))

}

func checkFork(f fork) {
	for {
		select {
		case msgLeft := <-f.fromLeftPhil:
			fmt.Println(msgLeft)
			f.toLeftPhil <- "Take me now!"

			<-f.fromLeftPhil
		case msgRight := <-f.fromRightPhil:
			fmt.Println(msgRight)
			f.toRightPhil <- "Take me now!"

			<-f.fromRightPhil
		}
	}

}

func infiniteTime() {
	for {
		var i int = 0
		i++
	}
}

func main() {
	f0 := fork{"Fork0", make(chan string), make(chan string), make(chan string), make(chan string)}
	f1 := fork{"Fork1", make(chan string), make(chan string), make(chan string), make(chan string)}
	f2 := fork{"Fork2", make(chan string), make(chan string), make(chan string), make(chan string)}
	f3 := fork{"Fork3", make(chan string), make(chan string), make(chan string), make(chan string)}
	f4 := fork{"Fork4", make(chan string), make(chan string), make(chan string), make(chan string)}

	p0 := philosopher{"Philosopher0", f0, f4, 0}
	p1 := philosopher{"Philosopher1", f1, f2, 0}
	p2 := philosopher{"Philosopher2", f2, f1, 0}
	p3 := philosopher{"Philosopher3", f3, f2, 0}
	p4 := philosopher{"Philosopher4", f3, f4, 0}

	go eat(p0)
	go eat(p1)
	go eat(p2)
	go eat(p3)
	go eat(p4)

	go checkFork(f0)
	go checkFork(f1)
	go checkFork(f2)
	go checkFork(f3)
	go checkFork(f4)

	infiniteTime()
}

type philosopher struct {
	name    string
	left    fork
	right   fork
	counter int
}

type fork struct {
	name          string
	toLeftPhil    chan string
	toRightPhil   chan string
	fromLeftPhil  chan string
	fromRightPhil chan string
}
