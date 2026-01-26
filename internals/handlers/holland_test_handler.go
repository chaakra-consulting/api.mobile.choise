package handlers

import (
	"api.choise/internals/helpers"
	"api.choise/internals/models"
	"api.choise/internals/requests"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// HollandTestQuestion returns a list of HollandTest questions from the database.
// The handler will return a 500 error response if an error occurs while retrieving the questions.
// Otherwise, it will return a 200 response with the list of questions in JSON format.
func HollandTestQuestion(c *gin.Context) {
	var hollandQuestions []models.HollandTest
	// Realistic,Investigative,Artistic,Social,Enterprising,Conventional
	hollandQuestions = []models.HollandTest{
		//Realistic - I as
		{ID: 1, QuestionText: "Praktis", Type: "Realistic", Identifier: "i_as"},
		{ID: 2, QuestionText: "Atletis", Type: "Realistic", Identifier: "i_as"},
		{ID: 3, QuestionText: "Berterus Terang", Type: "Realistic", Identifier: "i_as"},
		{ID: 4, QuestionText: "Cenderung Mekanistik", Type: "Realistic", Identifier: "i_as"},
		{ID: 5, QuestionText: "Pecinta Alam", Type: "Realistic", Identifier: "i_as"},
		{ID: 6, QuestionText: "Dapat mengoprasikan mesin dan peralatan", Type: "Realistic", Identifier: "i_as"},

		//Realistic - I can
		{ID: 7, QuestionText: "Memperbaiki peralatan elektronik", Type: "Realistic", Identifier: "i_can"},
		{ID: 8, QuestionText: "Memecahkan masalah mekanistik", Type: "Realistic", Identifier: "i_can"},
		{ID: 9, QuestionText: "Memasang/merakit barang elektronik", Type: "Realistic", Identifier: "i_can"},
		{ID: 10, QuestionText: "Melakukan kegiatan olahraga", Type: "Realistic", Identifier: "i_can"},
		{ID: 11, QuestionText: "Membaca petunjuk arah", Type: "Realistic", Identifier: "i_can"},
		{ID: 12, QuestionText: "Memperbaiki kendaraan bermotor", Type: "Realistic", Identifier: "i_can"},

		//Realistic - I like
		{ID: 13, QuestionText: "Kegiatan yang menuntut untuk berpikir secara kritis", Type: "Realistic", Identifier: "i_like"},
		{ID: 14, QuestionText: "Pekerjaan Lapangan", Type: "Realistic", Identifier: "i_like"},
		{ID: 15, QuestionText: "Kegiatan yang aktif secara fisik", Type: "Realistic", Identifier: "i_like"},
		{ID: 16, QuestionText: "Keterampilan tangan", Type: "Realistic", Identifier: "i_like"},
		{ID: 17, QuestionText: "Kegiatan merakit atau menyusun sesuatu", Type: "Realistic", Identifier: "i_like"},
		{ID: 18, QuestionText: "-", Type: "Realistic", Identifier: "i_like"},

		//Investigative - I as
		{ID: 19, QuestionText: "Memiliki rasa ingin tahu", Type: "Investigative", Identifier: "i_as"},
		{ID: 20, QuestionText: "Analitis", Type: "Investigative", Identifier: "i_as"},
		{ID: 21, QuestionText: "Berpikir Ilmiah", Type: "Investigative", Identifier: "i_as"},
		{ID: 22, QuestionText: "Pengamat", Type: "Investigative", Identifier: "i_as"},
		{ID: 23, QuestionText: "Memiliki ketepatan", Type: "Investigative", Identifier: "i_as"},
		{ID: 24, QuestionText: "Dapat mengoprasikan mesin dan peralatan", Type: "Investigative", Identifier: "i_as"},
		{ID: 25, QuestionText: "-", Type: "Investigative", Identifier: "i_as"},

		//Investigative - I can
		{ID: 26, QuestionText: "Berpikir Abstrak", Type: "Investigative", Identifier: "i_can"},
		{ID: 27, QuestionText: "Memecahkan masalah hitungan", Type: "Investigative", Identifier: "i_can"},
		{ID: 28, QuestionText: "memahami teori-teori fisika", Type: "Investigative", Identifier: "i_can"},
		{ID: 29, QuestionText: "Melakukan hitungan kompleks", Type: "Investigative", Identifier: "i_can"},
		{ID: 30, QuestionText: "Menggunakan alat mikroskop", Type: "Investigative", Identifier: "i_can"},
		{ID: 31, QuestionText: "Memperbaiki mobil/kendaraan bermotor", Type: "Investigative", Identifier: "i_can"},
		{ID: 32, QuestionText: "Menganalisis data", Type: "Investigative", Identifier: "i_can"},

		//Investigative - I like
		{ID: 33, QuestionText: "Mengeksplorasikan ide-ide", Type: "Investigative", Identifier: "i_like"},
		{ID: 34, QuestionText: "Menggunakan komputer", Type: "Investigative", Identifier: "i_like"},
		{ID: 35, QuestionText: "Bekerja secara mandiri", Type: "Investigative", Identifier: "i_like"},
		{ID: 36, QuestionText: "Melakukan kegiatan di Laboratorium", Type: "Investigative", Identifier: "i_like"},
		{ID: 37, QuestionText: "Membaca majalah ilmiah atau teknik", Type: "Investigative", Identifier: "i_like"},
		{ID: 38, QuestionText: "-", Type: "Investigative", Identifier: "i_like"},
		{ID: 39, QuestionText: "-", Type: "Investigative", Identifier: "i_like"},

		//Artistic - I as
		{ID: 40, QuestionText: "Kreatif", Type: "Artistic", Identifier: "i_as"},
		{ID: 41, QuestionText: "Intuitif", Type: "Artistic", Identifier: "i_as"},
		{ID: 42, QuestionText: "Imaginative", Type: "Artistic", Identifier: "i_as"},
		{ID: 43, QuestionText: "Inovatif", Type: "Artistic", Identifier: "i_as"},
		{ID: 44, QuestionText: "Individualistik", Type: "Artistic", Identifier: "i_as"},

		//Artistic - I can
		{ID: 45, QuestionText: "Membuat sketsa, menggambar dan melukis", Type: "Artistic", Identifier: "i_can"},
		{ID: 46, QuestionText: "Menmainkan alat musik", Type: "Artistic", Identifier: "i_can"},
		{ID: 47, QuestionText: "Menulis cerita,puisi, maupun lirik lagu", Type: "Artistic", Identifier: "i_can"},
		{ID: 48, QuestionText: "Merancang busana atau interior ruangan", Type: "Artistic", Identifier: "i_can"},
		{ID: 49, QuestionText: "-", Type: "Artistic", Identifier: "i_can"},

		//Artistic - I like
		{ID: 50, QuestionText: "Menghadiri konser, pertunjukan teater dan seni", Type: "Artistic", Identifier: "i_like"},
		{ID: 51, QuestionText: "Membaca bacaan fiksi dan puisi", Type: "Artistic", Identifier: "i_like"},
		{ID: 52, QuestionText: "Membuat kerajinan tangan", Type: "Artistic", Identifier: "i_like"},
		{ID: 53, QuestionText: "Kegiatan fotografi", Type: "Artistic", Identifier: "i_like"},
		{ID: 54, QuestionText: "Mengekspresian diri secara kreatif", Type: "Artistic", Identifier: "i_like"},

		//Social - I as
		{ID: 55, QuestionText: "Bersahabat", Type: "Social", Identifier: "i_as"},
		{ID: 56, QuestionText: "Penolong", Type: "Social", Identifier: "i_as"},
		{ID: 57, QuestionText: "Idealis", Type: "Social", Identifier: "i_as"},
		{ID: 58, QuestionText: "Berwawasan", Type: "Social", Identifier: "i_as"},
		{ID: 59, QuestionText: "Ramah", Type: "Social", Identifier: "i_as"},
		{ID: 60, QuestionText: "Pengertian", Type: "Social", Identifier: "i_as"},

		//Social - I can
		{ID: 61, QuestionText: "Mengajar/membimbing orang lain", Type: "Social", Identifier: "i_can"},
		{ID: 62, QuestionText: "Mengekspresikan diri", Type: "Social", Identifier: "i_can"},
		{ID: 63, QuestionText: "Memimpin diskusi kelompok", Type: "Social", Identifier: "i_can"},
		{ID: 63, QuestionText: "Menengahi konflik", Type: "Social", Identifier: "i_can"},
		{ID: 64, QuestionText: "Merencanakan dan mengawasi aktifitas", Type: "Social", Identifier: "i_can"},
		{ID: 65, QuestionText: "Bekerjasama dengan orang lain", Type: "Social", Identifier: "i_can"},

		//Social - I like
		{ID: 66, QuestionText: "Bekerja dalam kelompok", Type: "Social", Identifier: "i_like"},
		{ID: 67, QuestionText: "Membantu permasalahan orang lain", Type: "Social", Identifier: "i_like"},
		{ID: 68, QuestionText: "Berpartisipasi dalam berbagai pertemuan", Type: "Social", Identifier: "i_like"},
		{ID: 69, QuestionText: "Berperan sebagai relawan yang membantu orang lain", Type: "Social", Identifier: "i_like"},
		{ID: 70, QuestionText: "Bekerja dengan anak muda", Type: "Social", Identifier: "i_like"},
		{ID: 71, QuestionText: "Berolahraga secara bersama-sama", Type: "Social", Identifier: "i_like"},

		//Konvensional - I as
		{ID: 72, QuestionText: "Berpenampilan rapi", Type: "Konvensional", Identifier: "i_as"},
		{ID: 73, QuestionText: "Tepat", Type: "Konvensional", Identifier: "i_as"},
		{ID: 74, QuestionText: "Cenderung numerik", Type: "Konvensional", Identifier: "i_as"},
		{ID: 75, QuestionText: "Metodis/Prosedural", Type: "Konvensional", Identifier: "i_as"},
		{ID: 76, QuestionText: "Teliti", Type: "Konvensional", Identifier: "i_as"},
		{ID: 77, QuestionText: "Efisien", Type: "Konvensional", Identifier: "i_as"},

		//Konvensional - I can
		{ID: 78, QuestionText: "Bekerja baik dalam sistem", Type: "Konvensional", Identifier: "i_can"},
		{ID: 79, QuestionText: "Melakukan berbagai pekerjaan dalam waktu yang singkat", Type: "Konvensional", Identifier: "i_can"},
		{ID: 80, QuestionText: "Menjaga rahasia secara akurat", Type: "Konvensional", Identifier: "i_can"},
		{ID: 81, QuestionText: "Menggunakan komputer", Type: "Konvensional", Identifier: "i_can"},
		{ID: 82, QuestionText: "Menulis surat-surat bisnis dengan bahasa yang efektif", Type: "Konvensional", Identifier: "i_can"},
		{ID: 83, QuestionText: "-", Type: "Konvensional", Identifier: "i_can"},

		//Konvensional - I like
		{ID: 84, QuestionText: "Mengikuti prosedur yang diberikan dengan benar", Type: "Konvensional", Identifier: "i_like"},
		{ID: 85, QuestionText: "Menggunakan peralatan pengolahan data", Type: "Konvensional", Identifier: "i_like"},
		{ID: 86, QuestionText: "Bekerja dengan angka", Type: "Konvensional", Identifier: "i_like"},
		{ID: 87, QuestionText: "Mengetik dengan cepat", Type: "Konvensional", Identifier: "i_like"},
		{ID: 88, QuestionText: "Bekerja secara teliti dan detail", Type: "Konvensional", Identifier: "i_like"},
		{ID: 89, QuestionText: "-", Type: "Konvensional", Identifier: "i_like"},

		//Enterprising - I as
		{ID: 90, QuestionText: "Percaya Diri", Type: "Enterprising", Identifier: "i_as"},
		{ID: 91, QuestionText: "Asertif", Type: "Enterprising", Identifier: "i_as"},
		{ID: 92, QuestionText: "Mudah menjalin hubungan sosial", Type: "Enterprising", Identifier: "i_as"},
		{ID: 93, QuestionText: "Persuasif", Type: "Enterprising", Identifier: "i_as"},
		{ID: 94, QuestionText: "Antusias", Type: "Enterprising", Identifier: "i_as"},
		{ID: 95, QuestionText: "Enerjik", Type: "Enterprising", Identifier: "i_as"},

		//Enterprising - I can
		{ID: 96, QuestionText: "Memulai proyek dan tantangan baru", Type: "Enterprising", Identifier: "i_can"},
		{ID: 97, QuestionText: "Meyakinkan orang lain untuk melakukan sesuatu sesuai dengan yang diinginkan", Type: "Enterprising", Identifier: "i_can"},
		{ID: 98, QuestionText: "Menjual barang dan mepromosikan berbagai ide", Type: "Enterprising", Identifier: "i_can"},
		{ID: 99, QuestionText: "Menjadi pembicara", Type: "Enterprising", Identifier: "i_can"},
		{ID: 100, QuestionText: "Mengelola suatu aktifitas dan kegiatan", Type: "Enterprising", Identifier: "i_can"},
		{ID: 101, QuestionText: "Memimpin suatu kelompok", Type: "Enterprising", Identifier: "i_can"},

		//Enterprising - I like
		{ID: 102, QuestionText: "Membuat keputusan yang mempengaruhi orang lain", Type: "Enterprising", Identifier: "i_like"},
		{ID: 103, QuestionText: "Terpilih untuk mewakili suatu kegiatan", Type: "Enterprising", Identifier: "i_like"},
		{ID: 104, QuestionText: "Memenangkan suatu penghargaan kepemimpinan atau penjualan", Type: "Enterprising", Identifier: "i_like"},
		{ID: 105, QuestionText: "Memulai kampanye politik", Type: "Enterprising", Identifier: "i_like"},
		{ID: 106, QuestionText: "Bertem dengan orang penting", Type: "Enterprising", Identifier: "i_like"},
		{ID: 107, QuestionText: "-", Type: "Enterprising", Identifier: "i_like"},
	}

	helpers.SuccessResponse(c, 200, hollandQuestions)
}

func HollandPostAnswer(c *gin.Context) {
	var hollandTestRequest requests.HollandTest
	if err := c.ShouldBindBodyWith(&hollandTestRequest, binding.JSON); err != nil {
		helpers.ErrorResponse(c, 400, c.Error(err))
		return
	}

}
