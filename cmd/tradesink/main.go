package main

// rest api
import (
	"encoding/json"
	"net/http"
)

// weapon struct
type Weapon struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Power int    `json:"power"`
}

var weapons []Weapon = []Weapon{
	{ID: 1, Name: "Sword", Power: 100},
	{ID: 2, Name: "Axe", Power: 200},
	{ID: 3, Name: "Bow", Power: 300},
}

func getWeapons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weapons)
}

func addWeapon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var weapon Weapon
	json.NewDecoder(r.Body).Decode(&weapon)
	weapons = append(weapons, weapon)
	json.NewEncoder(w).Encode(weapons)
}

func main() {
	http.HandleFunc("/weapons", getWeapons)
	http.ListenAndServe(":8080", nil)
}
