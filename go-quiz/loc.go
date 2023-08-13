package main

import "fmt"

type Loc struct {
	Lat float64
	Lng float64
}

func main() {
	agents := map[Loc][]string{
		{51.4871871, -0.1270605}: {"007", "M"},
		{51.4850225, -0.1269318}: {"Q"},
	}
	fmt.Println(agents[Loc{51.4850225, -0.1269318}])
}
