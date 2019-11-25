package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type List struct {
	Value string
	Next  *List
}

func NewList(value string) *List {
	return &List{Value: value}
}

func (l *List) String() string {
	var values []string
	for ; l != nil; l = l.Next {
		values = append(values, l.Value)
	}
	return strings.Join(values, ", ")
}

func (l *List) Append(value string) {
	for ; l.Next != nil; l = l.Next {
	}
	l.Next = NewList(value)
}

func (l *List) MarshalJSON() ([]byte, error) {
	var values []string
	for ; l != nil; l = l.Next {
		values = append(values, l.Value)
	}
	return json.Marshal(values)
}

func (l *List) UnmarshalJSON(data []byte) error {
	var values []string
	if err := json.Unmarshal(data, &values); err != nil {
		return err
	}

	if len(values) == 0 {
		return fmt.Errorf("empty list")
	}

	l.Value = values[0]
	for _, v := range values[1:] {
		l.Append(v)
	}
	return nil
}

func main() {
	cart := NewList("onions")
	cart.Append("steak")
	cart.Append("potatos")

	out, err := json.Marshal(cart)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out)) // ["onions","steak","potatos"]

	var user struct {
		Name string
		Cart *List
	}
	data := []byte(`{
		"name": "Wile E. Coyote",
		"cart": ["canon", "brick", "gun powder"]
	}`)
	if err := json.Unmarshal(data, &user); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", user.Cart)
}
