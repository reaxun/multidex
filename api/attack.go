package api

import (
    "database/sql"
    "encoding/csv"
    "encoding/json"
    "fmt"
    "os"
    "net/http"
    "strings"
    "time"

    "github.com/gorilla/mux"
)

type Attack struct {
    Name        string `json:"name,omitempty"`
    Type        string `json:"type,omitempty"`
    Category    string `json:"category,omitempty"`
    PP          int `json:"pp,omitempty"`
    Power       int `json:"pow,omitempty"`
    Accuracy    int `json:"acc,omitempty"`
}

func NewAttack(rows *sql.Rows) *Attack {
    var name, atkType, category string
    var pp, power, accuracy int
    rows.Scan(&name, &atkType, &category, &pp, &power, &accuracy)
    return &Attack{Name: name, Type: atkType, Category: category, PP: pp, Power: power, Accuracy: accuracy}
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
            s[i] = v
        }
        if _, err = stmt.Exec(s...); err != nil {
            fmt.Println("Exec failed")
            return err
        }
    }
    return tx.Commit()
}

func GetAttackFromAttacks(w http.ResponseWriter, req *http.Request) {
    searchStartTime := time.Now()
    params := mux.Vars(req)
    name := strings.Title(params["name"])
    rows, err := runQuery("select * from attacks where Name='" + name + "'")
    if err != nil {
        fmt.Println("Error running query")
        return
    }
    defer rows.Close()

    for rows.Next() {
        a := NewAttack(rows)
        searchDuration := time.Since(searchStartTime)
        fmt.Println("Took " + searchDuration.String() + " to search Attack database")
        json.NewEncoder(w).Encode(a)
        return
    }
    json.NewEncoder(w).Encode(&Attack{})
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
