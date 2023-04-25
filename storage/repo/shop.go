package repo

import (
	"context"
	"genproto/common"
)

type ShopI interface {
	Upsert(*common.ShopCreatedModel) error
	Delete(ctx context.Context) error
}
