package model

type FasilitasKampus struct {
	ID           int    `json:"id" form:"id"`
	RumahSakit   string `json:"rumahSakit" form:"rumahSakit"`
	Taman        string `json:"taman" form:"taman"`
	Kantin       string `json:"kantin" form:"kantin"`
	Lapangan     string `json:"lapangan" form:"lapangan"`
	ATM          string `json:"atm" form:"atm"`
	GOR          string `json:"gor" form:"gor"`
	Perpustakaan string `json:"perpustakaan" form:"perpustakaan"`
	Aula         string `json:"aula" form:"aula"`
}
