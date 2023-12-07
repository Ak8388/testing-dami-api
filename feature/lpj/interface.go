package lpj

import (
	"dami-api/feature/tpb/data"
	"time"
)

type LpjEntity struct {
	Id                string    `json:"id"`
	NoLpj             string    `json:"noLpj"`
	TanggalLpj        time.Time `json:"tanggalLpj"`
	NoTpb             data.Tpb  `json:"noTpb"`
	UraianTbp         string    `json:"uraianTpb"`
	Nilai             int       `json:"nilai"`
	Tahapan           string    `json:"tahapan"`
	TanggalVerify     time.Time `json:"tanggalVerify"`
	TanggalPengesahan time.Time `json:"tanggalPengesahan"`
}

type DataLpjInterface interface {
	Select() ([]LpjEntity, error)
	Create(LpjEntity) (LpjEntity, error)
	Update(LpjEntity) (LpjEntity, error)
	Delete(string) (LpjEntity, error)
}

type ServiceLpjInterface interface {
	Select() ([]LpjEntity, error)
	Create(LpjEntity) (LpjEntity, error)
	Update(LpjEntity) (LpjEntity, error)
	Delete(string) (LpjEntity, error)
}
