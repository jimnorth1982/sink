package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sink/items/db"
	"sink/items/types"
)

func GetAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(db.GetItems())
}

func AddItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var item types.Item

	err := decoder.Decode(&item)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	if item.Name == "" {
		http.Error(w, "missing field 'name' from JSON object", http.StatusBadRequest)
		log.Println("missing field 'name' from JSON object")
		return
	}

	if decoder.More() {
		http.Error(w, "found extra data after JSON object", http.StatusBadRequest)
		log.Println("found extra data after JSON object")
		return
	}

	new_item, err := db.AddItem(item)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println(fmt.Printf("Successfully added item: %s", item.Name))
	json.NewEncoder(w).Encode(new_item)
}
