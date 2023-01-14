package models

import "time"

type Chats struct {
	IsMember        bool      `firestore:"is_member"`
	IsAdmin         bool      `firestore:"is_admin"`
	IsPidorEnabled  bool      `firestore:"is_pidor_enabled"`
	IsToxicEnabled  bool      `firestore:"is_toxic_enabled"`
	IsAnimalEnabled bool      `firestore:"is_animal_enabled"`
	JoinedAt        time.Time `firestore:"joined_at"`
	LeavedAt        time.Time `firestore:"leaved_at,omitempty"`
}
