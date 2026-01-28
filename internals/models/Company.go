package models

type Company struct {
	IDPerusahaan   uint         `gorm:"primaryKey;autoIncrement" json:"id_perusahaan"`
	NamaPerusahaan string       `gorm:"not null" json:"nama_perusahaan"`
	Alamat         string       `gorm:"not null" json:"alamat"`
	JenisUsaha     string       `gorm:"not null" json:"jenis_usaha"`
	Email          string       `gorm:"not null" json:"email"`
	NoHp           string       `gorm:"not null" json:"no_hp"`
	LogoPerusahaan string       `json:"logo_perusahaan"`
	Status         int          `gorm:"not null;" json:"status"`
	JobVacancies   []JobVacancy `gorm:"foreignKey:IDPerusahaan"`
}

func (Company) TableName() string {
	return "tb_perusahaan"
}
