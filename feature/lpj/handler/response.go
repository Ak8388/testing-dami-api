package handler

import (
	"dami-api/feature/lpj"
	"time"
)

type LpjResponseModel struct {
	Id                string    `json:"id"`
	NoLpj             string    `json:"noLpj"`
	TanggalLpj        time.Time `json:"tanggalLpj"`
	NoTpb             string    `json:"noTpb"`
	TanggalTbp        time.Time `json:"tanggalTbp"`
	UraianTbp         string    `json:"uraianTpb"`
	Nilai             int       `json:"nilai"`
	Tahapan           string    `json:"tahapan"`
	TanggalVerify     time.Time `json:"tanggalVerify"`
	TanggalPengesahan time.Time `json:"tanggalPengesahan"`
}

func LpjResponse(Entity lpj.LpjEntity) LpjResponseModel {
	return LpjResponseModel{
		Id:                Entity.Id,
		NoLpj:             Entity.NoLpj,
		TanggalLpj:        Entity.TanggalLpj,
		NoTpb:             Entity.NoTpb.NomerTpb,
		TanggalTbp:        Entity.NoTpb.TanggalTpb,
		UraianTbp:         Entity.UraianTbp,
		Nilai:             Entity.Nilai,
		Tahapan:           Entity.Tahapan,
		TanggalVerify:     Entity.TanggalVerify,
		TanggalPengesahan: Entity.TanggalPengesahan,
	}
}

func LpjResponseMulti(Entity []lpj.LpjEntity) (res []LpjResponseModel) {
	for _, v := range Entity {
		lpj := LpjResponseModel{
			Id:                v.Id,
			NoLpj:             v.NoLpj,
			TanggalLpj:        v.TanggalLpj,
			NoTpb:             v.NoTpb.NomerTpb,
			TanggalTbp:        v.NoTpb.TanggalTpb,
			UraianTbp:         v.UraianTbp,
			Nilai:             v.Nilai,
			Tahapan:           v.Tahapan,
			TanggalVerify:     v.TanggalVerify,
			TanggalPengesahan: v.TanggalPengesahan,
		}

		res = append(res, lpj)
	}

	return
}
