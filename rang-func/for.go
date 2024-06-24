package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// START_FOR OMIT
	cart := []string{"milk", "honey"}
	for i := range cart {
		fmt.Printf("%d ", i) // 0 1
	}

	for i, v := range cart {
		fmt.Printf("(%d %s) ", i, v) // (0 milk) (1 honey)
	}
	// END_FOR OMIT

	r := strings.NewReader(poem)
	// START_SCAN OMIT
	s := bufio.NewScanner(r)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	if err := s.Err(); err != nil { // HL
		fmt.Println("ERROR:", err)
	}
	// END_SCAN OMIT

	// START_WALK OMIT
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	// END_WALK OMIT

}

var poem = `The Road goes ever on and on
Down from the door where it began.
Now far ahead the Road has gone,
And I must follow, if I can,
Pursuing it with eager feet,
Until it joins some larger way
Where many paths and errands meet.
And whither then? I cannot say.
`
