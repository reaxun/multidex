package api

import (
	"database/sql"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var (
	router *mux.Router
)

func getNewRequest(resource string) (*http.Request, error) {
	u, _ := url.ParseRequestURI("http://localhost:12345")
	u.Path = resource
	urlStr := u.String()
	return http.NewRequest("GET", urlStr, nil)
}

func openFakeDB() {
	db.Close()
	db, _ = sql.Open("sqlite3", "test.db")
}

func closeFakeDB() {
	db.Close()
	db, _ = sql.Open("sqlite3", "multidex.db")
}

func init() {
	os.Remove("multidex.db")
	var err error
	db, err = sql.Open("sqlite3", "multidex.db")
	if err != nil {
		panic("Failed to open database")
	}

	filename := "../assets/pokemon.csv"
	tableStr := "create table pokemon (Number INT, Name TEXT, HP INT, Attack INT, Defense INT, SAttack INT, SDefense INT, Speed INT, TypeA TEXT, TypeB TEXT)"
	prepareStr := "insert into pokemon (Number, Name, HP, Attack, Defense, SAttack, SDefense, Speed, TypeA, TypeB) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	if err = ReadDatabase(filename, tableStr, prepareStr); err != nil {
		panic("Pokemon database not set up properly!")
	}

	filename = "../assets/attacks.csv"
	tableStr = "create table attacks (Name TEXT, Type TEXT, Category TEXT, PP INT, Power INT, Accuracy INT)"
	prepareStr = "insert into attacks (Name, Type, Category, PP, Power, Accuracy) values (?, ?, ?, ?, ?, ?)"
	if err = ReadDatabase(filename, tableStr, prepareStr); err != nil {
		panic("Attack database not set up properly!")
	}

	router = mux.NewRouter()
	addRoutes(router)
}

func TestReadDatabaseFailures(t *testing.T) {
	old_db := db
	var err error
	filename := "../assets/pokemon.csv"
	tableStr := "create table test (Number INT, Name TEXT)"
	prepareStr := "insert into test (Number, Name) values (?, ?)"
	tables := []struct {
		f string
		t string
		p string
		e string
	}{
		{"badfilename", tableStr, prepareStr, "open badfilename: no such file or directory"},
		{filename, "badtablestring", prepareStr, `near "badtablestring": syntax error`},
		{filename, tableStr, "badpreparestring", `near "badpreparestring": syntax error`},
		{filename, tableStr, prepareStr, "sql: expected 2 arguments, got 10"},
	}

	for _, table := range tables {
		os.Remove("test.db")
		db, err = sql.Open("sqlite3", "test.db")
		err = ReadDatabase(table.f, table.t, table.p)
		assert.Error(t, err)
		assert.Equal(t, table.e, err.Error())
		db.Close()
	}
	db = old_db
}
