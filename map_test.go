package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

/* 
1. map berfungsi untuk membuat data json menjadi dynamic
2. Secara otomatis, atribut akan menjadi key di map, dan value menjadi value di map
3. Namun karena value berupa interface{}, maka kita harus lakukan konversi secara manual jika ingin mengambil value nya
4. Dan tipe data map tidak mendukung JSON Tag lai
*/

func TestMapDecode(t *testing.T) {
	jsonString := `{"id":"p0001", "name":"Apple macbook", "price":"20000000"}` //bebas menentukan attribute nya
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])//["id"] -> key
	fmt.Println(result["name"])//["name"] -> key
	fmt.Println(result["price"])//["price"] -> key
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id" : "P0001",
		"name" : "Apple Mac Book Pro",
		"price" : 20000000,
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}