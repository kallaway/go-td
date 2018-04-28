package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type ToDo struct {
	Task string
	Done bool
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	jsonData, err := ioutil.ReadFile("data/todos.json")
	if err != nil {
		fmt.Println("Error reading file")
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonData))
}

func main() {
	http.HandleFunc("/", sayHello)

	if err := http.ListenAndServe(":8100", nil); err != nil {
		panic(err)
	}
}
