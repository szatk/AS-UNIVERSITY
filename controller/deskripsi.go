package controller

//import packages
import (
	"AS-UNIVERSITY/config"
	"AS-UNIVERSITY/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//func memanggil seluruh data Deskripsi
func GetDeskripsi(c echo.Context) error {
	var deskripsi []model.Deskripsi

	if err := config.DB.Find(&deskripsi).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Semua Data Deskripsi",
		"data":    deskripsi,
	})
}

//Fungsi get deskripsi by ID
func GetDeskripsiByID(c echo.Context) error {
	var deskripsi model.Deskripsi
	id, _ := strconv.Atoi(c.Param("id"))
	//config.DB.Where("id = ?", id).Delete(&admin)

	if err := config.DB.Where("id= ?", id).Find(&deskripsi).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Deskripsi",
		"data":    deskripsi,
	})
}

//fungsi create new deskripsi
func CreateDeskripsi(e echo.Context) error {
	deskripsi := model.Deskripsi{}
	e.Bind(&deskripsi)

	if err := config.DB.Save(&deskripsi).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success menambahkan data Deskripsi",
		"data":    deskripsi,
	})
}

//Fungsi Update Tabel Deskripsi
func UpdateDeskripsiByID(e echo.Context) error {
	deskripsi := model.Deskripsi{}
	id, _ := strconv.Atoi(e.Param("id"))
	e.Bind(&deskripsi)
	deskripsi.ID = id

	if err := config.DB.Where("id= ?", deskripsi.ID).Updates(&deskripsi).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Mengubah Data Deskripsi",
		"data":    deskripsi,
	})
}

//Fungsi hapus data deskripsi
func DeleteDeskripsiByID(e echo.Context) error {
	var deskripsi model.Deskripsi
	id, _ := strconv.Atoi(e.Param("id"))
	config.DB.Where("id = ?", id).Delete(&deskripsi)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"descs":   deskripsi,
		"message": "Data Berhasil Dihapus",
	})
}
