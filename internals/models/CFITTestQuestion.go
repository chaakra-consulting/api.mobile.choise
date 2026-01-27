package models

type CFITTestQuestion struct {
	IDSoal    int    `gorm:"primaryKey" json:"id_soal"`
	NomorSoal int    `json:"nomor_soal"`
	Soal      string `json:"soal"`
	OpsiA     string `json:"opsi_a"`
	OpsiB     string `json:"opsi_b"`
	OpsiC     string `json:"opsi_c"`
	OpsiD     string `json:"opsi_d"`
	OpsiE     string `json:"opsi_e"`
	OpsiF     string `json:"opsi_f"`
	Jawaban   string `json:"jawaban"`
	Jawaban2  string `json:"jawaban_2"`
	TypeSoal  string `json:"type_soal"`
	Subtes    int    `json:"subtes"`
}

func (CFITTestQuestion) TableName() string {
	return "tb_soal_cfit"
}
