package main

import (
	"fmt"

	"daduic.com/spotify"
)

func main() {
	message := spotify.Hello("Gladys")
	fmt.Println(message)
}
