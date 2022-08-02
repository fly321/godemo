package main

import "fmt"

type Usber1 interface {
	start()
	stop()
}
type Computer1 struct {
}

func (c Computer1) work(usber Usber1) {
	usber.stop()
	usber.start()
}

type Phone1 struct {
	Name string
}

func (p Phone1) start() {
	fmt.Println(p.Name, " is starting")
}
func (p Phone1) stop() {
	fmt.Println(p.Name, " is stoping")
}

type Camera1 struct {
}

func main() {
	var computer1 Computer1
	var phone1 = Phone1{"小i米哈哈"}
	computer1.work(phone1)
}
