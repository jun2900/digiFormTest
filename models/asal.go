package models

import "gorm.io/gorm"

type Lokasi struct {
	gorm.Model
	ProvinsiAsal    string `json:"provinsi_asal" gorm:"not null;size:50"`
	TipeAsal        string `json:"tipe_asal" gorm:"not null;size:50"`
	KotaAsal        string `json:"kota/kabupaten_asal" gorm:"not null;size:50"`
	KecamatanAsal   string `json:"kecamatan_asal" gorm:"not null;size:50"`
	KodeposAsal     string `json:"kodepos_asal" gorm:"not null;size:50"`
	AlamatAsal      string `json:"alamat_asal" gorm:"not null;size:50"`
	ProvinsiTujuan  string `json:"provinsi_tujuan" gorm:"not null;size:50"`
	TipeTujuan      string `json:"tipe_tujuan" gorm:"not null;size:50"`
	KotaTujuan      string `json:"kota/kabupaten_tujuan" gorm:"not null;size:50"`
	KecamatanTujuan string `json:"kecamatan_tujuan" gorm:"not null;size:50"`
	KodeposTujuan   string `json:"kodepos_tujuan" gorm:"not null;size:50"`
	AlamatTujuan    string `json:"alamat_tujuan" gorm:"not null;size:50"`
	Service         string `json:"service" gorm:"not null;size:50"`
}
