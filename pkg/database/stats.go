package database

import (
	"context"
	"mutabornaya/global"
	"mutabornaya/pkg/models"
	"strconv"
)

func (db *Database) UpdateStats(chanID int64, stats models.Stats) error {
	_, err := db.db.Collection(global.ColStats).Doc(strconv.FormatInt(chanID, 10)).Set(context.Background(), stats)

	return err
}

func (db *Database) GetStats(chanID int64) (currState *models.Stats, err error) {
	dsnap, err := db.db.Collection(global.ColStats).Doc(strconv.FormatInt(chanID, 10)).Get(context.Background())
	if err != nil {
		return nil, err
	}
	dsnap.DataTo(&currState)

	return
}
