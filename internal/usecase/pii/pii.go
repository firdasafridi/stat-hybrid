package pii

import (
	"context"

	piidomain "github.com/firdasafridi/stat-hybrid/internal/entity/pii"
)

type PIIUC interface {
	GetPIIData(ctx context.Context) (piiList []piidomain.TrxPII, err error)
	countriesUC
}

type PII struct {
}

func New(pii *PII) *PII {
	return pii
}


func (uc *PII) GetPIIData(ctx context.Context) (piiList []piidomain.TrxPII, err error) {
	return []piidomain.TrxPII{
		{
			IDCard: "637123123123123123",
			FullName: "Manusia Baja Hitam",
			PhoneNumber: "123SAMPLEPHONE",
		},
		{
			IDCard: "637123123123123123",
			FullName: "Power Ranger Kuning",
			PhoneNumber: "123SAMPLEPHONE",
		},
	}, nil
}
