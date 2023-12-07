package data

import "time"

type Tpb struct {
	Id               string    `json:"id"`
	NomerTpb         string    `json:"nomerTbp"`
	TanggalTpb       time.Time `json:"tanggalTbp"`
	TujuanPembayaran string    `json:"tujuanPembayaran"`
	NilaiTpb         int       `json:"nilaiTbp"`
	Tahapan          string    `json:"tahapan"`
}
