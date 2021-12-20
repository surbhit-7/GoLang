package main

import (
	"encoding/json"
	"fmt"
)

/*type Book struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
}
type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	author := Author{Name: "Shaggy", Age: 69}
	name := Book{Title: "Hello World", Author: author}
	byteArray, err := json.MarshalIndent(name, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteArray))
}*/
/*type Person struct {
	Name   string `json:"name"`
	Height int    `json:"height"`
}
*/
func main() {
	JsonString := `{"name":"Surbhit Vishwakarma" , "height" : 6}`

	var person map[string]interface{}

	err := json.Unmarshal([]byte(JsonString), &person)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(person)
}
