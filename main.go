package main

import (
	"html/template"
	"net/http"
	"sink/items/service"
	"sink/items/types"
)

var itemsTemplate = template.Must(template.ParseFiles("templates/items.html"))

func main() {
	http.Handle("/static/css/", http.StripPrefix("/static/css/", http.FileServer(http.Dir("static/css"))))
	http.Handle("/static/img/", http.StripPrefix("/static/img/", http.FileServer(http.Dir("static/img"))))
	http.HandleFunc("/items", itemsHandler)
	http.HandleFunc("/addItem", addItemHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	itemsMap := service.GetAllItems()
	items := make([]types.Item, 0, len(itemsMap))
	for _, item := range itemsMap {
		items = append(items, item)
	}
	w.Header().Set("Content-Type", "text/html")
	if err := itemsTemplate.Execute(w, items); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func addItemHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`<p>Login functionality is not implemented yet.</p>`))
}
