package postgres

import (
	"context"
	"genproto/common"

	"github.com/Invan2/invan_report_service/models"
	"github.com/Invan2/invan_report_service/pkg/logger"
	"github.com/Invan2/invan_report_service/storage/repo"
	"github.com/pkg/errors"
)

type productRepo struct {
	db  models.DB
	log logger.Logger
}

func NewProductRepo(log logger.Logger, db models.DB) repo.ProductI {
	return &productRepo{
		db:  db,
		log: log,
	}
}

func (p *productRepo) Upsert(ctx context.Context, entity *common.CreateProductCopyRequest) error {

	query := `
		INSERT INTO
			"product"
		(
			id,
			name,
			is_marking,
			sku,
			mxik_code,
			company_id
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (id) DO
		UPDATE
		SET name = $2;
	`

	_, err := p.db.Exec(
		ctx,
		query,
		entity.Id,
		entity.Name,
		entity.IsMarking,
		entity.Sku,
		entity.MxikCode,
		entity.Request.CompanyId,
	)
	if err != nil {
		return errors.Wrap(err, "error while insert product")
	}

	return nil
}

func (p *productRepo) Delete(ctx context.Context, req *common.RequestID) (*common.ResponseID, error) {
	return nil, nil
}
