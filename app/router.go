package app

import (
	"renaksiopdService/controller"

	"github.com/julienschmidt/httprouter"
)

type RouteController struct {
}

func NewRouter(
	rencanaAksiOpdController controller.RencanaAksiOpdController,
) *httprouter.Router {
	router := httprouter.New()

	router.GET("/rencana-aksi-opd/:sasaran_opd_id/:tahun", rencanaAksiOpdController.FindBySasaranOpdAndTahun)
	router.POST("/rencana-aksi-opd/sync_jadwal/:rekin_id", rencanaAksiOpdController.SyncJadwalPelaksanaan)
	router.POST("/rencana-aksi-opd/create", rencanaAksiOpdController.Create)
	router.PUT("/rencana-aksi-opd/update/:id", rencanaAksiOpdController.Update)
	router.DELETE("/rencana-aksi-opd/delete/:id", rencanaAksiOpdController.Delete)
	router.GET("/renaksi-opd/detail/:id", rencanaAksiOpdController.FindById)
	router.GET("/sasaran_opd/all/:kode_opd/:tahun", rencanaAksiOpdController.FindAllSasaranByTahun)

	return router
}
