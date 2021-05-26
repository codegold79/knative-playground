package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

const (
	secretPath = "/etc/secret/secret.toml"
	cmPath     = "/etc/config/configmap.toml"
)

func main() {
	log.Println("Starting server at port 8080.")

	http.HandleFunc("/", baseHandler)
	http.HandleFunc("/cm", cmHandler)
	http.HandleFunc("/secret", secretHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func requestCheck(r *http.Request) error {
	if r.Method != "GET" {
		return errors.New("only GET requests are supported")
	}
	return nil
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	err := requestCheck(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "<h1>Paths</h1><p>Look at what is saved in a secret at <a href='http://%s/secret'>%[1]s/secret</a> or in a configmap at <a href='http://%[1]s/cm'>%[1]s/secret</a>.</p>", r.Host)
}

func cmHandler(w http.ResponseWriter, r *http.Request) {
	err := requestCheck(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	msg, err := readMessage(cmPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusProcessing)
		return
	}

	fmt.Fprintf(w, "<h1>ConfigMap Message</h1><p>%s</p>", msg)
}

func secretHandler(w http.ResponseWriter, r *http.Request) {
	err := requestCheck(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	msg, err := readMessage(secretPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusProcessing)
		return
	}

	fmt.Fprintf(w, "<h1>Secret Message</h1><p>%s</p>", msg)
}
