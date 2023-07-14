package hybridencryption

import (
	"crypto/aes"
	"encoding/base64"
)

type AES struct {
	Key []byte
}

// NewAESDefault is function to create new configuration of aes algorithm option
func NewAES(key []byte) (*AES, error) {
	return &AES{
		Key: key,
	}, nil
}

func (aesDefault *AES) AESEncrypt(src []byte) (encrypted []byte, err error) {
	cipher, err := aes.NewCipher(aesDefault.Key)
	if err != nil {
		return
	}
	length := (len(src) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, src)
	pad := byte(len(plain) - len(src))
	for i := len(src); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))

	for bs, be := 0, cipher.BlockSize(); bs <= len(src); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	encrypted = Base64Encode(encrypted)

	return encrypted, nil
}

func (aesDefault *AES) AESDecrypt(encrypted []byte) (decrypted []byte, err error) {
	cipher, err := aes.NewCipher(aesDefault.Key)
	if err != nil {
		return decrypted, err
	}
	encrypted, err = Base64Decode(encrypted)
	if err != nil {
		return decrypted, err
	}

	decrypted = make([]byte, len(encrypted))
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim], nil
}

func Base64Encode(message []byte) []byte {
	b := make([]byte, base64.StdEncoding.EncodedLen(len(message)))
	base64.StdEncoding.Encode(b, message)
	return b
}

func Base64Decode(message []byte) (b []byte, err error) {
	var l int
	b = make([]byte, base64.StdEncoding.DecodedLen(len(message)))
	l, err = base64.StdEncoding.Decode(b, message)
	if err != nil {
		return
	}
	return b[:l], nil
}
