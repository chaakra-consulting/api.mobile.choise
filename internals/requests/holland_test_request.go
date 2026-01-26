package requests

type HollandTest struct {
	IDPelamar  int `json:"id_pelamar" binding:"required"`
	IDLowongan int `json:"id_lowongan" binding:"required"`
	IDUjian    int `json:"id_ujian" binding:"required"`
	NilaiR     int `json:"nilai_r" binding:"required"`
	NilaiI     int `json:"nilai_i" binding:"required"`
	NilaiA     int `json:"nilai_a" binding:"required"`
	NilaiS     int `json:"nilai_s" binding:"required"`
	NilaiE     int `json:"nilai_e" binding:"required"`
	NilaiK     int `json:"nilai_k" binding:"required"`
}
