package main

import (
	"fmt"
)

func main() {
	links, err := importLinks("docs/tabs.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("%+v\n", links)
}
