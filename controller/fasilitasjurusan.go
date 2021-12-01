package controller

//import packages
import (
	"AS-UNIVERSITY/config"
	"AS-UNIVERSITY/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//func memanggil seluruh data Fasilitas Jurusan
func GetFasilitasJurusan(c echo.Context) error {
	var fasilitasjurusan []model.FasilitasJurusan

	if err := config.DB.Find(&fasilitasjurusan).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Semua Data Fasilitas Jurusan",
		"data":    fasilitasjurusan,
	})
}

//Fungsi get Fasilitas Jurusan by ID
func GetFasilitasJurusanByID(c echo.Context) error {
	var fasilitasjurusan model.FasilitasJurusan
	id, _ := strconv.Atoi(c.Param("id"))
	//config.DB.Where("id = ?", id).Delete(&admin)

	if err := config.DB.Where("id= ?", id).Find(&fasilitasjurusan).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Fasilitas Jurusan ",
		"data":    fasilitasjurusan,
	})
}

//fungsi create Fasilitas Jurusan
func CreateFasilitasJurusan(e echo.Context) error {
	fasilitasjurusan := model.FasilitasJurusan{}
	e.Bind(&fasilitasjurusan)

	if err := config.DB.Save(&fasilitasjurusan).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Menambahkan Data Fasilitas Jurusan",
		"data":    fasilitasjurusan,
	})
}

//Fungsi Update Tabel Fasilitas Jurusan
func UpdateFasilitasJurusan(e echo.Context) error {
	fasilitasjurusan := model.FasilitasJurusan{}
	id, _ := strconv.Atoi(e.Param("id"))
	e.Bind(&fasilitasjurusan)
	fasilitasjurusan.ID = id

	if err := config.DB.Where("id= ?", fasilitasjurusan.ID).Updates(&fasilitasjurusan).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Mengubah Data Fasilitas Jurusan",
		"data":    fasilitasjurusan,
	})
}

//Fungsi hapus data Fasilitas Jurusan
func DeleteFasilitasJurusan(e echo.Context) error {
	var fasilitasjurusan model.FasilitasJurusan
	id, _ := strconv.Atoi(e.Param("id"))
	config.DB.Where("id = ?", id).Delete(&fasilitasjurusan)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"data":    fasilitasjurusan,
		"message": "Data Berhasil Dihapus",
	})
}
