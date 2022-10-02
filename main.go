package main

import "fmt"

type person struct {
	first string
}

func (p person) speak() {
	fmt.Println(p.first)
}

//an interace says

type human interface {
	speak()
}

func main() {
	p := person{
		first: "ake",
	}

	fmt.Printf("%T\n", p)
}
