package main

import (
    "fmt"

    "daduic.com/spotify"
)

func main() {
    // Get a greeting message and print it.
    message := spotify.Hello("Gladys")
    fmt.Println(message)
}
