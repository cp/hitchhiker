package main

import (
	"github.com/go-martini/martini"
	"github.com/google/go-github/github"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"

	"log"
	"os"
)

func getClient() *github.Client {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env")
	}

	authToken := os.Getenv("GITHUB_TOKEN")

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: authToken},
	)

	tc := oauth2.NewClient(oauth2.NoContext, ts)

	return github.NewClient(tc)
}

func setAssignee(pull *github.PullRequest) error {
	log.Printf(github.Stringify(pull.User))
	return nil
}

func main() {
	m := martini.Classic()

	m.Get("/webhook", func() string {
		client := getClient()

		pull, _, err := client.PullRequests.Get("reflect", "reflect-js", 100)

		if err != nil {
			log.Fatal(err)
			return "Error getting pull request"
		}

		err = setAssignee(pull)

		if err != nil {
			log.Fatal(err)
			return "Error setting assignee"
		}

		return "well done!"
	})

	m.Get("/", func() string {
		return "Hello world!"
	})

	m.Run()
}
