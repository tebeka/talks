package main

import (
	"context"
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

func ctxRequestID(ctx context.Context) string {
	return "e262b9f6-2b43-474b-aae6-a1a6045dd120"
}

// START_LOGIN OMIT
func Login(ctx context.Context, login, passwd string) (*User, error) {
	var err error // HL
	if !checkPasswd(login, passwd) {
		err := AuthError{
			Login:     login,
			Reason:    "password",
			RequestID: ctxRequestID(ctx),
		}
		return nil, &err
	}

	// TODO
	return nil, err
}

// END_LOGIN OMIT

func main() {
	// START_MAIN OMIT
	if u, err := Login(context.Background(), "daffy", "duck season"); err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println(u)
	}
	// END_MAIN OMIT
}
