package models

type CepatTelitiQuestion struct {
	IDSoal    uint   `gorm:"primaryKey" json:"id_soal"`
	NomorSoal int    `json:"no_soal"`
	Soal      string `json:"soal"`
	OpsiA     string `json:"opsi_a"`
	OpsiB     string `json:"opsi_b"`
	OpsiC     string `json:"opsi_c"`
	OpsiD     string `json:"opsi_d"`
	OpsiE     string `json:"opsi_e"`
	Jawaban   string `json:"jawaban"`
}

func (CepatTelitiQuestion) TableName() string {
	return "tb_soal_cepat"
}
