package model

import (
	"strings"
	"time"
	"gorm.io/gorm"
)

// CarOwner構造体
type CarOwner struct {
	ID                uint		`json:"id" gorm:"primaryKey"`
	FirstName         string	`json:"first_name" gorm:"column:first_name;type:varchar(25);not null"`
	MiddleName        string	`json:"middle_name" gorm:"column:middle_name;type:varchar(25)"`
	LastName          string	`json:"last_name" gorm:"column:last_name;type:varchar(25)"`
	LicenseExpiration time.Time `json:"license_expiration" gorm:"column:license_expiration;type:date;not null"`
	CreatedAt         time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt         time.Time	`json:"updated_at" gorm:"autoUpdateTime"`
    DeletedAt         gorm.DeletedAt	`json:"-" gorm:"index"`
}

//テーブル名を明示的に指定
func (CarOwner) TableName() string {
	return "car_owners"
}

// 免許証の期限切れ確認
func (c *CarOwner) IsLicenseExpired() bool {
	//現在時刻が c.LicwenseExpiration より後かどうか(現在時刻が後ならtrueを返す)
	return time.Now().After(c.LicenseExpiration)
}

//name枠に該当する名前があるか確認する
func (c *CarOwner) IsContainsName (name string) bool {
	return strings.Contains(c.FirstName, name) ||
		strings.Contains(c.MiddleName, name) ||
		strings.Contains(c.LastName, name)
}

//IDが正の数かどうか確認する
func (c *CarOwner)IsIDPositive() bool {
    if c.ID <= 0 {
        return false
    }
    return true
}

