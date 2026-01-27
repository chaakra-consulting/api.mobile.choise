package models

type MSDTTestQuestion struct {
	IDSoal      uint   `gorm:"primaryKey" json:"id_soal"`
	NoSoal      int    `json:"no_soal"`
	Pernyataan1 string `json:"pernyataan_1"`
	Pernyataan2 string `json:"pernyataan_2"`
	Aspek1      string `json:"aspek_1"`
	Aspek2      string `json:"aspek_2"`
}

func (MSDTTestQuestion) TableName() string {
	return "tb_soal_msdt"
}
