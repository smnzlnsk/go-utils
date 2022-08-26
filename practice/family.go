package practice

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func PrintPerson(p *Person) {
	fmt.Printf("name: %s; age: %d\n", p.Name, p.Age)
}
