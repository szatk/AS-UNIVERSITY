package model

type FasilitasJurusan struct {
	ID            int    `json:"id" form:"id"`
	AC            string `json:"ac" form:"ac"`
	Ruang_Kelas   string `json:"ruang_kelas" form:"ruang_kelas"`
	Administrasi  string `json:"administrasi" form:"administrasi"`
	Jumlah_Kelas  int    `json:"jumlah_kelas" form:"jumlah_kelas"`
	Laboratorium  string `json:"laboratorium" form:"laboratorium"`
	Kamar_Mandi   string `json:"kamar_mandi" form:"kamar_mandi"`
	Perpustakaan  string `json:"perpustakaan" form:"perpustakaan"`
	Tempat_Parkir string `json:"tempat_parkir" form:"tempat_parkir"`
	Wifi          string `json:"wifi" form:"wifi"`
}
