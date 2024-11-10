package main

import (
	"fmt"
	"time"
)

var (
	fooFlag = false
	barFlag = false
)

func main() {
	fooPet := Pet{name: "Foo", walkingTime: 2}
	foo := FuncOwner{checkYardFunc: walk, pet: fooPet}
	barPet := Pet{name: "Bar", walkingTime: 3}
	bar := FuncOwner{checkYardFunc: walkVictim, pet: barPet}

	go startDay(foo)
	go startDay(bar)

	time.Sleep(30 * time.Second)
}

type Owner interface {
	checkYard()
	releasePet()
}

type Pet struct {
	name        string
	walkingTime int
}

func (p Pet) walk() {
	fmt.Printf("Started: %v\n", p.name)
	time.Sleep(time.Duration(p.walkingTime) * time.Second)
	fmt.Printf("Ended: %v\n", p.name)
}

type FuncOwner struct {
	pet           Pet
	checkYardFunc func(owner Owner)
}

func (owner FuncOwner) checkYard() {
	owner.checkYardFunc(owner)
}

func (owner FuncOwner) releasePet() {
	owner.pet.walk()
}

func walk(owner Owner) {
	for barFlag {
	}
	fooFlag = true

	owner.releasePet()

	fooFlag = false
}

func walkVictim(owner Owner) {
	for fooFlag {
	}
	barFlag = true
	if fooFlag {
		barFlag = false
		for fooFlag {
		}
	}
	barFlag = true

	owner.releasePet()

	barFlag = false
}

func startDay(owner Owner) {
	for i := 0; i < 5; i++ {
		owner.checkYard()
		time.Sleep(time.Duration(i) * time.Second)
	}
}
