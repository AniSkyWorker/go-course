package main

import (
	"fmt"
)

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

func (b NoFlyBehavior) fly(direction string) string {
	return "nowhere"
}

type CommonFlyBehavior struct{}

func (b CommonFlyBehavior) fly(direction string) string {
	return "to " + direction
}

type FastFlyBehavior struct{}

func (b FastFlyBehavior) fly(direction string) string {
	return "fast to " + direction
}

func SimpleQuackBehavior() string {
	return "Qaaaack!"
}

func LongQuackBehavior() string {
	return "Q--aaa-ck!"
}

func PeepQuackBehavior() string {
	return "Peee!"
}

type RubberDuck struct {
	Duck
}

type MallardDuck struct {
	Duck
}

type RedHeadDuck struct {
	Duck
}

func getDuckName(duck IDuck) string {
	switch duck.(type) {
	case *RubberDuck:
		return "Rubber duck"
	case *MallardDuck:
		return "Mallard duck"
	case *RedHeadDuck:
		return "Read head duck"
	}
	return ""
}

func playWithDuck(duck IDuck) {
	direction := "north"
	fmt.Println("Sent " + getDuckName(duck) + " flying " + duck.Fly(direction) + " and duck says " + duck.Quack())
}

func NewRubberDuck() *RubberDuck {
	return &RubberDuck{Duck{"Rubber duck", NoFlyBehavior{}, PeepQuackBehavior}}
}

func NewMallardDuck() *MallardDuck {
	var duck MallardDuck
	duck.fly = FastFlyBehavior{}
	duck.quack = LongQuackBehavior
	return &duck
}

func NewRedHeadDuck() *RedHeadDuck {
	var duck RedHeadDuck
	duck.fly = CommonFlyBehavior{}
	duck.quack = SimpleQuackBehavior
	return &duck
}

func main() {
	playWithDuck(NewRedHeadDuck())
	playWithDuck(NewMallardDuck())
	playWithDuck(NewRubberDuck())
}
