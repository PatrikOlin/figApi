package datastore

import (	
	"bufio"
	"database/sql"
	"fmt"	
	"log"
	"os"
	"figApi/util"
			
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/lib/pq"
)

	 
func Initdb() {

	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	// host, port, user, password, dbname)
	
	// db, err := sql.Open("postgres", psqlInfo)
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()

	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }

	// createDb()
	// initFirstnameTable()
	// initSurnameTable()
	// initStreetPrefixTable()
	// initStreetSuffixTable()
	// initPostalAddressTable()
	// initPasswordTable()
	// initEmailDomainsTable()
	// initCompanynameTable()
	// initArticlesTable()
}

// func getConn() string {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
// 	host, port, user, password, dbname)

// 	return psqlInfo;
// }

func createDb() {
	path := "./datastore/store.db"
	var _, err = os.Stat(path)

	if os.IsNotExist(err) {
		var file, err = os.Create(path)
	 	util.Check(err)
		defer file.Close()
	}
}

func initFirstnameTable() {
	src := "./datastore/fname"
	dropStmt := "DROP TABLE IF EXISTS firstnames"
	initStmt := "CREATE TABLE IF NOT EXISTS firstnames (id INTEGER PRIMARY KEY, firstname TEXT)"
	inputStmt := "INSERT INTO firstnames (firstname) VALUES (?)"

	initTable(dropStmt, initStmt)
	populateTable(src, inputStmt)
}

func initSurnameTable() {
	src := "./datastore/lname"
	dropStmt := "DROP TABLE IF EXISTS surnames"
	initStmt := "CREATE TABLE IF NOT EXISTS surnames (id INTEGER PRIMARY KEY, surname TEXT)"
	inputStmt := "INSERT INTO surnames (surname) VALUES (?)"

	initTable(dropStmt, initStmt)
	populateTable(src, inputStmt)
}

func initStreetPrefixTable() {
	src := "./datastore/stpre"
	dropStmt := "DROP TABLE IF EXISTS streetprefixes"
	initStmt := "CREATE TABLE IF NOT EXISTS streetprefixes (id INTEGER PRIMARY KEY, streetprefix TEXT)"
	inputStmt := "INSERT INTO streetprefixes (streetprefix) VALUES (?)"

	initTable(dropStmt, initStmt)
	populateTable(src, inputStmt)
}

func initStreetSuffixTable() {
	src := "./datastore/stsuf"
	dropStmt := "DROP TABLE IF EXISTS streetsuffixes"
	initStmt := "CREATE TABLE IF NOT EXISTS streetsuffixes (id INTEGER PRIMARY KEY, streetsuffix TEXT)"
	inputStmt := "INSERT INTO streetsuffixes (streetsuffix) VALUES (?)"

	initTable(dropStmt, initStmt)
	populateTable(src, inputStmt)
}

func initPasswordTable() {
	src := "./datastore/pwords"
	dropStmt := "DROP TABLE IF EXISTS passwords"
	initStmt := "CREATE TABLE IF NOT EXISTS passwords (id INTEGER PRIMARY KEY, password TEXT)"
	inputStmt := "INSERT INTO passwords (password) VALUES (?)"

	initTable(dropStmt, initStmt)
	populateTable(src, inputStmt)
}

func initEmailDomainsTable() {
	src := "./datastore/emailaddresses"
	dropStmt := "DROP TABLE IF EXISTS emaildomains"
	initStmt := "CREATE TABLE IF NOT EXISTS emaildomains (id INTEGER PRIMARY KEY, emaildomain TEXT)"
	inputStmt := "INSERT INTO emaildomains (emaildomain) VALUES (?)"

	initTable(dropStmt, initStmt)
	populateTable(src, inputStmt)
}

func initArticlesTable() {
	src := "./datastore/articles"
	dropStmt := "DROP TABLE IF EXISTS articles"
	initStmt := "CREATE TABLE IF NOT EXISTS articles (id INTEGER PRIMARY KEY, article TEXT)"
	inputStmt := "INSERT INTO articles (article) VALUES (?)"

	initTable(dropStmt, initStmt)
	populateTable(src, inputStmt)
}

func initCompanynameTable() {
	src := "./datastore/companynames"
	dropStmt := "DROP TABLE IF EXISTS companynameparts"
	initStmt := "CREATE TABLE IF NOT EXISTS companynameparts (id INTEGER PRIMARY KEY, companynamepart TEXT)"
	inputStmt := "INSERT INTO companynameparts (companynamepart) VALUES (?)"

	initTable(dropStmt, initStmt)
	file, err := os.Open(src)
	util.Check(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	db, _ := sql.Open("sqlite3", "./datastore/store.db")
	statement, _ := db.Prepare(inputStmt)
	for scanner.Scan() {
		statement.Exec(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	statement.Close()
	db.Close()
}

func initPostalAddressTable() {
	src := "./datastore/pnumort"
	dropStmt := "DROP TABLE IF EXISTS postaladdresses"
	initStmt := "CREATE TABLE IF NOT EXISTS postaladdresses (id INTEGER PRIMARY KEY, postalcode TEXT, posttown TEXT)"
	inputStmt := "INSERT INTO postaladdresses (postalcode, posttown) VALUES (?, ?)"

	initTable(dropStmt, initStmt)

	file, err := os.Open(src)
	util.Check(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	db, _ := sql.Open("sqlite3", "./datastore/store.db")
	statement, _ := db.Prepare(inputStmt)
	var line []string
	for scanner.Scan() {
		line = append(line, scanner.Text())
		if len(line) == 2 {
			statement.Exec(line[0], line[1])
			line = line[2:]
		}
	}
	statement.Close()
	db.Close()

}

func initTable(dropStmt string, initStmt string) {
	db, _ := sql.Open("sqlite3", "./datastore/store.db")
	statement, _ := db.Prepare(dropStmt)
	statement.Exec()
	statement, _ = db.Prepare(initStmt)
	statement.Exec()
	statement.Close()

	db.Close()
}

func populateTable(sourcefile string, inputStmt string) {
	file, err := os.Open(sourcefile)
	util.Check(err)

	scanner := bufio.NewScanner(file)
	db, _ := sql.Open("sqlite3", "./datastore/store.db")
	statement, _ := db.Prepare(inputStmt)
	for scanner.Scan() {
		statement.Exec(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	statement.Close()
	db.Close()
}

func GetRandomLine(tablename string) string {
	var result string
	var id int
	// db, _ := sql.Open("sqlite3", "./datastore/store.db")
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load env-file")
	}
	db, err := sql.Open("postgres", os.ExpandEnv("host=${HOST} user=${USER} dbname=${DBNAME} sslmode=disable password=${PASSWORD}"))
	if err != nil {
		panic("Failed to connect to db")
	}
	
	row := db.QueryRow("SELECT * FROM " + tablename + " ORDER BY random() LIMIT 1;")
	error := row.Scan(&result, &id)

	if error != nil {
		if error == sql.ErrNoRows {
			fmt.Println("No rows found")
		} else {
			panic(error)
		}
	}

	db.Close()
	return result
}
			
	 
