package data

import (
	"dami-api/feature/tpb/data"
	"time"
)

type Lpj struct {
	Id                int       `json:"id"`
	NoLpj             string    `json:"noLpj"`
	TanggalLpj        time.Time `json:"tanggalLpj"`
	NoTpb             data.Tpb  `json:"noTpb"`
	UraianTbp         string    `json:"uraianTpb"`
	Nilai             int       `json:"nilai"`
	Tahapan           string    `json:"tahapan"`
	TanggalVerify     time.Time `json:"tanggalVerify"`
	TanggalPengesahan time.Time `json:"tanggalPengesahan"`
}
