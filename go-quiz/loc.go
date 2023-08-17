package main

import "fmt"

type Location struct {
	Lat float64
	Lng float64
}

func main() {
	agents := map[Location][]string{
		{51.4871871, -0.1270605}: {"007", "M"},
		{51.4850225, -0.1269318}: {"Q"},
	}
	fmt.Println(agents[Location{51.4850225, -0.1269318}])
}
