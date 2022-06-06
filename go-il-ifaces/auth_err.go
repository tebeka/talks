package main

import (
	"fmt"
)

type User struct{}

// START_AE OMIT
type AuthError struct {
	Login     string
	Reason    string
	RequestID string
}

func (e *AuthError) Error() string {
	return e.Reason
}

// END_AE OMIT

func checkPasswd(login, passwd string) bool {
	return true
}

// START_LOGIN OMIT
func Login(login, passwd string) (*User, error) {
	var err *AuthError
	if !checkPasswd(login, passwd) {
		err := AuthError{
			Login:  login,
			Reason: "password",
		}
		return nil, &err
	}

	// TODO
	return nil, err
}

// END_LOGIN OMIT

func main() {
	// START_MAIN OMIT
	if u, err := Login("daffy", "duck season"); err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println(u)
	}
	// END_MAIN OMIT
}
