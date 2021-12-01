package model

type Deskripsi struct {
	ID                  int `json:"id" form:"id"`
	Id_Akreditasi       int `json:"id_akreditasi" form:"id_akreditasi"`
	Id_Jurusans         int `json:"id_jurusans" form:"id_jurusans"`
	Id_Fakultas         int `json:"id_fakultas" form:"id_fakultas"`
	Id_FasilitasJurusan int `json:"id_fasilitasJurusan" form:"id_fasilitasJurusan"`
	Id_FasilitasKampus  int `json:"id_fasilitasKampus" form:"id_fasilitasKampus"`
	Id_FasilitasKelas   int `json:"id_fasilitasKelas" form:"id_fasilitasKelas"`
	Id_UKM              int `json:"id_ukm" form:"id_ukm"`
	Id_Universitas      int `json:"id_universitas" form:"id_universitas"`
}
