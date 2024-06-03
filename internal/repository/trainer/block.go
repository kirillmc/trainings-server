package trainer

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/kirillmc/platform_common/pkg/db"
	"github.com/kirillmc/trainings-server/internal/model"
)

func (r *repo) BlockTrainer(ctx context.Context, req *model.TrainerClient) error {
	builder := sq.Update(usersTrainersProgramsTable).Set(isUserAgreeColumn, false).
		PlaceholderFormat(sq.Dollar).Where(sq.Eq{userIdColumn: req.ClientId, trainerIdColumn: req.TrainerId, programIdColumn: req.ProgramId})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "moder_repository.BlockTrainer",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) BlockClient(ctx context.Context, req *model.TrainerClient) error {
	builder := sq.Update(usersTrainersProgramsTable).Set(isTrainerAgreeColumn, false).
		PlaceholderFormat(sq.Dollar).Where(sq.Eq{userIdColumn: req.ClientId, trainerIdColumn: req.TrainerId, programIdColumn: req.ProgramId})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "moder_repository.BlockClient",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
