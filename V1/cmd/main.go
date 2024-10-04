package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v65/github"
)

func main() {

	token := os.Getenv("GITHUB_AUTH_TOKEN")

	if token == "" {
		log.Fatal("token is needed")
	}

	args := os.Args

	if len(args) <= 1 {
		log.Fatal("username is needed")
	}

	user := args[1]
	ctx := context.Background()
	myClient := github.NewClient(nil).WithAuthToken(token)

	events, _, err := myClient.Activity.ListEventsPerformedByUser(ctx, user, true, nil)

	if err != nil {
		log.Fatalf("errors occur while accessing the user:%v", err)
	}

	for _, event := range events {
		printEvent(event)
	}

}

func printEvent(event *github.Event) {

	fmt.Printf("Event: %v\nRepo: %v\nActor: %v\nCreatedAt: %v\n\n",
		safeDereference(event.Type),
		safeDereference(event.Repo.Name),
		safeDereference(event.Actor.Login),
		event.CreatedAt.String())
}

func safeDereference(s *string) string {
	if s == nil {
		return "nil"
	} else {
		return *s
	}
}
