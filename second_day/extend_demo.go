package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	name  string
	age   int
	phone string
}
type Student struct {
	Human
	school string
	loan   float32
}
type Employee struct {
	Human
	company string
	money   float32
}

func (h *Human) SayHi() {
	fmt.Printf("Hi ,I am %s you can call me on %s\n", h.name, h.phone)
}
func (h *Human) Sing() {
	fmt.Println("la la la la")
}
func (e *Employee) SayHi() {
	fmt.Printf("Hi,I am %s ,I work at %s .Call me on %s\n", e.name, e.company, e.phone)
}
func (s *Student) BorrowMoney(amount float32) {
	s.loan += amount
}
func (e *Employee) SpendSalary(amount float32) {
	e.money -= amount
}

type Men interface {
	SayHi()
	Sing()
}
type YoungChap interface {
	SayHi()
	Sing()
	BorrowMoney(amount float32)
}
type ElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}

func (h Human) String() string {
	return "❰" + h.name + " - " + strconv.Itoa(h.age) + " years -  ✆ " + h.phone + "❱"
}

func main() {
	mark := Student{Human{"Mark", 6, "15622088256"}, "新华小学", 50}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 30, "18918915462"}, "星宫", 10000}
	tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}
	fmt.Println("This Human is :", tom)
	fmt.Println("======================================")
	var i Men
	mark.SayHi()
	sam.SayHi()
	i = &paul
	fmt.Println("This is Mike,a student")
	fmt.Println("======================================")
	i.SayHi()
	i = &tom
	i.SayHi()
	fmt.Println("======================================")
	x := make([]Men, 4)
	x[0], x[1], x[2], x[3] = &paul, &sam, &tom, &mark
	for _, value := range x {
		value.SayHi()
	}
}
