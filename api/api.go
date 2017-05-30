package api

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "net/http"
    "time"

    "github.com/gorilla/mux"
    _ "github.com/mattn/go-sqlite3"
)

var db sql.DB

func runQuery(query string) (*sql.Rows, error) {
    db, err := sql.Open("sqlite3", "multidex.db")
    if err != nil {
        fmt.Println("Failed to open database")
        return nil, err
    }
    defer db.Close()
    return db.Query(query)
}

func setup() {
    os.Remove("multidex.db")
    db, err := sql.Open("sqlite3", "multidex.db")
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

func StartAPI() {
    setup()

    router := mux.NewRouter()
    router.HandleFunc("/pokemon", GetPokedex).Methods("GET")
    router.HandleFunc("/pokemon/{name}", GetPokemonFromPokedex).Methods("GET")
    router.HandleFunc("/pokemon/type/{type}", GetPokemonByType).Methods("GET")
    router.HandleFunc("/attacks", GetAttacks).Methods("GET")
    router.HandleFunc("/attacks/{name}", GetAttackFromAttacks).Methods("GET")
    router.HandleFunc("/attacks/type/{type}", GetAttacksByType).Methods("GET")
    router.HandleFunc("/types", GetTypes).Methods("GET")
    log.Fatal(http.ListenAndServe(":12345", router))
}
