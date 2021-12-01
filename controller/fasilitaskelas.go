package controller

//import packages
import (
	"AS-UNIVERSITY/config"
	"AS-UNIVERSITY/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//func memanggil seluruh data Fasilitas Kelas
func GetFasilitasKelas(c echo.Context) error {
	var fasilitaskelas []model.FasilitasKelas

	if err := config.DB.Find(&fasilitaskelas).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Semua Data Fasilitas Kelas",
		"data":    fasilitaskelas,
	})
}

//Fungsi get FasilitasKelas by ID
func GetFasilitasKelasByID(c echo.Context) error {
	var fasilitaskelas model.FasilitasKelas
	id, _ := strconv.Atoi(c.Param("id"))
	//config.DB.Where("id = ?", id).Delete(&admin)

	if err := config.DB.Where("id= ?", id).Find(&fasilitaskelas).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Fasilitas Kelas ",
		"data":    fasilitaskelas,
	})
}

//fungsi create Fasilitas Kelas
func CreateFasilitasKelas(e echo.Context) error {
	fasilitaskelas := model.FasilitasKelas{}
	e.Bind(&fasilitaskelas)

	if err := config.DB.Save(&fasilitaskelas).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Menambahkan Data Fasilitas Kelas",
		"data":    fasilitaskelas,
	})
}

//Fungsi Update Tabel Fasilitas Kelas
func UpdateFasilitasKelas(e echo.Context) error {
	fasilitaskelas := model.FasilitasKampus{}
	id, _ := strconv.Atoi(e.Param("id"))
	e.Bind(&fasilitaskelas)
	fasilitaskelas.ID = id

	if err := config.DB.Where("id= ?", fasilitaskelas.ID).Updates(&fasilitaskelas).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Mengubah Data Fasilitas Kelas",
		"data":    fasilitaskelas,
	})
}

//Fungsi hapus data Fasilitas Kelas
func DeleteFasilitasKelas(e echo.Context) error {
	var fasilitaskelas model.FasilitasKelas
	id, _ := strconv.Atoi(e.Param("id"))
	config.DB.Where("id = ?", id).Delete(&fasilitaskelas)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"data":    fasilitaskelas,
		"message": "Data Berhasil Dihapus",
	})
}
