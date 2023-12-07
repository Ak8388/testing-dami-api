package handler

import (
	"dami-api/feature/tpb/data"
	"time"
)

type TpbRequest struct {
	NomerTpb         string    `json:"nomerTbp"`
	TanggalTpb       time.Time `json:"tanggalTbp"`
	TujuanPembayaran string    `json:"tujuanPembayaran"`
	NilaiTpb         int       `json:"nilaiTbp"`
	Tahapan          string    `json:"tahapan"`
}

func TbpRequest(tbpReq TpbRequest) data.Tpb {
	return data.Tpb{
		NomerTpb:         tbpReq.NomerTpb,
		TanggalTpb:       tbpReq.TanggalTpb,
		TujuanPembayaran: tbpReq.TujuanPembayaran,
		NilaiTpb:         tbpReq.NilaiTpb,
		Tahapan:          tbpReq.Tahapan,
	}
}
