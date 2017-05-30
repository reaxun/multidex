package api

import (
    "database/sql"
    "encoding/csv"
    "encoding/json"
    "fmt"
    "os"
    "net/http"
    "strings"

    "github.com/gorilla/mux"
)

type Stats struct {
    HP          int `json:"hp,omitempty`
    Attack      int `json:"attack,omitempty`
    Defense     int `json:"defense,omitempty`
    SAttack     int `json:"sattack,omitempty`
    SDefense    int `json:"sdefense,omitempty`
    Speed       int `json:"speed,omitempty`
}

type Pokemon struct {
    Number      int     `json:"number,omitempty"`
    Name        string  `json:"name,omitempty"`
    Base        *Stats  `json:"basestats,omitempty"`
    EVs         *Stats  `json:"evs,omitempty"`
    IVs         *Stats  `json:"ivs,omitempty"`
    TypeA       string  `json:"typea,omitempty"`
    TypeB       string  `json:"typeb,omitempty"`
}

func NewPokemon(rows *sql.Rows) *Pokemon {
    var number, hp, atk, def, satk, sdef, spd int
    var name, typea, typeb string
    rows.Scan(&number, &name, &hp, &atk, &def, &satk, &sdef, &spd, &typea, &typeb)
    return &Pokemon{Number: number, Name: name, Base: &Stats{HP: hp, Attack: atk, Defense: def, SAttack: satk, SDefense: sdef, Speed: spd}, TypeA: typea, TypeB: typeb}
}

func ReadPokemonDatabase(db *sql.DB) error {
    f, err := os.Open("assets/pokemon.csv")
    if err != nil {
        return err
    }
    defer f.Close()

    csvr := csv.NewReader(f)
    rows, err := csvr.ReadAll()

    table := `
    create table pokemon (
        Number INT,
        Name TEXT,
        HP INT,
        Attack INT,
        Defense INT,
        SAttack INT,
        SDefense INT,
        Speed INT,
        TypeA TEXT,
        TypeB TEXT
    )
    `
    if _, err := db.Exec(table); err != nil {
        return err
    }

    tx, err := db.Begin()
    if err != nil {
        return err
    }

    stmt, err := tx.Prepare("insert into pokemon (Number, Name, HP, Attack, Defense, SAttack, SDefense, Speed, TypeA, TypeB) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
    if err != nil {
        return err
    }
    defer stmt.Close()

    for _, row := range rows {
        s := make([]interface{}, len(row))
        for i, v := range row {
            s[i] = v
        }
        if _, err = stmt.Exec(s...); err != nil {
            fmt.Println("Exec failed")
            return err
        }
    }
    return tx.Commit()
}

func GetPokemonFromPokedex(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    name := strings.Title(params["name"])
    rows, err := runQuery("select * from pokemon where Name='" + name + "'")
    if err != nil {
        fmt.Println("Error running query")
        return
    }
    defer rows.Close()

    for rows.Next() {
        p := NewPokemon(rows)
        json.NewEncoder(w).Encode(p)
        return
    }
    json.NewEncoder(w).Encode(&Pokemon{})
}

func GetPokemonByType(w http.ResponseWriter, req *http.Request) {
    var pokedex []Pokemon
    params := mux.Vars(req)
    pokemonType := strings.Title(params["type"])
    rows, err := runQuery("select * from pokemon where TypeA='" + pokemonType + "' OR TypeB='" + pokemonType + "'")
    if err != nil {
        fmt.Println("Error running query")
        return
    }
    defer rows.Close()

    for rows.Next() {
        p := NewPokemon(rows)
        pokedex = append(pokedex, *p)
    }
    json.NewEncoder(w).Encode(pokedex)
}

func GetPokedex(w http.ResponseWriter, req *http.Request) {
    var pokedex []Pokemon
    rows, err := runQuery("select * from pokemon")
    if err != nil {
        fmt.Println("Error running query")
        return
    }
    defer rows.Close()

    for rows.Next() {
        p := NewPokemon(rows)
        pokedex = append(pokedex, *p)
    }
    json.NewEncoder(w).Encode(pokedex)
}
