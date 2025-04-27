package model

import "time"

type Registration struct {
	ID                  int64
	UserID              int64
	MicrophoneStatus    int32
	ParticipationDevice string
	WishedRewards       string
	RegisteredAt        time.Time
}
