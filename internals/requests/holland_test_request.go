package requests

type HollandTest struct {
	IDPelamar  int `json:"id_pelamar" binding:"required"`
	IDLowongan int `json:"id_lowongan" binding:"required"`
	IDUjian    int `json:"id_ujian" binding:"required"`
	NilaiR     int `json:"nilai_r"`
	NilaiI     int `json:"nilai_i"`
	NilaiA     int `json:"nilai_a"`
	NilaiS     int `json:"nilai_s"`
	NilaiE     int `json:"nilai_e"`
	NilaiK     int `json:"nilai_k"`
}
