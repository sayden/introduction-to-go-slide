package main

type Animal interface {
	Move()
}

type Dog struct {
	numberOfLegs int
	hasOwner     bool
}

func (d *Dog) bark() {
	println("Woof!")
}

func (d *Dog) Move() {
	println("Running!")
}

func main() {
	dog := Dog{
		4,
		true,
	}

	dog.bark()
	dog.Move()
}
