package main

import (
	"encoding/json"
)

func main() {
	var info map[string]string
	j := []byte(`{"wat": "wat"}`)
	
	json.Unmarshal(j, &info)
	println(info)
}