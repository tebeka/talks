package main

import (
	"context"
	"errors"
	"fmt"
	"os"
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
	return false
}

func ctxRequestID(ctx context.Context) string {
	return "e262b9f6-2b43-474b-aae6-a1a6045dd120"
}

// START_LOGIN OMIT
func Login(ctx context.Context, login, passwd string) (*User, error) {
	if !checkPasswd(login, passwd) {
		err := AuthError{
			Login:     login,
			Reason:    "password",
			RequestID: ctxRequestID(ctx),
		}
		return nil, &err
	}

	// TODO
	return nil, nil
}

// END_LOGIN OMIT

func main() {
	// START_MAIN OMIT
	u, err := Login(context.Background(), "daffy", "duck season")
	if err != nil {
		var ae *AuthError
		if errors.As(err, &ae) { // HL
			fmt.Fprintf(os.Stderr, "error: bad password %#v\n", ae)
		} else {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
		}
		os.Exit(1)
	}
	fmt.Println(u)
	// END_MAIN OMIT
}
