package practice

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func PrintPerson(p *Person) {
	fmt.Printf("name: %s; age: %d", p.Name, p.Age)
}

func PrintPerson2(p Person) {
	fmt.Printf("name: %s; age: %d", p.Name, p.Age)
}
