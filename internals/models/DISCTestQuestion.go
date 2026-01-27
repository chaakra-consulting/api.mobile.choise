package models

type DISCTestQuestion struct {
	IDSoal      uint   `gorm:"primaryKey" json:"id_soal"`
	NoSoal      int    `json:"no_soal"`
	Pernyataan1 string `json:"pernyataan_1"`
	Pernyataan2 string `json:"pernyataan_2"`
	Pernyataan3 string `json:"pernyataan_3"`
	Pernyataan4 string `json:"pernyataan_4"`
	AspekM1     string `json:"aspek_m1"`
	AspekM2     string `json:"aspek_m2"`
	AspekM3     string `json:"aspek_m3"`
	AspekM4     string `json:"aspek_m4"`
	AspekL1     string `json:"aspek_l1"`
	AspekL2     string `json:"aspek_l2"`
	AspekL3     string `json:"aspek_l3"`
	AspekL4     string `json:"aspek_l4"`
}

func (DISCTestQuestion) TableName() string {
	return "tb_soal_disc"
}
