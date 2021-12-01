package controller

import (
	"AS-UNIVERSITY/config"
	"AS-UNIVERSITY/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//func memanggil seluruh data Akreditas
func GetAkreditas(c echo.Context) error {
	var akreditas []model.Akreditasi

	if err := config.DB.Find(&akreditas).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "Berhasil Menampilkan Semua Akreditas",
		"akreditas": akreditas,
	})
}

//Fungsi get Akreditas by ID
func GetAkreditasByID(c echo.Context) error {
	var akreditas model.Akreditasi
	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Where("id= ?", id).Find(&akreditas).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "Berhasil Menampilkan Akreditas ",
		"akreditas": akreditas,
	})
}

//fungsi create Akreditas
func CreateAkreditas(e echo.Context) error {
	akreditas := model.Akreditasi{}
	e.Bind(&akreditas)

	if err := config.DB.Save(&akreditas).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message":   "success menambahkan Akreditas",
		"akreditas": akreditas,
	})
}

//Fungsi Update Tabel Akreditas
func UpdateAkreditas(e echo.Context) error {
	akreditas := model.Akreditasi{}
	id, _ := strconv.Atoi(e.Param("id"))
	e.Bind(&akreditas)
	akreditas.ID = id

	if err := config.DB.Where("id= ?", akreditas.ID).Updates(&akreditas).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message":   "Berhasil Mengubah Data Akreditas",
		"akreditas": akreditas,
	})
}

//Fungsi hapus data Akreditas
func DeleteAkreditas(e echo.Context) error {
	var akreditas model.Akreditasi
	id, _ := strconv.Atoi(e.Param("id"))
	config.DB.Where("id = ?", id).Delete(&akreditas)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"akreditas": akreditas,
		"message":   "Data Berhasil Dihapus",
	})
}
