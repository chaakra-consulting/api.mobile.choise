package models

import "gorm.io/gorm"

type Applicant struct {
	gorm.Model
	IDApply       uint          `gorm:"primaryKey;autoIncrement" json:"id_apply"`
	IDPelamar     uint          ` json:"id_pelamar"`
	IDLowongan    uint          ` json:"id_lowongan"`
	IDPerusahaan  int           ` json:"id_perusahaan"`
	StatusLamaran string        ` json:"status_lamaran"`
	StatusUjian   string        ` json:"status_ujian"`
	JobVacancy    JobVacancy    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:IDLowongan"`
	ApplicantUser ApplicantUser `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:IDPelamar"`
}

func (Applicant) TableName() string {
	return "tb_apply"
}
