package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	//	"os"
	"net/http"
)

const (
	urlTemplte = "https://api.stocktwits.com/api/2/streams/symbol/%s.json"
)

// START_PARSE OMIT
func parseStocks(r io.Reader, symbol string) (map[string]int, error) {
	var reply struct { // HL
		Messages []struct {
			Symbols []struct {
				Name string `json:"symbol"`
			}
		}
	}

	if err := json.NewDecoder(r).Decode(&reply); err != nil {
		return nil, err
	}

	counts := make(map[string]int)
	for _, msg := range reply.Messages {
		for _, sym := range msg.Symbols {
			if sym.Name != symbol {
				counts[sym.Name]++
			}
		}
	}
	return counts, nil
}

// END_PARSE OMIT

func relatedStocks(symbol string) (map[string]int, error) {
	url := fmt.Sprintf(urlTemplte, symbol)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return parseStocks(resp.Body, symbol)
}

func main() {
	counts, err := relatedStocks("AAPL")
	if err != nil {
		log.Fatal(err)
	}

	for sym, count := range counts {
		fmt.Printf("%5s -> %d\n", sym, count)
	}
}
