package service

import (
	"context"
	"renaksiopdService/model/web"
)

type RencanaAksiOpdService interface {
	FindBySasaranOpdAndTahun(ctx context.Context, sasaranOpdId int, tahun string) ([]web.RencanaAksiOpdResponse, error)
	SyncJadwalPelaksanaan(ctx context.Context, rekinId string) error
	Create(ctx context.Context, request web.RencanaAksiOpdCreateRequest) (web.RencanaAksiOpdRequestResponse, error)
	Update(ctx context.Context, request web.RencanaAksiOpdUpdateRequest) (web.RencanaAksiOpdRequestResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (web.RencanaAksiOpdByIdResponse, error)
	FindAllSasaranByTahun(ctx context.Context, kodeOpd string, tahun string) ([]web.SasaranOpdDetailResponse, error)
}
