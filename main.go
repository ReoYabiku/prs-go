package main

import (
	"log"
	"os"
	"prs-go/infra"
	"prs-go/usecase"
)

func main() {
	endpoint := "https://api.github.com/graphql"
	accessToken := os.Getenv("GITHUB_TOKEN")

	graphQL := infra.NewGraphQL(endpoint, accessToken)
	u := usecase.NewUsecase(graphQL)

	err := u.OpenURLs()
	if err != nil {
		log.Fatal(err)
	}
}
