package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	log.Printf("Starting on port 8081")

	http.Handle("/memory", handler(decodeMemory))
	http.Handle("/stream", handler(decodeStream))

	log.Fatal(http.ListenAndServe(":8081", nil))
}

type handler func(http.ResponseWriter, *http.Request) error

func (fn handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func decodeMemory(w http.ResponseWriter, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)

	var action struct {
		Results []struct {
			DeviceName string `json:"device_name"`
		} `json:"results"`
	}

	if err := decoder.Decode(&action); err != nil {
		return err
	}

	w.WriteHeader(200)

	for _, r := range action.Results {
		fmt.Fprint(w, "%s <br />", r.DeviceName)
	}

	return nil
}

func decodeStream(w http.ResponseWriter, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)

	w.WriteHeader(200)

	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if token == "device_name" {
			name, errName := decoder.Token()
			if errName != nil {
				return err
			}

			fmt.Fprint(w, "%s <br />", name)
		}
	}

	return nil
}
