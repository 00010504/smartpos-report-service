package repo

import (
	"context"
	"genproto/common"
)

type CompanyI interface {
	Upsert(context.Context, *common.CompanyCreatedModel) error
	Delete(context.Context, *common.RequestID) error
}
