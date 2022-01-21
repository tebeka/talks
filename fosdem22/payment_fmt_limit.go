package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// START_PTIME OMIT
type PTime struct {
	time.Time // HL
}

func (pt PTime) MarshalJSON() ([]byte, error) {
	// Step 1: Convert to value encoding/json can handle
	u := pt.Unix()
	// Step 2: Use json.Marshal
	return json.Marshal(u)
	// Step 3: There is no step 3
}

func (t *PTime) UnmarshalJSON(data []byte) error {
	var u int64
	if _, err := fmt.Fscanf(bytes.NewReader(data), "%d", &u); err != nil {
		return err
	}

	t.Time = time.Unix(u, 0)
	return nil
}

// END_PTIME OMIT

func main() {
	// START_MAIN OMIT
	t1 := PTime{time.Now()}
	t2 := t1.Add(time.Second)
	json.NewEncoder(os.Stdout).Encode(t2)
	// END_MAIN OMIT
}
