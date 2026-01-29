package requests

type MSDTTestRequest struct {
	IDPelamar  int    `json:"id_pelamar" binding:"required"`
	IDLowongan int    `json:"id_lowongan" binding:"required"`
	IDUjian    int    `json:"id_ujian" binding:"required"`
	NoSoal     int    `json:"no_soal" binding:"required"`
	Jawaban    string `json:"jawaban" binding:"required"`
}
