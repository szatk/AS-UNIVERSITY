package model

type Universitas struct {
	ID              int    `json:"id" form:"id"`
	NamaUniversitas string `json:"namaUniversitas" form:"namaUniversitas"`
	Luas            string `json:"luas" form:"luas"`
	Id_Akreditasi   int    `json:"id_akreditasi" form:"id_akreditasi"`
	Akreditasi      string `json:"akreditasi"`
}
