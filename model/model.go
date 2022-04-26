package model

import "time"

type Country struct {
	Id      int64     `gorm:"type:BIGINT UNSIGNED AUTO_INCREMENT" json:"id"`
	Created time.Time `gorm:"type:DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL;autoCreateTime" json:"created"`
	Updated time.Time `gorm:"type:DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL;autoUpdateTime" json:"updated"`
	Name    string    `gorm:"type:VARCHAR(100) NOT NULL;index" json:"name"`
	Keyword string    `gorm:"type:VARCHAR(5) NOT NULL;index" json:"keyword"`
}

type WebRequest struct {
	Id        int64     `gorm:"type:BIGINT UNSIGNED AUTO_INCREMENT" json:"id"`
	Created   time.Time `gorm:"type:DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL;autoCreateTime" json:"created"`
	Updated   time.Time `gorm:"type:DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL;autoUpdateTime" json:"updated"`
	CountryId int64     `gorm:"type:BIGINT UNSIGNED NOT NULL" json:"country_id"`
	DeviceId  int64     `gorm:"type:BIGINT UNSIGNED NOT NULL" json:"device_id"`
	Count     int64     `gorm:"type:BIGINT UNSIGNED NOT NULL;default:0" json:"count"`
}

type TimeSpent struct {
	Id        int64     `gorm:"type:BIGINT UNSIGNED AUTO_INCREMENT" json:"id"`
	Created   time.Time `gorm:"type:DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL;autoCreateTime" json:"created"`
	Updated   time.Time `gorm:"type:DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL;autoUpdateTime" json:"updated"`
	CountryId int64     `gorm:"type:BIGINT UNSIGNED NOT NULL" json:"country_id"`
	DeviceId  int64     `gorm:"type:BIGINT UNSIGNED NOT NULL" json:"device_id"`
	Count     int64     `gorm:"type:BIGINT UNSIGNED NOT NULL;default:0" json:"count"`
}

type Device struct {
	Id      int64     `gorm:"type:BIGINT UNSIGNED AUTO_INCREMENT" json:"id"`
	Created time.Time `gorm:"type:DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL;autoCreateTime" json:"created"`
	Updated time.Time `gorm:"type:DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL;autoUpdateTime" json:"updated"`
	Name    string    `gorm:"type:VARCHAR(50) NOT NULL" json:"name"`
}
