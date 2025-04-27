package model

import "time"

type PreliminaryPoints struct {
	ID        int64     `gorm:"column:id;type:integer;primaryKey;autoIncrement:true"`
	UserID    int64     `gorm:"column:user_id;type:integer;not null;uniqueIndex:idx_user_id"`
	Points    int64     `gorm:"column:points;type:integer;not null"`
	TotalRank int64     `gorm:"column:total_rank;type:integer;not null"`
	Advanced  bool      `gorm:"column:advanced;type:boolean;not null"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
}
