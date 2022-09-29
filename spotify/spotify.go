package spotify

import (
	"io"
	"log"
	"net/http"
)

// TODO: Fetch the current playback from API
func Hello(name string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://pokeapi.co/api/v2/pokemon/ditto", nil)
	//req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return string(body)
}
