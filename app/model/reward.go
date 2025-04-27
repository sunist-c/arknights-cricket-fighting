package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type Reward struct {
	ID          int64
	RewardName  string
	RewardCount int64
	WishedCount int64
	RewardLink  string
	RewardPrice decimal.Decimal
	CreatedAt   time.Time
}
