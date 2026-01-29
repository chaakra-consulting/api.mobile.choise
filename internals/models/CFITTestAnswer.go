package models

type CFITTestAnswer struct {
	IDJawabanCFIT uint   `gorm:"primaryKey;autoIncrement" json:"id_jawaban_cfit"`
	IDPelamar     uint   `json:"id_pelamar"`
	IDLowongan    uint   ` json:"id_lowongan"`
	NomorSoal     int    `json:"nomor_soal"`
	IDUjian       int    `json:"id_ujian"`
	Subtes        string `json:"subtes"`
	Jawaban       string `json:"jawaban"`
	Jawaban2      string `json:"jawaban2"`
	JawabanKunci  string `json:"jawaban_kunci"`
	JawabanKunci2 string `json:"jawaban_kunci2"`
}

func (CFITTestAnswer) TableName() string {
	return "tb_data_jawaban_cfit"
}
