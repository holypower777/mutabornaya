package database

import (
	"context"
	"mutabornaya/global"
	"mutabornaya/pkg/models"
	"strconv"
	"time"
)

func (db *Database) AddChat(chatID int64) error {
	_, err := db.db.Collection(global.ColChats).Doc(strconv.FormatInt(chatID, 10)).Set(context.Background(), &models.Chats{
		IsMember:        true,
		IsAdmin:         false,
		IsPidorEnabled:  false,
		IsToxicEnabled:  false,
		IsAnimalEnabled: false,
		JoinedAt:        time.Now(),
	})

	return err
}
