package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func DbConnect(data Data) {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "people",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	albID, err := AddPeople(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added ppl: %v\n", albID)
}

func GetData() ([]Data, error) {
	var data []Data
	rows, err := db.Query("SELECT * FROM ppl")
	if err != nil {
		return nil, fmt.Errorf("getData: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var d Data
		if err := rows.Scan(&d.ID, &d.Name, &d.Email, &d.Phone, &d.Message); err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("getData: %v", err)
		}
		data = append(data, d)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getData: %v", err)
	}
	fmt.Printf("Data: %v\n", data)
	return data, nil
}

func AddPeople(data Data) (int64, error) {
	result, err := db.Exec("INSERT INTO ppl (name, email, phone, msg) VALUES (?, ?, ?, ?)", data.Name, data.Email, data.Phone, data.Message)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}
