package api

import (
	"io"
	"log"
	"net/http"
)

func getRequest(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("Reponse failed with status code: %d and \nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	return body
}
