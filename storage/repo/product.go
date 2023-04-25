package repo

import (
	"context"
	"genproto/common"
)

type ProductI interface {
	Upsert(ctx context.Context, entity *common.CreateProductCopyRequest) error
	Delete(ctx context.Context, req *common.RequestID) (*common.ResponseID, error)
}
