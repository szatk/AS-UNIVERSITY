package controller

//import packages
import (
	"AS-UNIVERSITY/config"
	"AS-UNIVERSITY/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//func memanggil seluruh data UKM
func GetUKM(c echo.Context) error {
	var ukm []model.UKM

	if err := config.DB.Find(&ukm).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Semua Data UKM",
		"data":    ukm,
	})
}

//Fungsi get UKM by ID
func GetUKMByID(c echo.Context) error {
	var ukm model.UKM
	id, _ := strconv.Atoi(c.Param("id"))
	//config.DB.Where("id = ?", id).Delete(&admin)

	if err := config.DB.Where("id= ?", id).Find(&ukm).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan UKM ",
		"data":    ukm,
	})
}

//fungsi create new UKM
func CreateUKM(e echo.Context) error {
	ukm := model.UKM{}
	e.Bind(&ukm)

	if err := config.DB.Save(&ukm).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Menambahkan Data UKM",
		"data":    ukm,
	})
}

//Fungsi Update Tabel UKM
func UpdateUKM(e echo.Context) error {
	ukm := model.UKM{}
	id, _ := strconv.Atoi(e.Param("id"))
	e.Bind(&ukm)
	ukm.ID = id

	if err := config.DB.Where("id= ?", ukm.ID).Updates(&ukm).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Mengubah Data UKM",
		"data":    ukm,
	})
}

//Fungsi hapus data UKM
func DeleteUKM(e echo.Context) error {
	var ukm model.UKM
	id, _ := strconv.Atoi(e.Param("id"))
	config.DB.Where("id = ?", id).Delete(&ukm)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"data":    ukm,
		"message": "Data Berhasil Dihapus",
	})
}
