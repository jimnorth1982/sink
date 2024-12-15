package main

import (
	"net/http"
	"sink/items/service"
)

func main() {
	http.HandleFunc("/items", itemsHandler)
	http.HandleFunc("/addItem", addItemHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	service.GetAllItems(w, r)
}

func addItemHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`<p>Login functionality is not implemented yet.</p>`))
}
