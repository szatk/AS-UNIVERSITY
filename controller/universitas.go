package controller

//import packages
import (
	"AS-UNIVERSITY/config"
	"AS-UNIVERSITY/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//func memanggil seluruh data Universitas
func GetUniv(c echo.Context) error {
	var universitas []model.Universitas

	if err := config.DB.Find(&universitas).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	for i, value := range universitas {
		var akreditas model.Akreditasi
		if err := config.DB.Where("id= ?", value.Id_Akreditasi).First(&akreditas).Error; err == nil {
			universitas[i].Akreditasi = akreditas.Akreditas
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "Berhasil Menampilkan Semua Universitas",
		"universitas": universitas,
	})
}

//Fungsi get Universitas by ID
func GetUnivByID(c echo.Context) error {
	var universitas model.Universitas
	id, _ := strconv.Atoi(c.Param("id"))
	//config.DB.Where("id = ?", id).Delete(&admin)

	if err := config.DB.Where("id= ?", id).Find(&universitas).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "Berhasil Menampilkan Universitas ",
		"universitas": universitas,
	})
}

//fungsi create new Universitas
func CreateUniv(e echo.Context) error {
	universitas := model.Universitas{}
	e.Bind(&universitas)

	if err := config.DB.Save(&universitas).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message":     "Berhasil Menambahkan Universitas",
		"universitas": universitas,
	})
}

//Fungsi Update Tabel Universitas
func UpdateUniv(e echo.Context) error {
	universitas := model.Universitas{}
	id, _ := strconv.Atoi(e.Param("id"))
	e.Bind(&universitas)
	universitas.ID = id

	if err := config.DB.Where("id= ?", universitas.ID).Updates(&universitas).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message":     "Berhasil Mengubah Data Universitas",
		"universitas": universitas,
	})
}

//Fungsi hapus data Univ
func DeleteUniv(e echo.Context) error {
	var universitas model.Universitas
	id, _ := strconv.Atoi(e.Param("id"))
	config.DB.Where("id = ?", id).Delete(&universitas)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message":     "Data Berhasil Dihapus",
		"universitas": universitas,
	})
}
