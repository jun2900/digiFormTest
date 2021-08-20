package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type AirwayBill struct {
	gorm.Model
	SendDate      datatypes.Date `json:"send_date" gorm:"not null;size:50"`
	Service       string         `json:"tipe" gorm:"not null;size:50"`
	Origin        string         `json:"origin" gorm:"not null;size:50"`
	KodeposTujuan string         `json:"kodepos" gorm:"not null;size:50"`
}
