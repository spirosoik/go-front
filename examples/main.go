package main

import (
	"fmt"

	"github.com/spirosoik/go-front/front"
)

func main() {
	cfg := front.Config{
		APIToken: "test_token",
		BaseURL:  "https://api2.frontapp.com",
	}

	c, err := front.New(&cfg)
	if err != nil {
		fmt.Print(err)
	}

	data, err := c.Inbox.List()
	fmt.Print(data.Results)
}
