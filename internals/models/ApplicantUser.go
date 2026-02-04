package models

type ApplicantUser struct {
	IDPelamar  uint        `gorm:"primaryKey;autoIncrement" json:"id_pelamar"`
	IDLevel    uint        `json:"id_level"`
	Email      string      `gorm:"unique;not null" json:"email"`
	Username   string      `gorm:"unique;not null" json:"username"`
	Password   string      `gorm:"not null" json:"-"`
	Foto       string      `json:"foto"`
	Status     int         `gorm:"not null;" json:"status"`
	UserData   UserData    `gorm:"foreignKey:IDPelamar"`
	Applicants []Applicant `gorm:"foreignKey:IDApply"`
}

func (ApplicantUser) TableName() string {
	return "tb_pelamar"
}
