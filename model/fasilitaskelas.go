package model

type FasilitasKelas struct {
	ID              int    `json:"id" form:"id"`
	AC              string `json:"ac" form:"ac"`
	Kursi           string `json:"kursi" form:"kursi"`
	Kapasitas_Kelas int    `json:"kapasitas_kelas" form:"kapasitas_kelas"`
	Papan_Tulis     string `json:"papan_tulis" form:"papan_tulis"`
	Proyektor       string `json:"proyektor" form:"proyektor"`
}
