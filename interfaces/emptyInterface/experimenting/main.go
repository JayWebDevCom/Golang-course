package main

import "fmt"

// empty interface

type vehicles interface {
	// no methods defined
	defa()
}

// func (v vehicles) d() {
// 	fmt.Println("this is method d")
// }
type defa struct {
}

func (d defa) defa() {
	fmt.Println("this is the method")
}

type vehicle struct {
	seats    int
	maxSpeed int
	Colout   string
}

type car struct {
	vehicle
	defau  defa
	wheels int
	doors  int
}

func (d car) defa() {
	d.defau.defa()
}

type plane struct {
	vehicle
	defau defa
	jet   bool
}

func (d plane) defa() {
	d.defau.defa()
}

type boat struct {
	vehicle
	defau  defa
	length int
}

func (d boat) defa() {
	d.defau.defa()
}

func main() {
	d := defa{}
	prius := car{vehicle{}, d, 5, 5}
	tacoma := car{vehicle{}, d, 6, 6}
	bmw := car{vehicle{}, d, 3, 3}
	ford := car{vehicle{}, d, 1, 6}
	boeing := plane{vehicle{}, d, true}

	vs := []vehicles{prius, tacoma, bmw, ford, boeing}
	for i, v := range vs {
		fmt.Printf("at %d\n", i)
		v.defa()
	}

}
