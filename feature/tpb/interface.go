package tpb

import "time"

type EntityTbp struct {
	Id               string    `json:"id"`
	NomerTpb         string    `json:"nomerTbp"`
	TanggalTpb       time.Time `json:"tanggalTbp"`
	TujuanPembayaran string    `json:"tujuanPembayaran"`
	NilaiTpb         int       `json:"nilaiTbp"`
	Tahapan          string    `json:"tahapan"`
}

type DataTbpInterface interface {
	SelectTbpById(string) (EntityTbp, error)
}

type ServiceTbpInterface interface {
	SelectTbpById(string) (EntityTbp, error)
}
