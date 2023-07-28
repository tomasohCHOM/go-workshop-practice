package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) SayHello() {
	fmt.Println("Hello, my name is", p.Name)
}

func main() {
	var mystring string = "Hello World"
	fmt.Println(mystring)
	x := []int{1, 2, 3}
	x = append(x, 5)
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}

	p := Person{Name: "Tomas", Age: 19}
	p.SayHello()
}
