package model

import "time"

type User struct {
	ID             int64
	QQUnionID      string
	QQNumber       int64
	LinkedGroupID  int64
	AvatarURL      string
	InGameNickname string
	GameChannel    string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
