package data

import (
	"dami-api/feature/lpj"
	"dami-api/helper"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
)

type DataLpj struct {
	db *sql.DB
}

func (db *DataLpj) Select() (data []lpj.LpjEntity, err error) {
	rows, err := db.db.Query("Select * From lpj")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		lpj := lpj.LpjEntity{}

		err = rows.Scan(&lpj.Id, &lpj.NoLpj, &lpj.TanggalLpj, &lpj.NoTpb.NomerTpb, &lpj.NoTpb.TanggalTpb, &lpj.UraianTbp, &lpj.Nilai, &lpj.Tahapan, &lpj.TanggalVerify, &lpj.TanggalPengesahan)

		if err != nil {
			return nil, err
		}

		data = append(data, lpj)
	}

	return data, nil
}

func (db *DataLpj) Create(payload lpj.LpjEntity) (lpj.LpjEntity, error) {
	err := db.db.QueryRow("Insert Into lpj (nomor_lpj,tanggal_lpj,nomor_tbp,tanggal_tbp,uraian_tbp,nilai,tahapan,tanggal_verifikasi_lpj,tanggal_pengesahan_lpj) Values($1,$2,$3,$4,$5,$6,$7,$8,$9) Returning id", payload.NoLpj, payload.TanggalLpj, payload.NoTpb.NomerTpb, payload.NoTpb.TanggalTpb, payload.UraianTbp, payload.Nilai, payload.Tahapan, payload.TanggalVerify, payload.TanggalPengesahan).Scan(&payload.Id)

	if err != nil {
		return lpj.LpjEntity{}, err
	}

	return payload, nil
}

func (db *DataLpj) Update(payload lpj.LpjEntity) (lpj.LpjEntity, error) {
	err := db.db.QueryRow("Select id From lpj Where id=$1", payload.Id).Scan(&payload.Id)
	fmt.Println("Tanggal Lpj :", payload.TanggalLpj)
	if err != nil {
		return lpj.LpjEntity{}, errors.New("data not found")
	}

	var data []any
	index := 1
	qry := "Update lpj Set "

	if payload.NoLpj != "" {
		qry += "nomor_lpj=$" + strconv.Itoa(index)
		data = append(data, payload.NoLpj)
		index++
	}

	if payload.TanggalLpj.After(helper.TimeParseing()) {
		if index > 1 {
			index++
			qry += ",tanggal_lpj=$" + strconv.Itoa(index)
			data = append(data, payload.TanggalLpj)
		} else {
			qry += "tanggal_lpj=$" + strconv.Itoa(index)
			data = append(data, payload.TanggalLpj)
			index++
		}
	}

	if payload.NoTpb.NomerTpb != "" {
		if index > 1 {
			index++
			qry += ",nomor_tbp=$" + strconv.Itoa(index)
			data = append(data, payload.NoTpb.NomerTpb)
		} else {
			qry += "nomor_tbp=$" + strconv.Itoa(index)
			data = append(data, payload.NoTpb.NomerTpb)
			index++
		}
	}

	if payload.TanggalLpj.After(helper.TimeParseing()) {
		fmt.Println(payload.TanggalLpj.After(helper.TimeParseing()))
		if index > 1 {
			index++
			qry += ",tanggal_tbp=$" + strconv.Itoa(index)
			data = append(data, payload.NoTpb.TanggalTpb)
		} else {
			qry += "tanggal_tbp=$" + strconv.Itoa(index)
			data = append(data, payload.NoTpb.TanggalTpb)
			index++
		}
	}

	if payload.UraianTbp != "" {
		if index > 1 {
			index++
			qry += ",uraian_tbp=$" + strconv.Itoa(index)
			data = append(data, payload.UraianTbp)
		} else {
			qry += "uraian_tbp=$" + strconv.Itoa(index)
			data = append(data, payload.UraianTbp)
			index++
		}
	}

	if payload.Nilai != 0 {
		if index > 1 {
			index++
			qry += ",nilai=$" + strconv.Itoa(index)
			data = append(data, payload.Nilai)
		} else {
			qry += "nilai=$" + strconv.Itoa(index)
			data = append(data, payload.Nilai)
			index++
		}
	}

	if payload.Tahapan != "" {
		if index > 1 {
			index++
			qry += ",tahapan=$" + strconv.Itoa(index)
			data = append(data, payload.Tahapan)
		} else {
			qry += "tahapan=$" + strconv.Itoa(index)
			data = append(data, payload.Tahapan)
			index++
		}
	}

	if payload.TanggalLpj.After(helper.TimeParseing()) {
		if index > 1 {
			index++
			qry += ",tanggal_verifikasi_lpj=$" + strconv.Itoa(index)
			data = append(data, payload.TanggalVerify)
		} else {
			qry += "tanggal_verifikasi_lpj=$" + strconv.Itoa(index)
			data = append(data, payload.TanggalVerify)
			index++
		}
	}

	if payload.TanggalLpj.After(helper.TimeParseing()) {
		if index > 1 {
			index++
			qry += ",tanggal_pengesahan_lpj=$" + strconv.Itoa(index)
			data = append(data, payload.TanggalPengesahan)
		} else {
			qry += "tanggal_pengesahan_lpj=$" + strconv.Itoa(index)
			data = append(data, payload.TanggalPengesahan)
			index++
		}
	}

	qry += " Where id=$" + strconv.Itoa(index) + " Returning id,nomor_lpj,tanggal_lpj,nomor_tbp,tanggal_tbp,uraian_tbp,nilai,tahapan,tanggal_verifikasi_lpj,tanggal_pengesahan_lpj"
	data = append(data, payload.Id)

	err = db.db.QueryRow(qry, data...).Scan(&payload.Id, &payload.NoLpj, &payload.TanggalLpj, &payload.NoTpb.NomerTpb, &payload.NoTpb.TanggalTpb, &payload.UraianTbp, &payload.Nilai, &payload.Tahapan, &payload.TanggalVerify, &payload.TanggalPengesahan)

	if err != nil {
		return lpj.LpjEntity{}, err
	}

	return payload, nil
}

func (db *DataLpj) Delete(lpjID string) (payloadLpj lpj.LpjEntity, err error) {
	err = db.db.QueryRow("Delete From lpj Where id=$1 Returning nomor_lpj,tanggal_lpj,nomor_tbp,tanggal_tbp,uraian_tbp,nilai,tahapan,tanggal_verifikasi_lpj,tanggal_pengesahan_lpj", lpjID).Scan(&payloadLpj.NoLpj, &payloadLpj.TanggalLpj, &payloadLpj.NoTpb.NomerTpb, &payloadLpj.NoTpb.TanggalTpb, &payloadLpj.UraianTbp, &payloadLpj.Nilai, &payloadLpj.Tahapan, &payloadLpj.TanggalVerify, &payloadLpj.TanggalPengesahan)

	return
}

func NewLpjQuery(db *sql.DB) lpj.DataLpjInterface {
	return &DataLpj{db}
}
