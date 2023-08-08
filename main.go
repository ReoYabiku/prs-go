package main

import (
	"log"
	"prs-go/entity"
)

func main() {
	var url entity.URL = "https://www.google.com"
	urls := []*entity.URL{&url}

	for _, url := range urls {
		err := url.Call()
		if err != nil {
			log.Fatal(err)
		}
	}
}