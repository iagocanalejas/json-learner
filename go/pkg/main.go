package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var parsed any
	json.Unmarshal([]byte("{}"), &parsed)
	fmt.Printf("%v", parsed)
	fmt.Printf("%v", '\n')
}
