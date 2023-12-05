package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {
	bytes, err := json.Marshal(data)//function untuk melakukan konversi data ke json adalah Marshal
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestEncode(t *testing.T) {
	logJson("Rin")
	logJson(1)
	logJson(true)
	logJson([]string{"kakashi", "rin", "obito"})
}