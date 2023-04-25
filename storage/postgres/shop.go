package postgres

import (
	"context"
	"genproto/common"

	"github.com/Invan2/invan_report_service/models"
	"github.com/Invan2/invan_report_service/pkg/logger"
	"github.com/Invan2/invan_report_service/storage/repo"
	"github.com/pkg/errors"
)

type shopRepo struct {
	db  models.DB
	log logger.Logger
}

func NewShopRepo(db models.DB, log logger.Logger) repo.ShopI {
	return &shopRepo{db: db, log: log}
}

func (sh *shopRepo) Delete(ctx context.Context) error {
	return nil
}

func (sh *shopRepo) Upsert(entity *common.ShopCreatedModel) error {
	query := `
		INSERT INTO
			"shop"
		(
			id,
			name,
			company_id,
			created_by
		)
		VALUES (
			$1,
			$2,
			$3,
			$4
		) ON CONFLICT (id) DO
		UPDATE
			SET
			"name" = $2,
			"company_id" = $3
			;
	`

	_, err := sh.db.Exec(
		context.Background(),
		query,
		entity.Id,
		entity.Name,
		entity.Request.CompanyId,
		entity.Request.UserId,
	)
	if err != nil {
		return errors.Wrap(err, "error while insert shop")
	}

	return nil
}
