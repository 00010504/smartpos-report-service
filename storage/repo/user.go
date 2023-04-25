package repo

import (
	"context"
	"genproto/common"
)

type UserI interface {
	Upsert(ctx context.Context, entity *common.UserCreatedModel) error
	Delete(ctx context.Context, req *common.RequestID) error
}
