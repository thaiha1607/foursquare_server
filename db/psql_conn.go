package db

import (
	"context"
	"log"

	_ "github.com/lib/pq"
	"github.com/thaiha1607/foursquare_server/common/env"
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
	// >>> AUTOMATIC MIGRATION START
	// It should be used only for development purpose.
	// USE IT WITH CAUTION!
	if !env.IsProdEnv() {
		ctx := context.Background()
		if err := db.Schema.Create(ctx); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
		log.Println("automatic migration done")
	}
	// AUTOMATIC MIGRATION END <<<
	return
}
