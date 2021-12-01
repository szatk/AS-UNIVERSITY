package model

type Fakultas struct {
	ID            int    `json:"id" form:"id"`
	Nama_Fakultas string `json:"nama_fakultas" form:"nama_fakultas"`
	Id_Jurusans   string `json:"id_jurusans" form:"id_jurusans"`
	Jurusans      string `json:"jurusans"`
}
