package database

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type Database struct {
	db *firestore.Client
}

func NewDatabase(sa option.ClientOption) (*Database, error) {
	app, err := firebase.NewApp(context.Background(), nil, sa)
	if err != nil {
		return nil, err
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		return nil, err
	}

	return &Database{
		db: client,
	}, nil
}

func (db *Database) Close() {
	db.db.Close()
}
