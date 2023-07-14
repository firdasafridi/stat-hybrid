package hybridencryption

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"hash"
	"log"
	"reflect"

	"github.com/pkg/errors"
)

var (
	ErrCheckKey    = errors.New("failed to parse PEM block containing the key")
	ErrStructCrypt = errors.New("please call new structcrypt")
	ErrPublicKey   = errors.New("please set the public key")
	ErrPrivateKey  = errors.New("please set the private key")
	ErrNotStruct   = errors.New("struct is not pointer")
	ErrStructIsNil = errors.New("struct is nil")
)

type RSAOption struct {
	PubKey     string
	PrivateKey string
}

type RSA struct {
	privKey  *rsa.PrivateKey
	pubKey   *rsa.PublicKey
	hashData hash.Hash
}

func NewRSA(opt RSAOption) (*RSA, error) {
	crypt := &RSA{}

	if opt.PubKey != "" {
		// The public key is a part of the *rsa.PrivateKey struct
		publicKey, err := parseRsaPublicKeyFromPemStr(opt.PubKey)
		if err != nil {
			return nil, errors.Wrap(err, "parseRsaPublicKeyFromPemStr")
		}
		crypt.pubKey = publicKey
	}

	if opt.PrivateKey != "" {
		// The GenerateKey method takes in a reader that returns random bits, and
		// the number of bits
		privateKey, err := parseRsaPrivateKeyFromPemStr(opt.PrivateKey)
		if err != nil {
			return nil, err
		}
		crypt.privKey = privateKey
	}

	crypt.hashData = sha512.New()

	return crypt, nil
}

func parseRsaPublicKeyFromPemStr(pubPEM string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pubPEM))
	if block == nil {
		return nil, ErrCheckKey
	}

	pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return pub, nil
}

func parseRsaPrivateKeyFromPemStr(privPEM string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privPEM))
	if block == nil {
		return nil, ErrCheckKey
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return priv, nil
}

func (crypt *RSA) Encrypt(fromData []byte) (string, error) {
	if crypt == nil {
		return "", ErrStructCrypt
	}

	if crypt.pubKey == nil {
		return "", ErrPublicKey
	}

	encryptedBytes, err := rsa.EncryptOAEP(
		crypt.hashData,
		rand.Reader,
		crypt.pubKey,
		fromData,
		nil)
	if err != nil {
		return "", errors.Wrap(err, "rsa.EncryptOAEP")
	}

	bs64EncryptedBytes := base64.StdEncoding.EncodeToString(encryptedBytes)

	return bs64EncryptedBytes, nil
}

func (crypt *RSA) Decrypt(fromData []byte) (string, error) {
	if crypt == nil {
		return "", ErrStructCrypt
	}

	if crypt.privKey == nil {
		return "", ErrPrivateKey
	}

	encryptedBytes, err := base64.StdEncoding.DecodeString(string(fromData))
	if err != nil {
		return "", errors.Wrap(err, "base64.StdEncoding.DecodeString")
	}

	decryptedBytes, err := crypt.privKey.Decrypt(nil, encryptedBytes, &rsa.OAEPOptions{Hash: crypto.SHA512})
	if err != nil {
		log.Fatalln(err)
	}

	return string(decryptedBytes), nil
}

func (crypt *RSA) EncryptFromStruct(fromStruct interface{}) (encryptData string, err error) {
	flatBytes, err := json.Marshal(fromStruct)
	if err != nil {
		return "", errors.Wrap(err, "json.Marshal")
	}

	encryptStr, err := crypt.Encrypt(flatBytes)
	if err != nil {
		return "", errors.Wrap(err, "crypt.Decrypt")
	}

	return encryptStr, nil
}

func (crypt *RSA) Unmarshal(fromStruct interface{}, encryptBytes []byte) (err error) {

	if reflect.ValueOf(fromStruct).Kind() != reflect.Ptr {
		return ErrNotStruct
	}

	if reflect.ValueOf(fromStruct).IsNil() {
		return ErrStructIsNil
	}

	encryptStr, err := crypt.Decrypt(encryptBytes)
	if err != nil {
		return errors.Wrap(err, "crypt.Decrypt")
	}

	err = json.Unmarshal([]byte(encryptStr), fromStruct)
	if err != nil {
		return errors.Wrap(err, "crypt.Decrypt")
	}

	return nil
}
