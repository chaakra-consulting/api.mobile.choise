package models

type DISCTestAnswer struct {
	IDJawabanDISC int    `gorm:"primaryKey;autoIncrement" json:"id_jawaban_disc"`
	IDPelamar     int    `json:"id_pelamar" binding:"required"`
	IDLowongan    int    `json:"id_lowongan" binding:"required"`
	IDUjian       int    `json:"id_ujian" binding:"required"`
	NoSoal        int    `json:"no_soal" binding:"required"`
	Jawaban       string `json:"jawaban" binding:"required"`
	Jawaban2      string `json:"jawaban2" binding:"required"`
}

func (DISCTestAnswer) TableName() string {
	return "tb_data_jawaban_disc"
}
