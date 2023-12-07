package handler

import (
	"dami-api/feature/lpj"
	"dami-api/feature/tpb/data"
	"time"
)

type RequestLpj struct {
	NoLpj             string    `json:"noLpj"`
	TanggalLpj        time.Time `json:"tanggalLpj"`
	NoTpb             string    `json:"noTpb"`
	UraianTbp         string    `json:"uraianTpb"`
	Nilai             int       `json:"nilai"`
	Tahapan           string    `json:"tahapan"`
	TanggalVerify     time.Time `json:"tanggalVerify"`
	TanggalPengesahan time.Time `json:"tanggalPengesahan"`
}

func LpjRequset(payloadLpj RequestLpj) lpj.LpjEntity {
	return lpj.LpjEntity{
		NoLpj:      payloadLpj.NoLpj,
		TanggalLpj: payloadLpj.TanggalLpj,
		NoTpb: data.Tpb{
			Id: payloadLpj.NoTpb,
		},
		UraianTbp:         payloadLpj.UraianTbp,
		Nilai:             payloadLpj.Nilai,
		Tahapan:           payloadLpj.Tahapan,
		TanggalVerify:     payloadLpj.TanggalVerify,
		TanggalPengesahan: payloadLpj.TanggalPengesahan,
	}
}
