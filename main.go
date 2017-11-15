package main

import (
	"log"

	"github.com/TMDeal/PokeDB/models"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	db, err := models.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
