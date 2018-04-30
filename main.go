package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ToDo represents a possible action - either done or not
type ToDo struct {
	Task string `json:"task"`
	Done bool   `json:"done"`
}

func serveTodos(w http.ResponseWriter, r *http.Request) {
	jsonData, err := ioutil.ReadFile("data/todos.json")
	if err != nil {
		fmt.Println("Error reading file")
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonData))
}

func serveAPIRules(w http.ResponseWriter, r *http.Request) {
	apiRules := `
		Here are the available API endpoints
		/todos - view the list of todos (served as JSON)
		/todos-from-data - view the list of todos generated from Go code
	`
	w.Write([]byte(apiRules))
}

func serveTodosFromData(w http.ResponseWriter, r *http.Request) {
	todoSlice := []ToDo{
		ToDo{Task: "walk the dog", Done: false},
		ToDo{Task: "wash the dishes", Done: true},
		ToDo{Task: "write in journal", Done: false},
	}

	todosJSON, err := json.Marshal(todoSlice)
	if err != nil {
		fmt.Errorf("Error creating JSON: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(todosJSON))
}

func main() {
	h := http.NewServeMux()

	h.HandleFunc("/todos", serveTodos)
	h.HandleFunc("/todos-from-data", serveTodosFromData)
	h.HandleFunc("/", serveAPIRules)
	fmt.Println("The server is now available at port :8100")
	if err := http.ListenAndServe(":8100", h); err != nil {
		panic(err)
	}
}
