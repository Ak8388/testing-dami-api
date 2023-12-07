package service

import (
	"dami-api/feature/tpb"
)

type serviceTbp struct {
	query tpb.ServiceTbpInterface
}

func (s *serviceTbp) SelectTbpById(id string) (tpb.EntityTbp, error) {
	return s.query.SelectTbpById(id)
}

func NewTbpService(query tpb.ServiceTbpInterface) tpb.ServiceTbpInterface {
	return &serviceTbp{query}
}
