package models

import "time"

type UserStats struct {
	Times  int   `firestore:"time"`
	UserId int64 `firestore:"user_id"`
	Streak int   `firestore:"streak"`
}

type ChatStats struct {
	IsEnabled bool        `firestore:"is_enabled"`
	Users     []UserStats `firestore:"users,omitempty"`
}

type Stats struct {
	Pidor       ChatStats `firestore:"pidor"`
	Toxic       ChatStats `firestore:"toxic"`
	Animal      ChatStats `firestore:"animals"`
	LastUpdated time.Time `firestore:"last_updated"`
}
