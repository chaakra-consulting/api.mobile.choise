package models

import "gorm.io/gorm"

type ApplicantUser struct {
	gorm.Model
	IDPelamar  uint        `gorm:"primaryKey;autoIncrement" json:"id_pelamar"`
	IDLevel    uint        `json:"id_level"`
	Username   string      `gorm:"unique;not null" json:"username"`
	Password   string      `gorm:"not null" json:"-"`
	Foto       string      `json:"foto"`
	Status     int         `gorm:"not null;" json:"status"`
	Applicants []Applicant `gorm:"foreignKey:IDApply"`
}

func (ApplicantUser) TableName() string {
	return "tb_pelamar"
}
