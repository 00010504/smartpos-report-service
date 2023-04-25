package postgres

import (
	"context"
	"genproto/common"

	"github.com/Invan2/invan_report_service/models"
	"github.com/Invan2/invan_report_service/pkg/logger"
	"github.com/Invan2/invan_report_service/storage/repo"
	"github.com/pkg/errors"
)

type userRepo struct {
	log logger.Logger
	db  models.DB
}

func NewUserRepo(log logger.Logger, db models.DB) repo.UserI {
	return &userRepo{
		log: log,
		db:  db,
	}
}

func (u *userRepo) Upsert(ctx context.Context, entity *common.UserCreatedModel) error {

	query := `
		INSERT INTO "user" (
			"id",
			"first_name",
			"last_name",
			"phone_number",
			"user_type_id"
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5
		) ON CONFLICT ("id") 
		DO 
			UPDATE SET
				"first_name" = $2,
				"last_name" = $3,
				"phone_number" = $4,
				"user_type_id" = $5
	`

	if _, err := u.db.Exec(ctx, query, entity.Id, entity.FirstName, entity.LastName, entity.PhoneNumber, entity.UserTypeId); err != nil {
		return errors.Wrap(err, "error while upsert user")
	}

	return nil
}

func (u *userRepo) Delete(ctx context.Context, entity *common.RequestID) error {
	query := `
		UPDATE "user" SET "deleted_at"=extract(epoch from now())::bigint WHERE "id" =  $1
	`

	_, err := u.db.Exec(ctx, query, entity.Id)
	if err != nil {
		return errors.Wrap(err, "error while delete user")
	}

	return nil
}
