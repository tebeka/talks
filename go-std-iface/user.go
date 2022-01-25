package main

import "fmt"

// START_USER OMIT
type User struct {
	Login  string
	APIKey string
}

func (u User) Format(state fmt.State, verb rune) { // HL
	switch verb {
	case 's':
		fmt.Fprint(state, u.Login)
	case 'q':
		fmt.Fprintf(state, "%q", u.Login)
	case 'v':
		switch {
		case state.Flag('+'):
			fmt.Fprintf(state, "{Login:%s, APIKey:***}", u.Login)
		case state.Flag('#'):
			fmt.Fprintf(state, "%T{Login:%q, APIKey:\"***\"}", u, u.Login)
		default:
			fmt.Fprintf(state, "{%s, ***}", u.Login)
		}
	}
}

// END_USER OMIT

func main() {
	// START_MAIN OMIT
	u := User{"Bugs", "duckseason"}

	fmt.Printf("s : %s\n", u)
	fmt.Printf("v : %v\n", u)
	fmt.Printf("+v: %+v\n", u)
	fmt.Printf("#v %#v\n", u)
	fmt.Println(u)
	// END_MAIN OMIT
}
