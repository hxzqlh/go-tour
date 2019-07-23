package main

import (
	"fmt"
	"sync"
)

var a, b, c, d []int

type MutexInfo struct {
 
	mutex sync.RWMutex
	 
	infos []int
	 
	}

func (m *MutexInfo)p1(p *[]int) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	*p = append(*p, 1)
}

func (m *MutexInfo)p2(p *[]int) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	*p = append(*p, 2)
}

func (m *MutexInfo)p3(p *[]int) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	*p = append(*p, 3)
}

func (m *MutexInfo)p4(p *[]int) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	*p = append(*p, 4)
}
func main() {
	ch1:=make(chan int)
	ch2:=make(chan int)
	ch3:=make(chan int)
	ch:=make(chan int)
	ch4:=make(chan int)
	m := MutexInfo{}
	for i:=0;i<1;i++{
	go func() {
		ch<-1
		<-ch4
		m.p1(&a)
		ch1<-1
		<-ch4
		m.p1(&d)
		ch2<-1
		<-ch4
		m.p1(&c)
		ch3<-1
		<-ch4
		m.p1(&b)
	}()
	go func() {
		ch<-1
		<-ch4
		m.p2(&b)
		ch1<-1
		<-ch4
		m.p2(&a)
		ch2<-1
		<-ch4
		m.p2(&d)
		ch3<-1
		<-ch4
		m.p2(&c)
	}()
	go func() {
		ch<-1
		<-ch4
		m.p3(&c)
		ch1<-1
		<-ch4
		m.p3(&b)
		ch2<-1
		<-ch4
		m.p3(&a)
		ch3<-1
		<-ch4
		m.p3(&d)
	}()
	go func() {
		ch<-1
		<-ch4
		m.p4(&d)
		ch1<-1
		<-ch4
		m.p4(&c)
		ch2<-1
		<-ch4
		m.p4(&b)
		ch3<-1
		<-ch4
		m.p4(&a)
	}()
	}
	for i:=0;i<4;i++{
		<-ch
		ch4<-1
	}
	for i:=0;i<4;i++{
		<-ch1
		ch4<-1
	}
	for i:=0;i<4;i++{
		<-ch2
		ch4<-1
	}
	for i:=0;i<4;i++{
		<-ch3
		ch4<-1
	}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
