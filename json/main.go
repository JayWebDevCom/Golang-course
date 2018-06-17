package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

/*
Marshal - produces json
Unmarshal -  parses json encoded data to a variable

Encode writes the JSON encoding of x to the stream
*/

func main() {
	fmt.Println()
	fmt.Println("Marshal")

	rex := Dog{"dalmation", "rex", 40, "male", "pietro"}
	byteSlice, _ := json.Marshal(rex)
	fmt.Println(rex)
	fmt.Printf("%T\n", rex)
	fmt.Printf("%T\n", byteSlice) // byte being 8 bits
	fmt.Println("String conversion of byteSlice:", string(byteSlice))

	fmt.Println()
	fmt.Println("UnMarshal")

	jsonData := `{"Breed":"dalmation","Name":"rex","Mass in Kg":40,"sex":"male","Owner":"Jaiye"}`
	data := []byte(jsonData)
	// use verbose declaration because it must be at zero value
	var dog Dog
	// unmarshal by passing in the address
	json.Unmarshal(data, &dog)
	fmt.Println("Dog after Unmarshall:", dog)
	fmt.Printf("Dog type %T:\n", dog)
	fmt.Println("Dog breed:", dog.Breed)
	fmt.Println("Dog name:", dog.Name)
	fmt.Println("Dog weight:", dog.Weight)
	fmt.Println("Dog sex:", dog.sex)
	fmt.Println("Dog :", dog.Owner)

	fmt.Println()
	fmt.Println("Encode")

	// Encode to send out to a stream - use a writer

	tyrion := person{"Tyrion", "Lannister", 39, true}
	// NewEncoder takes a writer and returns a pointer to a new encoder
	encoder := json.NewEncoder(os.Stdout)
	fmt.Printf("Encoder is a %T\n", encoder)
	encoder.Encode(tyrion)

	fmt.Println()
	// Decode to read from a stream - use a reader
	fmt.Println("DeCode")
	// use verbose declaration because it must be at zero value
	var catelyn person
	catelynJSON := `{"FirstName":"Catelyn","LastName":"Tully","Age":50}`
	reader := strings.NewReader(catelynJSON)
	decoder := json.NewDecoder(reader)
	// get a pointer to a decoder
	fmt.Printf("Decoder is a %T\n", decoder)
	// decode by passing in the addressq
	decoder.Decode(&catelyn)
	fmt.Println("Catelyn:", catelyn)
}

// Dog is exported but does not need to be here
type Dog struct {
	// capitalized are exported and Marshaled
	// must be exported to be Marshaled out of the package
	// lowercase will not be Marshaled!
	// use tags to rename at Matshaling
	Breed  string
	Name   string
	Weight int    `json:"Mass in Kg"`
	sex    string // will not be marshaled or unmarshaled because field is not exported
	Owner  string `json:"-"` // will not be marshaled or unmarshaled bcause of tag
}

type person struct {
	FirstName   string
	LastName    string
	Age         int
	notExported bool
}
