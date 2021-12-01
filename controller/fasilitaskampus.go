package controller

//import packages
import (
	"AS-UNIVERSITY/config"
	"AS-UNIVERSITY/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//func memanggil seluruh Fasilitas Kampus
func GetFasilitasKampus(c echo.Context) error {
	var fasilitaskampus []model.FasilitasKampus

	if err := config.DB.Find(&fasilitaskampus).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Semua Data Fasilitas Kampus",
		"data":    fasilitaskampus,
	})
}

//Fungsi get FasilitasKampus by ID
func GetFasilitasKampusByID(c echo.Context) error {
	var fasilitaskampus model.FasilitasKampus
	id, _ := strconv.Atoi(c.Param("id"))
	//config.DB.Where("id = ?", id).Delete(&admin)

	if err := config.DB.Where("id= ?", id).Find(&fasilitaskampus).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Fasilitas Kampus ",
		"data":    fasilitaskampus,
	})
}

//fungsi create Fasilitas Kampus
func CreateFasilitasKampus(e echo.Context) error {
	fasilitaskampus := model.FasilitasKampus{}
	e.Bind(&fasilitaskampus)

	if err := config.DB.Save(&fasilitaskampus).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Menambahkan Data Fasilitas Kampus",
		"data":    fasilitaskampus,
	})
}

//Fungsi Update Tabel Fasilitas Kampus
func UpdateFasilitasKampus(e echo.Context) error {
	fasilitaskampus := model.FasilitasKampus{}
	id, _ := strconv.Atoi(e.Param("id"))
	e.Bind(&fasilitaskampus)
	fasilitaskampus.ID = id

	if err := config.DB.Where("id= ?", fasilitaskampus.ID).Updates(&fasilitaskampus).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Mengubah Data Fasilitas Kampus",
		"data":    fasilitaskampus,
	})
}

//Fungsi hapus data Fasilitas Kampus
func DeleteFasilitasKampus(e echo.Context) error {
	var fasilitaskampus model.FasilitasKampus
	id, _ := strconv.Atoi(e.Param("id"))
	config.DB.Where("id = ?", id).Delete(&fasilitaskampus)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"data":    fasilitaskampus,
		"message": "Data Berhasil Dihapus",
	})
}
