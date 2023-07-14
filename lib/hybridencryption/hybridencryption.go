package hybridencryption

import (
	"context"

	"github.com/firdasafridi/stat-hybrid/lib/common"
)

type HBE interface {
	Encrypt(ctx context.Context, plain string) (chiper string, key string, err error)
}

type HybridEncryption struct {
	RSARepo *RSA
}

// HybridEncryption is function to create new configuration of aes algorithm option
func NewHybridEncryption(rsaOption RSAOption) (hb *HybridEncryption, err error) {
	rsa, err := NewRSA(rsaOption)
	if err != nil {
		return
	}

	return &HybridEncryption{
		RSARepo: rsa,
	}, nil
}

func (h *HybridEncryption) Encrypt(ctx context.Context, plain string) (chiperData string, chiperKey string, err error) {
	aes, err := NewAES([]byte(common.RandomString(32)))
	if err != nil {
		return
	}

	chiperKey, err = h.RSARepo.Encrypt([]byte(aes.Key))
	if err != nil {
		return
	}

	chiperDataByte, err := aes.AESEncrypt([]byte(plain))
	if err != nil {
		return
	}
	chiperData = string(chiperDataByte)
	return
}
