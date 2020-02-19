package jth

import (
	"database/sql"
	"fmt"

	"github.com/ezeoleaf/just-the-headlines/models"
)

func Fetch(db *sql.DB) *models.NewspaperCollection {
	n := models.GetNewspapers(DB)
	fmt.Println(n)
}
