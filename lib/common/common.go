package common

import (
	"fmt"
	"math/rand"
	"net/url"
	"time"

	"github.com/firdasafridi/stat-hybrid/lib/common/commonerr"
	"github.com/gorilla/schema"
)

type ctxKey string

const (
	YYYYMMDDDash = "2006-01-02"
)

func DecodeSchema(values url.Values, val interface{}) error {
	decoder := schema.NewDecoder()
	if err := decoder.Decode(val, values); err != nil {
		return commonerr.SetNewBadRequest("URL Params", err.Error())
	}
	return nil
}

func RandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}

func SetCustomKey(key string) ctxKey {
	return ctxKey(key)
}
