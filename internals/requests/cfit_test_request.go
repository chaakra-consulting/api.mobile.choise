package requests

type CFITTestRequest struct {
	IDPelamar     int    `json:"id_pelamar" binding:"required"`
	IDLowongan    int    `json:"id_lowongan" binding:"required"`
	IDUjian       int    `json:"id_ujian" binding:"required"`
	NomorSoal     int    `json:"nomor_soal" binding:"required"`
	Jawaban       string `json:"jawaban" binding:"required"`
	Jawaban2      string `json:"jawaban2"`
	JawabanKunci  string `json:"jawaban_kunci" binding:"required"`
	JawabanKunci2 string `json:"jawaban_kunci2"`
}
