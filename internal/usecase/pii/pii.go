package pii

import (
	"context"
	"encoding/json"

	piidomain "github.com/firdasafridi/stat-hybrid/internal/entity/pii"
	"github.com/firdasafridi/stat-hybrid/lib/hybridencryption"
)

type PIIUC interface {
	GetPIIData(ctx context.Context) (piiList []piidomain.TrxPII, err error)
	GetPIIDataHybridEncrypt(ctx context.Context) (chiperData string, chiperKey string, err error)
	countriesUC
}

type PII struct {
	HBELIB hybridencryption.HBE
}

func New(pii *PII) *PII {
	return pii
}

func (uc *PII) GetPIIData(ctx context.Context) (piiList []piidomain.TrxPII, err error) {
	return []piidomain.TrxPII{
		{
			IDCard:      "637123123123123123",
			FullName:    "Manusia Baja Hitam",
			PhoneNumber: "123SAMPLEPHONE",
		},
		{
			IDCard:      "637123123123123123",
			FullName:    "Power Ranger Kuning",
			PhoneNumber: "123SAMPLEPHONE",
		},
	}, nil
}

func (uc *PII) GetPIIDataHybridEncrypt(ctx context.Context) (chiperData string, chiperKey string, err error) {
	PIIPlainData, err := uc.GetPIIData(ctx)
	if err != nil {
		return
	}

	PIIPlainDataByte, _ := json.Marshal(PIIPlainData)

	chiperData, chiperKey, err = uc.HBELIB.Encrypt(ctx, string(PIIPlainDataByte))
	if err != nil {
		return
	}

	return
}
