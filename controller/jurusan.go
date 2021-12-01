package controller

import (
	"AS-UNIVERSITY/config"
	"AS-UNIVERSITY/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//func memanggil seluruh data Jurusan
func GetJurusan(c echo.Context) error {
	var jurusan []model.Jurusans

	if err := config.DB.Find(&jurusan).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Semua Jurusan",
		"data":    jurusan,
	})
}

//Fungsi get Jurusan by ID
func GetJurusanByID(c echo.Context) error {
	var jurusan model.Jurusans
	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Where("id= ?", id).Find(&jurusan).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Jurusan",
		"data":    jurusan,
	})
}

//fungsi create Jurusan
func CreateJurusan(e echo.Context) error {
	jurusan := model.Jurusans{}
	e.Bind(&jurusan)

	if err := config.DB.Save(&jurusan).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menambahkan Jurusan",
		"data":    jurusan,
	})
}

//Fungsi Update Tabel Jurusan
func UpdateJurusan(e echo.Context) error {
	jurusan := model.Jurusans{}
	id, _ := strconv.Atoi(e.Param("id"))
	e.Bind(&jurusan)
	jurusan.ID = id

	if err := config.DB.Where("id= ?", jurusan.ID).Updates(&jurusan).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Mengubah Data Jurusan",
		"data":    jurusan,
	})
}

//Fungsi hapus data Jurusan
func DeleteJurusan(e echo.Context) error {
	var jurusan model.Jurusans
	id, _ := strconv.Atoi(e.Param("id"))
	config.DB.Where("id = ?", id).Delete(&jurusan)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"data":    jurusan,
		"message": "Data Berhasil Dihapus",
	})
}
