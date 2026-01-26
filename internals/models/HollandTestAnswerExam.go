package models

type HollandTestAnswerExam struct {
	IdJawabanHolland int `gorm:"column:id_jawaban_holland;primaryKey;autoIncrement:true;"`
	IDPelamar        int `gorm:"column:id_pelamar"`
	IDLowongan       int `gorm:"column:id_lowongan"`
	IDUjian          int `gorm:"column:id_ujian"`
	NilaiR           int `gorm:"column:nilai_r"`
	NilaiI           int `gorm:"column:nilai_i"`
	NilaiA           int `gorm:"column:nilai_a"`
	NilaiS           int `gorm:"column:nilai_s"`
	NilaiE           int `gorm:"column:nilai_e"`
	NilaiK           int `gorm:"column:nilai_k"`
}

func (HollandTestAnswerExam) TableName() string {
	return "tb_data_jawaban_holland"
}
