package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type RespActor struct {
	Login string `json:"login"`
}

type RespRepo struct {
	Name string `json:"name"`
}

type Resp struct {
	Type       string    `json:"type"`
	Actor      RespActor `json:"actor"`
	Repo       RespRepo  `json:"repo"`
	Created_at time.Time `json:"created_at"`
}

func main() {
	token := os.Getenv("GITHUB_AUTH_TOKEN")

	if token == "" {
		log.Fatal("token is needed")
	}

	arg := os.Args

	if len(arg) <= 1 {
		log.Fatal("usrname is needed")
	}

	user := arg[1]

	client := &http.Client{}

	url := "https://api.github.com/users/" + user + "/events"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatalf("errors occur while creating http request:%v", err)
	}

	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalf("errors occur while sending http request:%v", err)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("error occur while reading the response:%v", err)
	}

	var items []Resp

	err = json.Unmarshal(data, &items)

	if err != nil {
		log.Fatalf("error occur while unmarshal the data:%v", err)
	}

	for _, event := range items {
		fmt.Printf("Event: %v\nRepo: %v\nActor: %v\nCreatedAt: %v\n\n",
			event.Type,
			event.Repo.Name,
			event.Actor.Login,
			event.Created_at)
	}

}
