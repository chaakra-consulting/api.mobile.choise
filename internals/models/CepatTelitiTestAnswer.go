package models

type CepatTelitiTestAnswer struct {
	IDJawabanCepat uint   `gorm:"primaryKey" json:"id_jawaban_cepat"`
	IDPelamar      uint   `json:"id_pelamar" binding:"required"`
	IDLowongan     uint   `json:"id_lowongan" binding:"required"`
	IDUjian        uint   `json:"id_ujian" binding:"required"`
	NoSoal         uint   `json:"no_soal" binding:"required"`
	Jawaban        string `json:"jawaban" binding:"required"`
}

func (CepatTelitiTestAnswer) TableName() string {
	return "tb_data_jawaban_cepat"
}
