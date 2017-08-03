package api

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

type Attack struct {
	Name     string `json:"name"`
	AtkType  Type   `json:"typea"`
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
	f, err := os.Open("assets/attacks.csv")
	if err != nil {
		return err
	}
	defer f.Close()

	csvr := csv.NewReader(f)
	rows, err := csvr.ReadAll()

	table := `
    create table attacks (
        Name TEXT,
        Type TEXT,
        Category TEXT,
        PP INT,
        Power INT,
        Accuracy INT
    )
    `

	if _, err := db.Exec(table); err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("insert into attacks (Name, Type, Category, PP, Power, Accuracy) values (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, row := range rows {
		s := make([]interface{}, len(row))
		for i, v := range row {
			if i <= 2 {
				s[i] = strings.ToLower(v)
			} else {
				s[i] = v
			}
		}
		if _, err = stmt.Exec(s...); err != nil {
			fmt.Println("Exec failed")
			return err
		}
	}
	return tx.Commit()
}

func GetAttackFromAttacks(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	name := strings.ToLower(params["name"])
	rows, err := runQuery("select * from attacks where Name='" + name + "'")
	if err != nil {
		fmt.Println("Error running query")
		return
	}
	defer rows.Close()

	for rows.Next() {
		a := NewAttack(rows)
		json.NewEncoder(w).Encode(a)
		return
	}
	json.NewEncoder(w).Encode(&Attack{})
}

func GetAttacksByType(w http.ResponseWriter, req *http.Request) {
	var attacks []Attack
	params := mux.Vars(req)
	attackType := strings.ToLower(params["type"])
	rows, err := runQuery("select * from attacks where Type='" + attackType + "'")
	if err != nil {
		fmt.Println("Error running query")
		return
	}

	for rows.Next() {
		a := NewAttack(rows)
		attacks = append(attacks, *a)
	}
	json.NewEncoder(w).Encode(attacks)
}

func GetAttacks(w http.ResponseWriter, req *http.Request) {
	var attacks []Attack
	rows, err := runQuery("select * from attacks")
	if err != nil {
		fmt.Println("Error running query")
		return
	}
	defer rows.Close()

	for rows.Next() {
		a := NewAttack(rows)
		attacks = append(attacks, *a)
	}
	json.NewEncoder(w).Encode(attacks)
}
