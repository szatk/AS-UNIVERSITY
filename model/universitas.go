package model

type Universitas struct {
	ID              uint    `json:"id" form:"id"`
	NamaUniversitas string `json:"namaUniversitas" form:"namaUniversitas"`
	Luas            string `json:"luas" form:"luas"`
	Id_Akreditasi   uint    `json:"id_akreditasi" form:"id_akreditasi"`
	Akreditasi      Akreditasi
}
