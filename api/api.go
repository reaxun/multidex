package api

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func runQuery(query string) (*sql.Rows, error) {
	db, err := sql.Open("sqlite3", "multidex.db")
	if err != nil {
		fmt.Println("Failed to open database")
		return nil, err
	}
	defer db.Close()
	return db.Query(query)
}

func ReadDatabase(filename, tableStr, prepareStr string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	csvr := csv.NewReader(f)
	rows, err := csvr.ReadAll()

	if _, err := db.Exec(tableStr); err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(prepareStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, row := range rows {
		s := make([]interface{}, len(row))
		for i, v := range row {
			if _, err = strconv.Atoi(v); err != nil {
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

func setup() {
	os.Remove("multidex.db")
	var err error
	db, err = sql.Open("sqlite3", "multidex.db")
	if err != nil {
		fmt.Println("Failed to open database")
		return
	}
	defer db.Close()

	pokemonStartTime := time.Now()
	if err := ReadPokemonDatabase(db); err != nil {
		fmt.Println(err)
		return
	}
	pokemonDuration := time.Since(pokemonStartTime)
	fmt.Println("Took " + pokemonDuration.String() + " to create Pokemon database")

	attackStartTime := time.Now()
	if err := ReadAttackDatabase(db); err != nil {
		fmt.Println("Failed to read file")
		return
	}
	attackDuration := time.Since(attackStartTime)
	fmt.Println("Took " + attackDuration.String() + " to create Attack database")
}

func addRoutes(router *mux.Router) {
	router.HandleFunc("/pokemon", GetPokedex).Methods("GET")
	router.HandleFunc("/pokemon/{name}", GetPokemonFromPokedex).Methods("GET")
	router.HandleFunc("/pokemon/type/{type}", GetPokemonByType).Methods("GET")
	router.HandleFunc("/attacks", GetAttacks).Methods("GET")
	router.HandleFunc("/attacks/{name}", GetAttackFromAttacks).Methods("GET")
	router.HandleFunc("/attacks/type/{type}", GetAttacksByType).Methods("GET")
	router.HandleFunc("/types", GetTypes).Methods("GET")
}

func StartAPI() {
	setup()

	router := mux.NewRouter()
	addRoutes(router)

	log.Fatal(http.ListenAndServe(":12345", router))
}
