package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Message struct {
	Padded string
}

func rightPad(s string, p string) []byte {
	pad, _ := strconv.Atoi(p)

	if pad < 1 {
		message := &Message{Padded: s}
		response, _ := json.Marshal(message)
		return response
	}
	padded := s + strings.Repeat("_", pad)
	message := &Message{Padded: padded}
	response, _ := json.Marshal(message)
	return response
}

func leftPad(s string, p string) []byte {
	pad, _ := strconv.Atoi(p)

	if pad < 1 {
		message := &Message{Padded: s}
		response, _ := json.Marshal(message)
		return response
	}
	padded := strings.Repeat("_", pad) + s
	message := &Message{Padded: padded}
	response, _ := json.Marshal(message)
	return response
}

func paas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application,json")
	switch r.Method {
	case "GET":
		s := r.URL.Query().Get("string")
		pad := r.URL.Query().Get("pad")
		response := leftPad(s, pad)

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	default:
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func raas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application,json")
	switch r.Method {
	case "GET":
		s := r.URL.Query().Get("string")
		pad := r.URL.Query().Get("pad")
		response := rightPad(s, pad)

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	default:
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func main() {
	http.HandleFunc("/raas", raas)
	http.HandleFunc("/paas", paas)

	log.Fatal(http.ListenAndServe(":1337", nil))
}
