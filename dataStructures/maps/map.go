package maps

import "fmt"

// Map1 is exported
func Map1() {

	fmt.Println("------- Maps --------")

	// optional capacity - make(map[T]T[, capacity])
	myMap := make(map[string]int)
	fmt.Println("myMap:", myMap)
	fmt.Println("myMap length at initialization:", len(myMap))

	myMap["zero"] = 0
	myMap["one"] = 1
	fmt.Println("myMap:", myMap)
	fmt.Println("myMap length:", len(myMap))

	x, prs := myMap["zero"]

	/*
		prs is present
			the optional second return value when getting a return value from a map
			indicates if the key was present. used to disambiguate between missing keys
			and keys with zero values.
	*/

	fmt.Println("x:", x)
	fmt.Println("prs:", prs)

	x, prs = myMap["absent"]
	fmt.Println("x:", x)
	fmt.Println("prs:", prs)

	// composite literal definition
	myMap2 := map[string]int{
		"four": 4,
		"five": 5,
	}
	fmt.Println("myMap2 length:", len(myMap2))
	val, prs := myMap2["four"]
	fmt.Println("myMap2:", myMap2)
	fmt.Println("val:", val)
	fmt.Println("prs:", prs)

	// declaration
	myMap3 := map[string]int{}
	fmt.Println("myMap3 length:", len(myMap3))
	val, prs = myMap3["four"]
	fmt.Println("myMap3:", myMap2)
	fmt.Println("val:", val)
	fmt.Println("prs:", prs)
}

//Map2 is exported
func Map2() {
	// composite literal
	greetings := map[int]string{
		0: "good morning",
		1: "bonjour",
		2: "buenos dias",
		3: "bongiorno",
	}

	fmt.Println("greetings map:", greetings)

	delete(greetings, 2)

	if val, prs := greetings[2]; prs {
		fmt.Println("val:", val, "prs:", prs)
	} else {
		fmt.Println("that value does not exist")
		fmt.Println("val:", val, "prs:", prs)
	}
}

//Map3 is exported
func Map3() {
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
	}
	fmt.Println("element:", elements["H"]["name"])
	fmt.Println("element:", elements["Li"]["state"])

	for k, v := range elements {
		for k2, v2 := range v {
			fmt.Println("Key:", k, "inner key:", k2, "inner val:", v2)
		}
	}
}
