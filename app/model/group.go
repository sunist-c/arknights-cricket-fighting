package model

import "time"

type Group struct {
	ID                 int64     `gorm:"column:id;type:integer;primaryKey;autoIncrement"`
	GroupNumber        int64     `gorm:"column:group_number;type:integer;not null;uniqueIndex:idx_group_number"`
	GroupName          string    `gorm:"column:group_name;type:varchar(255);not null;index:idx_group_name"`
	Description        string    `gorm:"column:description;type:varchar(255);not null"`
	ManagerRealName    string    `gorm:"column:manager_real_name;type:varchar(32);not null"`
	ManagerQQNumber    int64     `gorm:"column:manager_qq_number;type:integer;not null"`
	ManagerOfficialID  string    `gorm:"column:manager_official_id;type:varchar(18);not null"`
	ManagerPhoneNumber string    `gorm:"column:manager_phone_number;type:varchar(32);not null"`
	CreatedAt          time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt          time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
}
