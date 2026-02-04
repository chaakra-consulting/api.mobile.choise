package models

type UserData struct {
	Nik                 string         `json:"nik"`
	IDPelamar           int64          `json:"id_pelamar"`
	NamaPelamar         string         `json:"nama_pelamar"`
	Alamat              string         `json:"alamat"`
	DomisiliJalan       string         `json:"domisili_jalan"`
	DomisiliProvinsiID  string         `json:"domisili_provinsi_id"`
	DomisiliKotaID      string         `json:"domisili_kota_id"`
	DomisiliKecamatanID string         `json:"domisili_kecamatan_id"`
	DomisiliKelurahanID string         `json:"domisili_kelurahan_id"`
	AlamatKtp           string         `json:"alamat_ktp"`
	KtpJalan            string         `json:"ktp_jalan"`
	KtpProvinsiID       string         `json:"ktp_provinsi_id"`
	KtpKotaID           string         `json:"ktp_kota_id"`
	KtpKecamatanID      string         `json:"ktp_kecamatan_id"`
	KtpKelurahanID      string         `json:"ktp_kelurahan_id"`
	StatusPerkawinan    string         `json:"status_perkawinan"`
	Agama               string         `json:"agama"`
	AnakKe              string         `json:"anak_ke"`
	DariXSDR            string         `json:"dari_x_sdr"`
	TempatLahir         string         `json:"tempat_lahir"`
	TanggalLahir        string         `json:"tanggal_lahir"`
	JenisKelamin        string         `json:"jenis_kelamin"`
	NoHP                string         `json:"no_hp"`
	Facebook            string         `json:"facebook"`
	Instagram           string         `json:"instagram"`
	Twitter             string         `json:"twitter"`
	Linkedin            string         `json:"linkedin"`
	ApplicantUser       *ApplicantUser `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:IDPelamar"`
}

func (UserData) TableName() string {
	return "tb_data_diri"
}
