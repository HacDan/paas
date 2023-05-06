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

func rightPad(s string, p string, c string) []byte {
	pad, _ := strconv.Atoi(p)
	var padder string

	if len(c) > 1 {
		padder = c
	} else {
		padder = " "
	}

	if pad < 1 {
		message := &Message{Padded: s}
		response, _ := json.Marshal(message)
		return response
	}
	message := &Message{Padded: s + strings.Repeat(padder, pad)}
	response, _ := json.Marshal(message)
	return response
}

func leftPad(s string, p string, c string) []byte {
	pad, _ := strconv.Atoi(p)
	var padder string

	if len(c) > 1 {
		padder = c
	} else {
		padder = " "
	}

	if pad < 1 {
		message := &Message{Padded: s}
		response, _ := json.Marshal(message)
		return response
	}
	message := &Message{Padded: strings.Repeat(padder, pad) + s}
	response, _ := json.Marshal(message)
	return response
}

func laas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application,json")
	switch r.Method {
	case "GET":
		s := r.URL.Query().Get("string")
		pad := r.URL.Query().Get("pad")
		c := r.URL.Query().Get("padChar")
		response := leftPad(s, pad, c)

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
		c := r.URL.Query().Get("padChar")
		response := rightPad(s, pad, c)

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	default:
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func main() {
	http.HandleFunc("/raas", raas)
	http.HandleFunc("/laas", laas)

	log.Fatal(http.ListenAndServe(":1337", nil))
}
