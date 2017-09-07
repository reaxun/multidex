package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Stats struct {
	HP       int `json:"hp,omitempty"`
	Attack   int `json:"attack,omitempty"`
	Defense  int `json:"defense,omitempty"`
	SAttack  int `json:"sattack,omitempty"`
	SDefense int `json:"sdefense,omitempty"`
	Speed    int `json:"speed,omitempty"`
}

type Pokemon struct {
	Number int    `json:"number"`
	Name   string `json:"name"`
	Base   *Stats `json:"basestats,omitempty"`
	EVs    *Stats `json:"evs,omitempty"`
	IVs    *Stats `json:"ivs,omitempty"`
	TypeA  Type   `json:"typea"`
	TypeB  Type   `json:"typeb"`
}

func NewPokemon(rows *sql.Rows) *Pokemon {
	var number, hp, atk, def, satk, sdef, spd int
	var name, ta, tb string
	rows.Scan(&number, &name, &hp, &atk, &def, &satk, &sdef, &spd, &ta, &tb)
	typea := TypeFromString(ta)
	typeb := TypeFromString(tb)
	return &Pokemon{Number: number, Name: name, Base: &Stats{HP: hp, Attack: atk, Defense: def, SAttack: satk, SDefense: sdef, Speed: spd}, TypeA: typea, TypeB: typeb}
}

func ReadPokemonDatabase(db *sql.DB) error {
	filename := "assets/pokemon.csv"
	tableStr := "create table pokemon (Number INT, Name TEXT, HP INT, Attack INT, Defense INT, SAttack INT, SDefense INT, Speed INT, TypeA TEXT, TypeB TEXT)"
	prepareStr := "insert into pokemon (Number, Name, HP, Attack, Defense, SAttack, SDefense, Speed, TypeA, TypeB) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	return ReadDatabase(filename, tableStr, prepareStr)
}

func GetPokemonFromPokedex(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	name := strings.ToLower(params["name"])
	rows, err := runQuery("select * from pokemon where Name='" + name + "'")
	if err != nil {
		fmt.Println("Error running query")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer rows.Close()

	for rows.Next() {
		p := NewPokemon(rows)
		json.NewEncoder(w).Encode(p)
		return
	}
	// If we are here, no rows were found
	w.WriteHeader(http.StatusNotFound)
}

func GetPokemonByType(w http.ResponseWriter, req *http.Request) {
	var pokedex []Pokemon
	params := mux.Vars(req)
	pokemonType := strings.ToLower(params["type"])
	rows, err := runQuery("select * from pokemon where TypeA='" + pokemonType + "' OR TypeB='" + pokemonType + "'")
	if err != nil {
		fmt.Println("Error running query")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer rows.Close()

	for rows.Next() {
		p := NewPokemon(rows)
		pokedex = append(pokedex, *p)
	}
	if len(pokedex) == 0 {
		// If we are here, the type was not found
		w.WriteHeader(http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(pokedex)
	}
}

func GetPokedex(w http.ResponseWriter, req *http.Request) {
	var pokedex []Pokemon
	rows, err := runQuery("select * from pokemon")
	if err != nil {
		fmt.Println("Error running query")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer rows.Close()

	for rows.Next() {
		p := NewPokemon(rows)
		pokedex = append(pokedex, *p)
	}
	json.NewEncoder(w).Encode(pokedex)
}
