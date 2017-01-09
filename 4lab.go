package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Num struct {
	data      string
	recipient int 
}

var ch1 chan Num = make(chan Num)
var deep int = 0

func set(t Num) { 
	ch1 <- t

}

func get() { 
	curNum := <-ch1
	go set(curNum)

	if deep != curNum.recipient {
		deep++
		go get()
	} else {
		fmt.Print("Поток: ", deep)
		fmt.Print("Номер: ", <-ch1)
	}

}
func main() {
	rand.Seed(time.Now().UnixNano())
	N := rand.Intn(100)
	fmt.Print("поток:  ", N)

	var startNum Num
	startNum.recipient = N
	startNum.data = "smth data"

	go set(startNum)
	go get()
	time.Sleep(10000000)

}