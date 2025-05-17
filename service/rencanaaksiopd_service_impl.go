package service

import (
	"context"
	"database/sql"
	"math/rand"
	"renaksiopdService/helper"
	"renaksiopdService/model/domain"
	"renaksiopdService/model/web"
	"renaksiopdService/repository"

	"github.com/go-playground/validator/v10"
)

type RencanaAksiOpdServiceImpl struct {
	RencanaAksiOpdRepository repository.RencanaAksiOpdRepository
	DB                       *sql.DB
	validator                *validator.Validate
}

func NewRencanaAksiOpdServiceImpl(rencanaAksiOpdRepository repository.RencanaAksiOpdRepository, db *sql.DB, validator *validator.Validate) *RencanaAksiOpdServiceImpl {
	return &RencanaAksiOpdServiceImpl{
		RencanaAksiOpdRepository: rencanaAksiOpdRepository,
		DB:                       db,
		validator:                validator,
	}
}

func (service *RencanaAksiOpdServiceImpl) FindBySasaranOpdAndTahun(ctx context.Context, sasaranOpdId int, tahun string) ([]web.RencanaAksiOpdResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	rencanaAksi, err := service.RencanaAksiOpdRepository.FindBySasaranOpdAndTahun(ctx, tx, sasaranOpdId, tahun)
	if err != nil {
		return nil, err
	}

	return helper.ToRencanaAksiOpdResponses(rencanaAksi), nil
}

func (service *RencanaAksiOpdServiceImpl) SyncJadwalPelaksanaan(ctx context.Context, rekinId string) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	err = service.RencanaAksiOpdRepository.SyncJadwalPelaksanaan(ctx, tx, rekinId)
	if err != nil {
		return err
	}

	return nil
}

func (service *RencanaAksiOpdServiceImpl) Create(ctx context.Context, request web.RencanaAksiOpdCreateRequest) (web.RencanaAksiOpdRequestResponse, error) {
	err := service.validator.Struct(request)
	if err != nil {
		return web.RencanaAksiOpdRequestResponse{}, err
	}

	tx, err := service.DB.Begin()
	if err != nil {
		return web.RencanaAksiOpdRequestResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	randomInt := rand.Intn(10000000)

	rencanaAksiOpdDomain := domain.RencanaAksiOpd{
		Id:           randomInt,
		RekinId:      request.RekinId,
		SasaranOpdId: request.SasaranOpdId,
		TahunRenaksi: request.TahunRenaksi,
		Keterangan:   &request.Keterangan,
	}

	rencanaAksiOpd, err := service.RencanaAksiOpdRepository.Create(ctx, tx, rencanaAksiOpdDomain)
	if err != nil {
		return web.RencanaAksiOpdRequestResponse{}, err
	}

	return helper.ToRencanaAksiOpdRequestResponse(rencanaAksiOpd), nil
}

func (service *RencanaAksiOpdServiceImpl) Update(ctx context.Context, request web.RencanaAksiOpdUpdateRequest) (web.RencanaAksiOpdRequestResponse, error) {
	err := service.validator.Struct(request)
	if err != nil {
		return web.RencanaAksiOpdRequestResponse{}, err
	}

	tx, err := service.DB.Begin()
	if err != nil {
		return web.RencanaAksiOpdRequestResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	rencanaAksiOpdDomain := domain.RencanaAksiOpd{
		Id:         request.Id,
		RekinId:    request.RekinId,
		Keterangan: &request.Keterangan,
	}

	rencanaAksiOpd := service.RencanaAksiOpdRepository.Update(ctx, tx, rencanaAksiOpdDomain)

	return helper.ToRencanaAksiOpdRequestResponse(rencanaAksiOpd), nil
}

func (service *RencanaAksiOpdServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	err = service.RencanaAksiOpdRepository.Delete(ctx, tx, id)
	if err != nil {
		return err
	}

	return nil
}

func (service *RencanaAksiOpdServiceImpl) FindById(ctx context.Context, id int) (web.RencanaAksiOpdByIdResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return web.RencanaAksiOpdByIdResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	rencanaAksiOpd, err := service.RencanaAksiOpdRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.RencanaAksiOpdByIdResponse{}, err
	}

	return helper.ToRencanaAksiOpdByIdResponse(rencanaAksiOpd), nil
}

func (service *RencanaAksiOpdServiceImpl) FindAllSasaranByTahun(ctx context.Context, kodeOpd string, tahun string) ([]web.SasaranOpdDetailResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	sasaranList, err := service.RencanaAksiOpdRepository.FindAllSasaranByTahun(ctx, tx, kodeOpd, tahun)
	if err != nil {
		return nil, err
	}

	var responses []web.SasaranOpdDetailResponse
	for _, sasaran := range sasaranList {
		responses = append(responses, helper.ToSasaranOpdDetailResponse(sasaran))
	}

	return responses, nil
}
