package main

import (
	"fmt"
)

func eat(p philosopher) {
	for p.counter < 3 {
		/* This is the if-statement we have to ensure no deadlock.
		By switching the order that two (or even one)
		philosophers grab the forks, we ensure that it
		won't deadlock since the philosophers will
		switch grabbing a right / left fork */
		if p.name == "Philosopher0" || p.name == "Philosopher2" {
			p.toLeftFork <- p.name + ": I'm hungry, let me grab my left fork"
			<-p.left.toLeftPhil

			p.toRightFork <- p.name + ": I'm hungry, let me grab my right fork"
			<-p.right.toRightPhil
		} else {
			p.toRightFork <- p.name + ": I'm hungry, let me grab my right fork"
			<-p.right.toRightPhil

			p.toLeftFork <- p.name + ": I'm hungry, let me grab my left fork"
			<-p.left.toLeftPhil
		}

		fmt.Println(p.name, " is eating.. nom nom...")
		// time.Sleep(time.Duration(3 * time.Second)) <--- uncomment this line for more realistic eating

		p.toLeftFork <- "I'm done"
		p.toRightFork <- "I'm done"

		p.counter++

		fmt.Println(p.name, " is thinking.. hmmm...")
		//time.Sleep(time.Duration(3 * time.Second)) <--- uncomment this line for more realistic thinking
	}
	fmt.Println(p.name, "is saying: \"I'm done eating and I'm really full.\"")
}

func checkFork(f fork) {
	for {
		select {
		case <-f.leftPhil[0].toLeftFork:
			// fmt.Println(msgLeft)
			f.toLeftPhil <- "Take me now!"

			<-f.leftPhil[0].toLeftFork
			// <-f.rightPhil[0].toRightFork
		case <-f.rightPhil[0].toRightFork:
			// fmt.Println(msgRight)
			f.toRightPhil <- "Take me now!"

			<-f.rightPhil[0].toRightFork
			// <-f.leftPhil[0].toLeftFork
		}
	}

}

func infiniteTime() {
	for {

	}
}

func main() {
	f0 := fork{"Fork0", make(chan string), make(chan string), make([]philosopher, 0), make([]philosopher, 0)}
	f1 := fork{"Fork1", make(chan string), make(chan string), make([]philosopher, 0), make([]philosopher, 0)}
	f2 := fork{"Fork2", make(chan string), make(chan string), make([]philosopher, 0), make([]philosopher, 0)}
	f3 := fork{"Fork3", make(chan string), make(chan string), make([]philosopher, 0), make([]philosopher, 0)}
	f4 := fork{"Fork4", make(chan string), make(chan string), make([]philosopher, 0), make([]philosopher, 0)}

	p0 := philosopher{"Philosopher0", f0, f4, 0, make(chan string), make(chan string)}
	p1 := philosopher{"Philosopher1", f1, f0, 0, make(chan string), make(chan string)}
	p2 := philosopher{"Philosopher2", f2, f1, 0, make(chan string), make(chan string)}
	p3 := philosopher{"Philosopher3", f3, f2, 0, make(chan string), make(chan string)}
	p4 := philosopher{"Philosopher4", f4, f3, 0, make(chan string), make(chan string)}

	f0.leftPhil = append(f0.leftPhil, p0)
	f0.rightPhil = append(f0.rightPhil, p1)

	f1.leftPhil = append(f1.leftPhil, p1)
	f1.rightPhil = append(f1.rightPhil, p2)

	f2.leftPhil = append(f2.leftPhil, p2)
	f2.rightPhil = append(f2.rightPhil, p3)

	f3.leftPhil = append(f3.leftPhil, p3)
	f3.rightPhil = append(f3.rightPhil, p4)

	f4.leftPhil = append(f4.leftPhil, p4)
	f4.rightPhil = append(f4.rightPhil, p0)

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
	name        string
	left        fork
	right       fork
	counter     int
	toLeftFork  chan string
	toRightFork chan string
}

type fork struct {
	name        string
	toLeftPhil  chan string
	toRightPhil chan string
	leftPhil    []philosopher
	rightPhil   []philosopher
}
