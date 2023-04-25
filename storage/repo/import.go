package repo

import "context"

type ImportI interface {
	Create(ctx context.Context) error
}
