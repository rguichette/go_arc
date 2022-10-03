package main

import (
	"github.com/rguichette/go_arc"
	"github.com/rguichette/go_arc/storage/mongo"
	"github.com/rguichette/go_arc/storage/postgres"
)

func main() {
	dbm := mongo.Mongo{}
	dbp := postgres.Postg{}

	p := go_arc.Person{
		First: "Ralph",
	}

	go_arc.Put(dbm, 1, p)

	go_arc.Get(dbp, 2)

}
