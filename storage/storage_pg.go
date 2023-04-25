package storage

import (
	"context"

	"github.com/Invan2/invan_report_service/pkg/logger"
	"github.com/Invan2/invan_report_service/storage/postgres"
	"github.com/Invan2/invan_report_service/storage/repo"
	"github.com/jackc/pgx/v5"
)

type repos struct {
	companyRepo repo.CompanyI
	userRepo    repo.UserI
	shopRepo    repo.ShopI
	productRepo repo.ProductI
}

type repoIs interface {
	Company() repo.CompanyI
	User() repo.UserI
	Shop() repo.ShopI
	Product() repo.ProductI
}

type storage struct {
	db  *pgx.Conn
	log logger.Logger
	repos
}

type storageTr struct {
	tr pgx.Tx
	repos
}

type StorageTrI interface {
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	repoIs
}

type StoragePgI interface {
	Product() repo.ProductI
	WithTransaction(ctx context.Context) (StorageTrI, error)
	repoIs
}

func NewStoragePg(log logger.Logger, db *pgx.Conn) StoragePgI {

	return &storage{
		db:  db,
		log: log,
		repos: repos{
			companyRepo: postgres.NewCompanyRepo(log, db),
			userRepo:    postgres.NewUserRepo(log, db),
			shopRepo:    postgres.NewShopRepo(db, log),
			productRepo: postgres.NewProductRepo(log, db),
		},
	}
}

func (s *storage) WithTransaction(ctx context.Context) (StorageTrI, error) {
	tr, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}

	return &storageTr{
		tr:    tr,
		repos: repos{},
	}, nil
}

func (s *storageTr) Commit(ctx context.Context) error {
	return s.tr.Commit(ctx)
}

func (s *storageTr) Rollback(ctx context.Context) error {
	return s.tr.Rollback(ctx)
}

func (r *repos) Company() repo.CompanyI {
	return r.companyRepo
}

func (r *repos) User() repo.UserI {
	return r.userRepo
}

func (r *repos) Shop() repo.ShopI {
	return r.shopRepo
}

func (r *repos) Product() repo.ProductI {
	return r.productRepo
}
