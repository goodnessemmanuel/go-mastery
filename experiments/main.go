package main

// Quack interface represents the ability to quack
/**
 * Quack This program demonstrates the use of interfaces in Go.
 * It defines a Quack interface that represents the ability to quack.
 * Two structs, RubberDuck and Squirrel, implement the Quack interface.
 * The makeItQuack function takes any type that satisfies the Quack interface and prints its quack.
 * In the main function, it creates instances of RubberDuck and Squirrel and passes them to makeItQuack.
 * In Go, "duck typing" refers to the language's implicit interface implementation. It is driven by the classic phrase:
 * "If it walks like a duck and quacks like a duck, then it's a duck." Unlike languages such as Java or C#, where
 * a class must explicitly state it implements an interface (e.g., class Sparrow implements Bird), a struct in Go
 * satisfies an interface automatically just by having the required methods.
 */
type Quack interface {
	Quack() string
}
type RubberDuck struct{}

func (r RubberDuck) Quack() string {
	return "Squeak"
}

type Squirrel struct{}

func (s Squirrel) Quack() string {
	return "Squawk"
}

// anything that satisfies the Quack interface can be passed to makeItQuack
func makeItQuack(q Quack) {
	println(q.Quack())
}

func main() {
	duck := RubberDuck{}
	squirrel := Squirrel{}
	// both the duck and the squirrel satisfies the Quack interface and they work seamlessly
	makeItQuack(duck)
	makeItQuack(squirrel)
}
