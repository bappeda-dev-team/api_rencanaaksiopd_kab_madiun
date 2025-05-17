// helper/model.go
package helper

import (
	"renaksiopdService/model/domain"
	"renaksiopdService/model/web"
)

func ToRencanaAksiOpdResponse(rencanaAksi domain.RencanaAksiOpd) web.RencanaAksiOpdResponse {
	return web.RencanaAksiOpdResponse{
		Id:             rencanaAksi.Id,
		SasaranOpdId:   rencanaAksi.SasaranOpdId,
		NamaSasaranOpd: rencanaAksi.NamaSasaranOpd,
		TahunRenaksi:   rencanaAksi.TahunRenaksi,
		Tw1:            rencanaAksi.Tw1,
		Tw2:            rencanaAksi.Tw2,
		Tw3:            rencanaAksi.Tw3,
		Tw4:            rencanaAksi.Tw4,
		Keterangan:     rencanaAksi.Keterangan,
		RencanaKinerja: ToRencanaKinerjaResponses(rencanaAksi.RencanaKinerja),
	}
}

func ToRencanaKinerjaResponses(rencanaKinerja []domain.RencanaKinerjaOpd) []web.RencanaKinerjaResponse {
	var responses []web.RencanaKinerjaResponse
	for _, rk := range rencanaKinerja {
		responses = append(responses, ToRencanaKinerjaResponse(rk))
	}
	return responses
}

func ToRencanaKinerjaResponse(rk domain.RencanaKinerjaOpd) web.RencanaKinerjaResponse {
	return web.RencanaKinerjaResponse{
		RekinId:            rk.RekinId,
		NamaRencanaKinerja: rk.NamaRencanaKinerja,
		NipPegawai:         rk.NipPegawai,
		NamaPegawai:        rk.NamaPegawai,
		KodeOpd:            rk.KodeOpd,
		TotalAnggaran:      rk.TotalAnggaran,
		SubKegiatan:        ToSubKegiatanResponses(rk.SubKegiatan),
	}
}

func ToSubKegiatanResponses(subKegiatan []domain.SubKegiatanOpd) []web.SubKegiatanResponse {
	var responses []web.SubKegiatanResponse
	for _, sk := range subKegiatan {
		responses = append(responses, ToSubKegiatanResponse(sk))
	}
	return responses
}

func ToSubKegiatanResponse(sk domain.SubKegiatanOpd) web.SubKegiatanResponse {
	return web.SubKegiatanResponse{
		KodeSubKegiatan: sk.KodeSubKegiatan,
		NamaSubKegiatan: sk.NamaSubKegiatan,
		Indikator:       ToIndikatorResponses(sk.Indikator),
	}
}

func ToIndikatorResponses(indikator []domain.IndikatorSubKegiatanOpd) []web.IndikatorResponse {
	var responses []web.IndikatorResponse
	for _, i := range indikator {
		responses = append(responses, ToIndikatorResponse(i))
	}
	return responses
}

func ToIndikatorResponse(i domain.IndikatorSubKegiatanOpd) web.IndikatorResponse {
	return web.IndikatorResponse{
		Id:        i.Id,
		Indikator: i.Indikator,
		Target:    i.Target,
		Satuan:    i.Satuan,
	}
}

func ToRencanaAksiOpdResponses(rencanaAksi []domain.RencanaAksiOpd) []web.RencanaAksiOpdResponse {
	var responses []web.RencanaAksiOpdResponse
	for _, ra := range rencanaAksi {
		responses = append(responses, ToRencanaAksiOpdResponse(ra))
	}
	return responses
}

func ToRencanaAksiOpdRequestResponse(rencanaAksi domain.RencanaAksiOpd) web.RencanaAksiOpdRequestResponse {
	return web.RencanaAksiOpdRequestResponse{
		SasaranOpdId: rencanaAksi.SasaranOpdId,
		RekinId:      rencanaAksi.RekinId,
		TahunRenaksi: rencanaAksi.TahunRenaksi,
		Keterangan:   rencanaAksi.Keterangan,
	}
}

func ToRencanaAksiOpdByIdResponse(rencanaAksi domain.RencanaAksiOpd) web.RencanaAksiOpdByIdResponse {
	return web.RencanaAksiOpdByIdResponse{
		Id:                 rencanaAksi.Id,
		RekinId:            rencanaAksi.RekinId,
		TahunRenaksi:       rencanaAksi.TahunRenaksi,
		Keterangan:         rencanaAksi.Keterangan,
		NamaRencanaKinerja: rencanaAksi.NamaRencanaKinerja,
		SasaranOpd:         ToSasaranOpdDetailResponse(rencanaAksi.SasaranOpd),
	}
}

func ToSasaranOpdDetailResponse(sasaranOpd domain.SasaranOpdDetail) web.SasaranOpdDetailResponse {
	return web.SasaranOpdDetailResponse{
		Id:             sasaranOpd.Id,
		NamaSasaranOpd: sasaranOpd.NamaSasaranOpd,
		TahunAwal:      sasaranOpd.TahunAwal,
		TahunAkhir:     sasaranOpd.TahunAkhir,
		JenisPeriode:   sasaranOpd.JenisPeriode,
		Indikator:      ToIndikatorSasaranOpdResponses(sasaranOpd.Indikator),
	}
}

func ToIndikatorSasaranOpdResponses(indikators []domain.IndikatorSasaranOpd) []web.IndikatorSasaranOpdResponse {
	var responses []web.IndikatorSasaranOpdResponse
	for _, indikator := range indikators {
		responses = append(responses, ToIndikatorSasaranOpdResponse(indikator))
	}
	return responses
}

func ToIndikatorSasaranOpdResponse(indikator domain.IndikatorSasaranOpd) web.IndikatorSasaranOpdResponse {
	return web.IndikatorSasaranOpdResponse{
		Id:               indikator.Id,
		Indikator:        indikator.Indikator,
		RumusPerhitungan: indikator.RumusPerhitungan,
		SumberData:       indikator.SumberData,
		Target:           ToTargetResponse(indikator.Target),
	}
}

func ToTargetResponse(target domain.Target) web.TargetResponse {
	return web.TargetResponse{
		Id:          target.Id,
		IndikatorId: target.IndikatorId,
		Tahun:       target.Tahun,
		Target:      target.Target,
		Satuan:      target.Satuan,
	}
}
