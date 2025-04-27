package model

import "time"

type SemifinalPoints struct {
	ID               int64
	SemifinalGroupID int64
	Users            string
	GroupPoints      int64
	Advanced         bool
	TotalRank        int64
	CreatedAt        time.Time
}
