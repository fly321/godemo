package main

import "fmt"

type Usber interface {
	start(string) string
	stop(string) string
}

// 相机，手机结构体
type Phone struct {
	Name string
}

// 实现Usber接口
func (p Phone) start(name string) string {
	return fmt.Sprintf("%s is starting", name)
}

func (p Phone) stop(name string) string {
	return fmt.Sprintf("%s is stopping", name)
}

func main() {
	p := Phone{"phone"}
	var p1 Usber
	p1 = p
	start := p1.start(p.Name)
	fmt.Println(start)
	stop := p1.stop(p.Name)
	fmt.Println(stop)
}
