package main

// gobyexample.com
// golang.org/ref/spec
// book - The Go Programming Language by A. Donovan
// book Go In Action by W Kennedy
// book The Wat Ti Go K Ascher

/*
reference types stored in the heap
value types stored in the stack

reference types - pass references to memory addresses
referencing types reference an underlying data type

map
reference type - stores an address to where data is stored
elements removed with the inbuild delete(mapName, key) function
if the key does not have a value, the key will return the zero value of that Type!

type
array
does not chandge in size i.e. not dynamic - has a fixed capacity
size is part of it's type
good performance

slice
is a reference type pointing to an underlying data structure
made of 3 things - address pointing to a data structure, length, capacity
can change in size i.e. dynamic - capacity is not fixed
make([]T, length[, capacity])
capacity is the length of the underlying array supporting a slice i.e. the array it
points to

if length has not ben set when using make(), indexed positions cannot
be referenced so need to use append

after capacity is exceeded
new underlying array is created and replaces the initial aray
*/

import (
	. "./array"
	. "./maps"
	. "./slice"
)

func main() {
	Array1()
	Array2()

	Slice1()
	Slice2()
	Slice3()
	Slice4()
	Slice5()
	Slice6()
	Slice7()
	Slice8()

	Map1()
	Map2()
	Map3()
}
