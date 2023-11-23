package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {

	customer := Customer {
		FirstName: "Laksa",
		MiddleName: "Bayu",
		LastName: "Trie Aji",
		Hobbies: []string{"Gaming", "Coding"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Laksa","MiddleName":"Bayu","LastName":"Trie Aji","Hobbies":["Gaming","Coding"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies[0])
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer {
		FirstName: "Gojo",
		Addresses: []Address{
			{
				Street : "Jujutsu Kaisen",
				Country : "Japan",
				PostalCode : "9999",
			},
			{
				Street : "Jujur Kasian",
				Country : "Japan",
				PostalCode : "8888",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Gojo","MiddleName":"","LastName":"","Hobbies":null,"Addresses":[{"Street":"Jujutsu Kaisen","Country":"Japan","PostalCode":"9999"},{"Street":"Jujur Kasian","Country":"Japan","PostalCode":"8888"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses[0].Street)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) { //json yg berisikan array
	jsonString := `[{"Street":"Jujutsu Kaisen","Country":"Japan","PostalCode":"9999"},{"Street":"Jujur Kasian","Country":"Japan","PostalCode":"8888"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}
	fmt.Println(addresses)
}

func TestOnlyJSONArrayComplex(t *testing.T) {
	
	addresses := []Address{
		{
			Street : "Jujutsu Kaisen",
			Country : "Japan",
			PostalCode : "9999",
		},
		{
			Street : "Jujur Kasian",
			Country : "Japan",
			PostalCode : "8888",
		},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}