package service

import (
	"dami-api/feature/lpj"
	"dami-api/feature/tpb"
	"errors"
)

type service struct {
	query      lpj.DataLpjInterface
	tbpService tpb.ServiceTbpInterface
}

func (s *service) Select() ([]lpj.LpjEntity, error) {
	return s.query.Select()
}

func (s *service) Create(lpjPayload lpj.LpjEntity) (lpj.LpjEntity, error) {
	tbp, err := s.tbpService.SelectTbpById(lpjPayload.NoTpb.Id)

	if err != nil {
		return lpj.LpjEntity{}, errors.New("tbp not found " + err.Error())
	}

	lpjPayload.NoTpb.TanggalTpb = tbp.TanggalTpb
	lpjPayload.NoTpb.NomerTpb = tbp.NomerTpb

	return s.query.Create(lpjPayload)
}

func (s *service) Update(lpj lpj.LpjEntity) (lpj.LpjEntity, error) {
	return s.query.Update(lpj)
}

func (s *service) Delete(lpjId string) (lpj.LpjEntity, error) {
	return s.query.Delete(lpjId)
}

func NewLpjService(query lpj.DataLpjInterface, tbpService tpb.ServiceTbpInterface) lpj.ServiceLpjInterface {
	return &service{
		query:      query,
		tbpService: tbpService,
	}
}
