package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id string `json:"id"` //`json:"id"` -> json tag
	Name string `json:"name"`//`json:"name"` -> json tag
	ImageURL string `json:"image_url"`//`json:"image_url"` -> json tag
}

func TestJSONTag(t *testing.T) {
	product := Product {
		Id: "P0001",
		Name: "Macbook",
		ImageURL: "http://example.com/image.png",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":"P0001","name":"Macbook","image_url":"http://example.com/image.png"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}

	json.Unmarshal(jsonBytes, product)
	fmt.Println(product)
}