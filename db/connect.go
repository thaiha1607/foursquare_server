package db

import "github.com/thaiha1607/foursquare_server/ent"

type DatabaseConnector interface {
	setupConnection() (db *ent.Client)
}

func SetupDbConnection(dbc DatabaseConnector) (db *ent.Client) {
	db = dbc.setupConnection()
	return
}
