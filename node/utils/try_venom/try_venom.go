package main

import (
	"fmt"
	"log"

	goever "github.com/move-ton/ever-client-go"
)

func main() {
	ever, err := goever.NewEver("", []string{"https://gql-testnet.venom.foundation/"}, "")
	if err != nil {
		log.Fatal(err)
	}

	defer ever.Client.Destroy()

	value, err := ever.Client.Version()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Version bindings is: ", value.Version)
}
