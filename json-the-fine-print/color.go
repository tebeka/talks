package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Color struct {
	Red   uint32
	Green uint32
	Blue  uint32
}

func (c *Color) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf("#%02x%02x%02x", c.Red, c.Green, c.Blue)
	return json.Marshal(s)
}

func (c *Color) UnmarshalJSON(data []byte) error {
	var r, g, b uint32
	data = data[2 : len(data)-1] // "#123456" -> 123456
	if _, err := fmt.Sscanf(string(data), "%02x%02x%02x", &r, &g, &b); err != nil {
		return err
	}

	c.Red, c.Green, c.Blue = r, g, b
	return nil
}

func main() {
	c := Color{0xFF, 0xD7, 0} // Gold
	data, err := json.Marshal(&c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

	var c2 Color
	if err := json.Unmarshal(data, &c2); err != nil {
		log.Fatal(err)
	}
	fmt.Println(c2)
}
