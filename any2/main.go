package main

import (
	"encoding/json"
	"fmt"
)

type model struct {
	ID      int
	Age     int
	Name    string
	Address string
}

func (m *model) Marshal(data model) string {
	str, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(str)
}

func main() {
	user := new(model)
	user.ID = 1
	user.Age = 20
	user.Name = "John"
	user.Address = "New York"

	fmt.Println(user.Marshal(*user)[01])

}
