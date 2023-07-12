package pii

import (
	"context"

	domaincounties "github.com/firdasafridi/stat-hybrid/internal/entity/countries"
	"github.com/firdasafridi/stat-hybrid/internal/repo/api/countries"
)

type countriesUC interface {
	GetCounties(ctx context.Context, country string) (detailCountries []domaincounties.ResponseCountry, err error)
}

func (uc *PII) GetCounties(ctx context.Context, country string) (detailCountries []domaincounties.ResponseCountry, err error) {
	return countries.RequestCountry(ctx, country)
}
