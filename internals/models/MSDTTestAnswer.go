package models

type MSDTTestAnswer struct {
	IDJawabanMSDT int    `gorm:"primaryKey;autoIncrement" json:"id_jawaban_msdt"`
	IDPelamar     int    `json:"id_pelamar" binding:"required"`
	IDLowongan    int    `json:"id_lowongan" binding:"required"`
	IDUjian       int    `json:"id_ujian" binding:"required"`
	NoSoal        int    `json:"no_soal" binding:"required"`
	Jawaban       string `json:"jawaban" binding:"required"`
}

func (MSDTTestAnswer) TableName() string {
	return "tb_data_jawaban_msdt"
}
