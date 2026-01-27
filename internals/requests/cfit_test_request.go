package requests

type CFITTestRequest struct {
	IDPelamar     int    `json:"id_pelamar" binding:"required"`
	IDLowongan    int    `json:"id_lowongan" binding:"required"`
	IDUjian       int    `json:"id_ujian" binding:"required"`
	Subtes        string `json:"subtes" binding:"required"`
	NomorSoal     int    `json:"nomor_soal" binding:"required"`
	Jawaban       string `json:"jawaban" binding:"required"`
	KunciJawaban  string `json:"kunci_jawaban" binding:"required"`
	KunciJawaban2 string `json:"kunci_jawaban2"`
}
