package model

import "time"

type Address struct {
	ID           int64     `gorm:"column:id;type:integer;primaryKey;autoIncrement:true"`
	UserID       int64     `gorm:"column:user_id;type:integer;not null;uniqueIndex:idx_user_id"`
	ReceiverName string    `gorm:"column:receiver_name;type:varchar(32);not null;index:idx_receiver_name"`
	PhoneNumber  string    `gorm:"column:phone_number;type:varchar(32);not null;index:idx_phone_number"`
	CountryCode  string    `gorm:"column:country_code;type:varchar(3);not null;default:'PRC'"`
	Province     string    `gorm:"column:province;type:varchar(32):not null"`
	City         string    `gorm:"column:city;type:varchar(32):not null"`
	District     string    `gorm:"column:district;type:varchar(32):not null"`
	AddressLine  string    `gorm:"column:address_line;type:varchar(255);not null"`
	PostCode     string    `gorm:"column:post_code;type:varchar(32);not null"`
	CreatedAt    time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
}
