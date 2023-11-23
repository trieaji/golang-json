package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName string
	MiddleName string
	LastName string
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName: "Laksa",
		MiddleName: "Bayu",
		LastName: "Trie",
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}