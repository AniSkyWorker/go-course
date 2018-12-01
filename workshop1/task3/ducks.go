package main

import "fmt"

type QuackBehavior func() string

type FlyBehavior interface {
	fly(direction string) string
}

type Duck struct {
	name  string
	fly   FlyBehavior
	quack QuackBehavior
}

type IDuck interface {
	Name() string
	Fly(direction string) string
	Quack() string
}

func (duck *Duck) Name() string {
	return duck.name
}

func (duck *Duck) Fly(direction string) string {
	return duck.fly.fly(direction)
}

func (duck *Duck) Quack() string {
	return duck.quack()
}

type NoFlyBehavior struct{}

func (beh NoFlyBehavior) fly(direction string) string {
	return "nowhere"
}

type CommonFlyBehavior struct{}

func (beh CommonFlyBehavior) fly(direction string) string {
	return "to " + direction
}

type FastFlyBehavior struct{}

func (beh FastFlyBehavior) fly(direction string) string {
	return "fast to " + direction
}

func RedHeadDuckQuackBehavior() string {
	return "Qaaaack!"
}

func MallockDuckQuackBehavior() string {
	return "Q--aaa-ck!"
}

func RubberDuckQuackBehavior() string {
	return "Peee!"
}

func playWithDuck(duck IDuck) {
	direction := "north"
	fmt.Println("Sent " + duck.Name() + " flying " + duck.Fly(direction) + " and duck says " + duck.Quack())
}

func main() {
	rubberDuck := Duck{"Rubber Duck", NoFlyBehavior{}, RubberDuckQuackBehavior}
	redHeadDuck := Duck{"ReadHead Duck", CommonFlyBehavior{}, RedHeadDuckQuackBehavior}
	mallockDuck := Duck{"Mallock Duck", FastFlyBehavior{}, MallockDuckQuackBehavior}
	playWithDuck(&rubberDuck)
	playWithDuck(&redHeadDuck)
	playWithDuck(&mallockDuck)
}
