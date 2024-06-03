package training

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/kirillmc/platform_common/pkg/db"
	"github.com/kirillmc/trainings-server/internal/model"
	"github.com/kirillmc/trainings-server/internal/repository/training/converter"
	modelRepo "github.com/kirillmc/trainings-server/internal/repository/training/model"
)

func (r *repo) GetDot(ctx context.Context, dotId int64) (*model.DotsToGet, error) {
	builder := sq.Select(idColumm, xCoordinate, yCoordinate, numberColumn).
		PlaceholderFormat(sq.Dollar).
		From(dotsTable).
		Where(sq.Eq{idColumm: dotId}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "training_repository.GetDot",
		QueryRaw: query,
	}

	var dot modelRepo.Dots

	err = r.db.DB().ScanOneContext(ctx, &dot, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToDotFromRepo(&dot), nil
}

func (r *repo) GetDotsIds(ctx context.Context, positionId int64) ([]int64, error) {
	builder := sq.Select(dotIdColumn).
		PlaceholderFormat(sq.Dollar).
		From(positionsDotsTable).
		Where(sq.Eq{positionIdColumn: positionId})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "training_repository.GetDotsIds",
		QueryRaw: query,
	}

	var ids []int64

	err = r.db.DB().ScanAllContext(ctx, &ids, q, args...)
	if err != nil {
		return nil, err
	}

	return ids, nil
}

func (r *repo) GetNumber(ctx context.Context, positionId int64) (int64, error) {
	builder := sq.Select(numberColumn).
		PlaceholderFormat(sq.Dollar).
		From(positionsTable).
		Where(sq.Eq{idColumm: positionId}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "training_repository.GetNumber",
		QueryRaw: query,
	}

	var id int64

	err = r.db.DB().ScanOneContext(ctx, &id, q, args...)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) GetPositionsIds(ctx context.Context, technicId int64) ([]int64, error) {
	builder := sq.Select(positionIdColumn).
		PlaceholderFormat(sq.Dollar).
		From(technicsPositionsTable).
		Where(sq.Eq{technicIdColumn: technicId})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "training_repository.GetPositionsIds",
		QueryRaw: query,
	}

	var ids []int64

	err = r.db.DB().ScanAllContext(ctx, &ids, q, args...)
	if err != nil {
		return nil, err
	}

	return ids, nil
}

func (r *repo) GetTechnicId(ctx context.Context, exerciseId int64) (int64, error) {
	builder := sq.Select(technicIdColumn).
		PlaceholderFormat(sq.Dollar).
		From(exercisesTechnicsTable).
		Where(sq.Eq{exerciseIdColumn: exerciseId}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "training_repository.GetTechnicId",
		QueryRaw: query,
	}

	var id int64

	err = r.db.DB().ScanOneContext(ctx, &id, q, args...)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) SetDot(ctx context.Context, req *model.Dot, positionId int64) error {
	builder := sq.Insert(dotsTable).PlaceholderFormat(sq.Dollar).
		Columns(xCoordinate, yCoordinate, numberColumn).
		Values(req.XCoordinate, req.YCoordinate, req.Number).
		Suffix(returnId)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "training_repository.SetDot",
		QueryRaw: query,
	}

	var dotId int64

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&dotId)
	if err != nil {
		return err
	}

	builder = sq.Insert(positionsDotsTable).PlaceholderFormat(sq.Dollar).
		Columns(positionIdColumn, dotIdColumn).
		Values(positionId, dotId)

	query, args, err = builder.ToSql()
	if err != nil {
		return err
	}

	q = db.Query{
		Name:     "training_repository.SetDot.positionsDotsTable",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) SetPosition(ctx context.Context, number int64, technicId int64) (int64, error) {
	builder := sq.Insert(positionsTable).PlaceholderFormat(sq.Dollar).
		Columns(numberColumn).
		Values(number).
		Suffix(returnId)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "training_repository.SetPosition",
		QueryRaw: query,
	}

	var positionId int64

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&positionId)
	if err != nil {
		return 0, err
	}

	builder = sq.Insert(technicsPositionsTable).PlaceholderFormat(sq.Dollar).
		Columns(technicIdColumn, positionIdColumn).
		Values(technicId, positionId)

	query, args, err = builder.ToSql()
	if err != nil {
		return 0, err
	}

	q = db.Query{
		Name:     "training_repository.SetDot.technicsPositionsTable",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return 0, err
	}
	return positionId, nil
}

func (r *repo) SetTechnic(ctx context.Context, exerciseId int64) (int64, error) {
	builder := sq.Insert(exercisesTechnicsTable).PlaceholderFormat(sq.Dollar).
		Columns(exerciseIdColumn).
		Values(exerciseId).
		Suffix(returnTechnicId)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "training_repository.SetTechnic",
		QueryRaw: query,
	}

	var technicId int64

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&technicId)
	if err != nil {
		return 0, err
	}

	return technicId, nil
}
