package models

import "gorm.io/gorm"

type JobVacancy struct {
	gorm.Model
	IDLowongan     uint        `gorm:"primaryKey;autoIncrement" json:"id_lowongan"`
	IDPerusahaan   uint        `gorm:"not null" json:"id_perusahaan"`
	NamaJabatan    string      `gorm:"not null" json:"nama_jabatan"`
	JadwalSeleksi  string      `gorm:"not null" json:"jadwal_seleksi"`
	KotaPenempatan string      `gorm:"not null" json:"kota_penempatan"`
	Persyaratan    string      `gorm:"not null" json:"persyaratan"`
	Status         string      `gorm:"not null;" json:"status"`
	Gaji           string      `gorm:"not null" json:"gaji"`
	Company        Company     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:IDPerusahaan;references:IDPerusahaan"`
	Applicants     []Applicant `gorm:"foreignKey:IDLowongan;references:IDLowongan"`
}

func (JobVacancy) TableName() string {
	return "tb_lowongan"
}
