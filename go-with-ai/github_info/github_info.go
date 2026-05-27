package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type GitHubUser struct {
	Name        string `json:"name"`
	PublicRepos int    `json:"public_repos"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: go run github_info.go <username>\n")
		os.Exit(1)
	}

	user, err := fetchGitHubUser(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("Name: %s\n", user.Name)
	fmt.Printf("Public repos: %d\n", user.PublicRepos)
}

func fetchGitHubUser(username string) (GitHubUser, error) {
	client := http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get("https://api.github.com/users/" + username)
	if err != nil {
		return GitHubUser{}, fmt.Errorf("fetch github user: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return GitHubUser{}, fmt.Errorf("github api returned %s", resp.Status)
	}

	var user GitHubUser
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return GitHubUser{}, fmt.Errorf("decode github user: %w", err)
	}

	return user, nil
}
