package golangjson

import (
	"encoding/json"
	"os"
	"testing"
)

//encoder itu berfungsi untuk mengirim data json ke dalam file

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")//Create berguna untuk menulis ke file nya lngsung
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName: "Tomioka",
		LastName: "Giyuu",
	}

	encoder.Encode(customer)
}