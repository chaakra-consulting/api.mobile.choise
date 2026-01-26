package requests

type HollandTest struct {
	IDPelamar  string `json:"id_pelamar" binding:"required"`
	IDLowongan string `json:"id_lowongan" binding:"required"`
	IDUjian    string `json:"id_ujian" binding:"required"`
	NilaiR     string `json:"nilai_r" binding:"required"`
	NilaiI     string `json:"nilai_i" binding:"required"`
	NilaiA     string `json:"nilai_a" binding:"required"`
	NilaiS     string `json:"nilai_s" binding:"required"`
	NilaiE     string `json:"nilai_e" binding:"required"`
	NilaiK     string `json:"nilai_k" binding:"required"`
}
