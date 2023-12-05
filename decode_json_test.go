package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJSON(t *testing.T) {
	jsonString := `{"FirstName":"Laksa","MiddleName":"Bayu","LastName":"Trie"}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)//function untuk melakukan konversi data dari json ke tipe data di golang
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
}