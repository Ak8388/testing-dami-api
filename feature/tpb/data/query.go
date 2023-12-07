package data

import (
	"dami-api/feature/tpb"
	"database/sql"
)

type DataTbp struct {
	db *sql.DB
}

func (db *DataTbp) SelectTbpById(id string) (tbp tpb.EntityTbp, err error) {
	err = db.db.QueryRow("Select * From tbp Where id=$1", id).Scan(&tbp.Id, &tbp.NomerTpb, &tbp.TanggalTpb, &tbp.TujuanPembayaran, &tbp.NilaiTpb, &tbp.Tahapan)

	if err != nil {
		return
	}

	return
}

func NewTbp(db *sql.DB) tpb.DataTbpInterface {
	return &DataTbp{
		db: db,
	}
}
