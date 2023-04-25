package postgres

import (
	"context"
	"genproto/common"

	"github.com/Invan2/invan_report_service/models"
	"github.com/Invan2/invan_report_service/pkg/logger"
	"github.com/Invan2/invan_report_service/storage/repo"
	"github.com/pkg/errors"
)

type companyRepo struct {
	db  models.DB
	log logger.Logger
}

func NewCompanyRepo(log logger.Logger, db models.DB) repo.CompanyI {
	return &companyRepo{
		db:  db,
		log: log,
	}
}

func (c *companyRepo) Upsert(ctx context.Context, entity *common.CompanyCreatedModel) error {

	query := `
		INSERT INTO
			"company"
		(
			id,
			name
		)
		VALUES (
			$1,
			$2
		) ON CONFLICT (id) DO
		UPDATE
			SET
			name = $2;
	`

	_, err := c.db.Exec(
		ctx,
		query,
		entity.Id,
		entity.Name,
	)
	if err != nil {
		return errors.Wrap(err, "error while insert company")
	}

	return nil
}

func (c *companyRepo) Delete(ctx context.Context, req *common.RequestID) error {

	query := `
	  	UPDATE
			"company"
	  	SET
			deleted_at = extract(epoch from now())::bigint
	  	WHERE
			id = $1 AND deleted_at = 0
	`

	_, err := c.db.Exec(
		ctx,
		query,
		req.Id,
	)
	if err != nil {
		return errors.Wrap(err, "error while delete company")
	}

	return nil
}
