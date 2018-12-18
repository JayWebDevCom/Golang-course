package main

import "fmt"

// empty interface

type vehicles interface {
	// no methods defined
}

type vehicle struct {
	vehicles
	seats    int
	maxSpeed int
	Colout   string
}

type car struct {
	vehicles
	wheels int
	doors  int
}

type plane struct {
	vehicles
	jet bool
}

type boat struct {
	vehicle
	length int
}

// another example

type animal struct {
	noise string
}

type dog struct {
	animal
	friendly bool
	annoying bool
}

type cat struct {
	animal
	friendly bool
	annoying bool
}

// accepts anything, like object
func specs(a ...interface{}) {
	for _, v := range a {
		fmt.Println(v)
	}
}

func main() {
	prius := car{}
	tacoma := car{}
	bmw := car{}
	ford := car{}
	boeing747 := plane{}
	boeing757 := plane{}
	airbusa28 := plane{}
	sanger := boat{}

	vs := []vehicles{prius, tacoma, bmw, ford, boeing747, boeing757, airbusa28, sanger}

	for k, v := range vs {
		// fmt.Println(k, " - ", v, "\n")
		fmt.Printf("type at %d is %T\n", k, v)
	}

	fido := dog{animal{"woof"}, true, true}
	specs(fido)

	kitty := cat{animal{"meow"}, false, false}
	specs(kitty)

	specs(prius)

	specs([]interface{}{fido, kitty, prius})
}
