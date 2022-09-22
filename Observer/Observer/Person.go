package Observer

import "fmt"

type Person struct {
	name string
}

func (p *Person) handleEvent(vacancies []string) {
	fmt.Println("Hi dear " + p.name)
	fmt.Println("Vacancies updated: ")
	for _, vacancy := range vacancies {
		fmt.Println(vacancy)
	}
}

func NewPerson(name string) *Person {
	p := new(Person)
	p.name = name
	return p
}
