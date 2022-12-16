package models
// naming convention should be plural or singular?

import "gorm.io/gorm"

type Bank struct {
	gorm.Model
	Name 			string
	IBAN 			string
}

// // type User struct {
//   ID        uint           `gorm:"primaryKey"`
//   CreatedAt time.Time
//   UpdatedAt time.Time
//   DeletedAt gorm.DeletedAt `gorm:"index"`
//   Name string
// }
