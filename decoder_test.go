package golangjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

//decoder itu dari json menuju stream json (menjuju tipe data di golang)
//decoder itu berfungsi untuk memanggil data json yang ada di dalam file

func TestStreamDecoder(t *testing.T) {

	reader, _ := os.Open("Customer.json")
	decoder := json.NewDecoder(reader)

	customer := &Customer{}
	decoder.Decode(customer) //mirip seperti unmarshal

	fmt.Println(customer)

}