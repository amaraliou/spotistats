package main

import (
	"fmt"
	"log"

	"github.com/amaraliou/spotistats/api"
)

func main() {
	api.Authenticate("", "")
	a, err := api.GetArtist("6oMuImdp5ZcFhWP0ESe6mG")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n%s\n", a.Name)
	fmt.Printf("\n%s\n", a.Href)
	fmt.Printf("\n%d\n", a.Popularity)
	fmt.Printf("\n%s\n", a.URI)
	fmt.Printf("\n%s\n", a.ID)
}
