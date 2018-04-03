package main

import (
	"os"
)

func msgchan(msg string) chan string {
	c := make(chan string)
	go func() {
		for {
			c <- msg
		}
	}()
	return c
}

func main() {
	rchan := [4]chan string{}

	for k := 0; k < 4; k++ {
		data := string(byte(k)+'1') + " "
		rchan[k] = msgchan(data)
	}

	fdA, errA := os.OpenFile("A.txt", os.O_CREATE|os.O_WRONLY, 0600)
	if errA != nil {
		panic(errA)
	}
	fdB, errB := os.OpenFile("B.txt", os.O_CREATE|os.O_WRONLY, 0600)
	if errB != nil {
		panic(errB)
	}
	fdC, errC := os.OpenFile("C.txt", os.O_CREATE|os.O_WRONLY, 0600)
	if errC != nil {
		panic(errC)
	}
	fdD, errD := os.OpenFile("D.txt", os.O_CREATE|os.O_WRONLY, 0600)
	if errD != nil {
		panic(errD)
	}

	defer func() {
		fdA.Close()
		fdB.Close()
		fdC.Close()
		fdD.Close()
	}()

	for j := 0; j < 1000; j++ {
		for i := 0; i < 4; i++ {
			if _, errA := fdA.WriteString(<-rchan[i%4]); errA != nil {
				panic(errA)
			}
			if _, errB := fdB.WriteString(<-rchan[(i+1)%4]); errB != nil {
				panic(errB)
			}
			if _, errC := fdC.WriteString(<-rchan[(i+2)%4]); errC != nil {
				panic(errC)
			}
			if _, errD := fdD.WriteString(<-rchan[(i+3)%4]); errD != nil {
				panic(errD)
			}
		}
	}
}
