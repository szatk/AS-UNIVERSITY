package route

import (
	"AS-UNIVERSITY/constants"
	"AS-UNIVERSITY/controller"
	"AS-UNIVERSITY/middleware"
	m "AS-UNIVERSITY/middleware"

	"github.com/labstack/echo/v4"
	midd "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	//CRUD ADMIN
	e.GET("/admins", controller.GetAdminsController)
	e.GET("/admins/:id", controller.GetAdminsByID)
	m.LogMiddleware(e)
	e.POST("/admins", controller.CreateAdminController)
	e.DELETE("/admins/:id", controller.DeleteAdmin)
	e.PUT("/admins/:id", controller.UpdateAdmin)
	e.POST("/login", controller.LoginAdminController)

	eAuthBasic := e.Group("/auth")
	eAuthBasic.Use(midd.BasicAuth(middleware.BasicAuthDB))
	eAuthBasic.GET("/admins", controller.GetAdminsController)

	eJwt := e.Group("/jwt")
	eJwt.Use(midd.JWT([]byte(constants.SECRET_JWT)))
	eJwt.GET("/admins", controller.GetAdminsController)

	//CRUD AKREDITASI
	e.GET("/akreditas", controller.GetAkreditas)
	e.GET("/akreditas/:id", controller.GetAkreditasByID)
	e.POST("/akreditas", controller.CreateAkreditas)
	e.DELETE("/akreditas/:id", controller.DeleteAkreditas)
	e.PUT("/akreditas/:id", controller.UpdateAkreditas)

	//CRUD DESKRIPSI
	e.GET("/deskripsi", controller.GetAkreditas)
	e.GET("/deskripsi/:id", controller.GetAkreditasByID)
	e.POST("/deskripsi", controller.CreateAkreditas)
	e.DELETE("/deskripsi/:id", controller.DeleteAkreditas)
	e.PUT("/deskripsi/:id", controller.UpdateAkreditas)

	//CRUD FAKULTAS
	e.GET("/fakultas", controller.GetFakultas)
	e.GET("/fakultas/:id", controller.GetFakultasByID)
	e.POST("/fakultas", controller.CreateFakultas)
	e.DELETE("/fakultas/:id", controller.DeleteFakultas)
	e.PUT("/fakultas/:id", controller.UpdateFakultas)

	//CRUD FASILITAS JURUSAN
	e.GET("/fasilitasjurusan", controller.GetFasilitasJurusan)
	e.GET("/fasilitasjurusan/:id", controller.GetFasilitasJurusanByID)
	e.POST("/fasilitasjurusan", controller.CreateFasilitasJurusan)
	e.DELETE("/fasilitasjurusan/:id", controller.DeleteFasilitasJurusan)
	e.PUT("/fasilitasjurusan/:id", controller.UpdateFasilitasJurusan)

	//CRUD FASILITAS KAMPUS
	e.GET("/fasilitaskampus", controller.GetFasilitasKampus)
	e.GET("/fasilitaskampus/:id", controller.GetFasilitasKampusByID)
	e.POST("/fasilitaskampus", controller.CreateFasilitasKampus)
	e.DELETE("/fasilitaskampus/:id", controller.DeleteFasilitasKampus)
	e.PUT("/fasilitaskampus/:id", controller.UpdateFasilitasKampus)

	//CRUD FASILITAS KELAS
	e.GET("/fasilitaskelas", controller.GetFasilitasKelas)
	e.GET("/fasilitaskelas/:id", controller.GetFasilitasKelasByID)
	e.POST("/fasilitaskelas", controller.CreateFasilitasKelas)
	e.DELETE("/fasilitaskelas/:id", controller.DeleteFasilitasKelas)
	e.PUT("/fasilitaskelas/:id", controller.UpdateFasilitasKelas)

	//CRUD JURUSAN
	e.GET("/jurusan", controller.GetJurusan)
	e.GET("/jurusan/:id", controller.GetJurusanByID)
	e.POST("/jurusan", controller.CreateJurusan)
	e.DELETE("/jurusan/:id", controller.DeleteJurusan)
	e.PUT("/jurusan/:id", controller.UpdateJurusan)

	//CRUD UKM
	e.GET("/ukm", controller.GetUKM)
	e.GET("/ukm/:id", controller.GetUKMByID)
	e.POST("/ukm", controller.CreateUKM)
	e.DELETE("/jurusan/:id", controller.DeleteUKM)
	e.PUT("/jurusan/:id", controller.UpdateUKM)

	//CRUD UNIVERSITAS
	e.GET("/universitas", controller.GetUniv)
	e.GET("/universitas/:id", controller.GetUnivByID)
	e.POST("/universitas", controller.CreateUniv)
	e.DELETE("/universitas/:id", controller.DeleteUniv)
	e.PUT("/universitas/:id", controller.UpdateUniv)

	return e
}
