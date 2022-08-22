package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// START_COLOR OMIT
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

// END_COLOR OMIT

func main() {
	// START_MAIN OMIT
	c := Color{0xFF, 0xD7, 0} // Gold

	json.NewEncoder(os.Stdout).Encode(&c) // HL
	// END_MAIN OMIT
}
