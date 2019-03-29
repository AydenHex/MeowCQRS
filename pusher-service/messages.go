package main

import (
	"time"
)

const (
	kindMeowCreated = iota + 1
)

type MeowCreatedMessage struct {
	Kind      uint32    `json:"king"`
	ID        string    `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

func newMeowCreatedMessage(id string, body string, createdAt time.Time) *MeowCreatedMessage {
	return &MeowCreatedMessage{
		Kind:      kindMeowCreated,
		ID:        id,
		Body:      body,
		CreatedAt: createdAt,
	}
}
