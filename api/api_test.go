package api

import (
	"database/sql"
	"net/http"
	"net/url"
	"os"
)

func getNewRequest(resource string) (*http.Request, error) {
	u, _ := url.ParseRequestURI("http://localhost:12345")
	u.Path = resource
	urlStr := u.String()
	return http.NewRequest("GET", urlStr, nil)
}

func init() {
	os.Remove("multidex.db")
	var err error
	db, err = sql.Open("sqlite3", "multidex.db")
	if err != nil {
		panic("Failed to open database")
	}
	defer db.Close()

	filename := "../assets/pokemon.csv"
	tableStr := `
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
	prepareStr := "insert into pokemon (Number, Name, HP, Attack, Defense, SAttack, SDefense, Speed, TypeA, TypeB) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	if err = ReadDatabase(filename, tableStr, prepareStr); err != nil {
		panic("Pokemon database not set up properly!")
	}

	filename = "../assets/attacks.csv"
	tableStr = `
    create table attacks (
        Name TEXT,
        Type TEXT,
        Category TEXT,
        PP INT,
        Power INT,
        Accuracy INT
    )
    `
	prepareStr = "insert into attacks (Name, Type, Category, PP, Power, Accuracy) values (?, ?, ?, ?, ?, ?)"
	if err = ReadDatabase(filename, tableStr, prepareStr); err != nil {
		panic("Attack database not set up properly!")
	}
}
