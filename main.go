package main

import "fmt"

type Speaker interface {
	speak()
}

type Human struct {
	name string
}

type Robot struct {
	serialNumb int
}

func (h Human) speak() {
	fmt.Println("Меня звать", h.name)

}

func (r Robot) speak() {
	fmt.Println("Мой номер", r.serialNumb, "я смотрю reelsы и курю айкос")
}

func main() {

	var object1 Speaker = Human{"Василий"}
	var object2 Speaker = Robot{153321}
	object1.speak()
	object2.speak()

}
