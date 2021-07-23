package db

type ApiLog struct {
	ID     uint   `gorm:"primaryKey"`
	Api    string `gorm:"type:varchar(32);not null"`
	Status string `gorm:"type:varchar(16);not null"`
	Time   string `gorm:"type:varchar(16);not null"`
}
