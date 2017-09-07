package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Attack struct {
	Name     string `json:"name"`
	AtkType  Type   `json:"type"`
	Category string `json:"category"`
	PP       int    `json:"pp"`
	Power    int    `json:"pow"`
	Accuracy int    `json:"acc"`
}

func NewAttack(rows *sql.Rows) *Attack {
	var name, t, category string
	var pp, power, accuracy int
	rows.Scan(&name, &t, &category, &pp, &power, &accuracy)
	atkType := TypeFromString(t)
	return &Attack{Name: name, AtkType: atkType, Category: category, PP: pp, Power: power, Accuracy: accuracy}
}

func ReadAttackDatabase(db *sql.DB) error {
	filename := "assets/attacks.csv"
	tableStr := "create table attacks (Name TEXT, Type TEXT, Category TEXT, PP INT, Power INT, Accuracy INT)"
	prepareStr := "insert into attacks (Name, Type, Category, PP, Power, Accuracy) values (?, ?, ?, ?, ?, ?)"
	return ReadDatabase(filename, tableStr, prepareStr)
}

func GetAttackFromAttacks(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	name := strings.ToLower(params["name"])
	rows, err := runQuery("select * from attacks where Name='" + name + "'")
	if err != nil {
		fmt.Println("Error running query")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer rows.Close()

	for rows.Next() {
		a := NewAttack(rows)
		json.NewEncoder(w).Encode(a)
		return
	}
	// If we are here, no rows were found
	w.WriteHeader(http.StatusNotFound)
}

func GetAttacksByType(w http.ResponseWriter, req *http.Request) {
	var attacks []Attack
	params := mux.Vars(req)
	attackType := strings.ToLower(params["type"])
	rows, err := runQuery("select * from attacks where Type='" + attackType + "'")
	if err != nil {
		fmt.Println("Error running query")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for rows.Next() {
		a := NewAttack(rows)
		attacks = append(attacks, *a)
	}
	if len(attacks) == 0 {
		// No attacks found, return 404 for bad type
		w.WriteHeader(http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(attacks)
	}
}

func GetAttacks(w http.ResponseWriter, req *http.Request) {
	var attacks []Attack
	rows, err := runQuery("select * from attacks")
	if err != nil {
		fmt.Println("Error running query")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer rows.Close()

	for rows.Next() {
		a := NewAttack(rows)
		attacks = append(attacks, *a)
	}
	json.NewEncoder(w).Encode(attacks)
}
