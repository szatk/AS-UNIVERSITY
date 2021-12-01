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
func GetFakultas(c echo.Context) error {
	var fakultas []model.Fakultas

	if err := config.DB.Find(&fakultas).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	for i, value := range fakultas {
		var jurusan model.Jurusans
		if err := config.DB.Where("id= ?", value.Id_Jurusans).First(&jurusan).Error; err == nil {
			fakultas[i].Jurusans = jurusan.Jurusan
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Semua Fakultas",
		"data":    fakultas,
	})
}

//Fungsi get Fakultas by ID
func GetFakultasByID(c echo.Context) error {
	var fakultas model.Fakultas
	id, _ := strconv.Atoi(c.Param("id"))
	//config.DB.Where("id = ?", id).Delete(&admin)

	if err := config.DB.Where("id= ?", id).Find(&fakultas).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Fakultas",
		"data":    fakultas,
	})
}

//fungsi create new Fakultas
func CreateFakultas(e echo.Context) error {
	fakultas := model.Fakultas{}
	e.Bind(&fakultas)

	if err := config.DB.Save(&fakultas).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menambahkan Fakultas",
		"data":    fakultas,
	})
}

//Fungsi Update Tabel Fakultas
func UpdateFakultas(e echo.Context) error {
	fakultas := model.Fakultas{}
	id, _ := strconv.Atoi(e.Param("id"))
	e.Bind(&fakultas)
	fakultas.ID = id

	if err := config.DB.Where("id= ?", fakultas.ID).Updates(&fakultas).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Mengubah Data Fakultas",
		"data":    fakultas,
	})
}

//Fungsi hapus data Fakultas
func DeleteFakultas(e echo.Context) error {
	var fakultas model.Fakultas
	id, _ := strconv.Atoi(e.Param("id"))
	config.DB.Where("id = ?", id).Delete(&fakultas)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data Berhasil Dihapus",
		"data":    fakultas,
	})
}
