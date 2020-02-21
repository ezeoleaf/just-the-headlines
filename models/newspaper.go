package models

import (
	"database/sql"
	"fmt"
)

type Newspaper struct {
	ID      int    `json:"id"`
	Name    string `json:"id"`
	Country string `json:"id"`
}

type Newspapers struct {
	Newspapers []Newspaper `json:"items"`
}

func NewNewspapers() *Newspapers {
	n := &Newspapers{}

	return n
}

func (n *Newspapers) Fetch(db *sql.DB) (self *Newspapers) {
	self = n
	sql := "SELECT * FROM newspaper"
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	result := &Newspapers{}
	for rows.Next() {
		n := Newspaper{}
		e := rows.Scan(&n.ID, &n.Name, &n.Country)
		fmt.Println(n)
		if e != nil {
			panic(err)
		}
		result.Newspapers = append(result.Newspapers, n)
	}

	return result
}
