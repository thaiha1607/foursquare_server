package db

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/thaiha1607/foursquare_server/ent"
)

type PostgresDatabaseConnector struct {
	uri string
}

func NewPostgresDatabaseConnector(uri string) DatabaseConnector {
	return &PostgresDatabaseConnector{
		uri: uri,
	}
}

func (p *PostgresDatabaseConnector) setupConnection() (db *ent.Client) {
	db, err := ent.Open("postgres", p.uri)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	log.Println("connected to postgres")
	return
}
