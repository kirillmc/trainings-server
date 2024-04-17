package training

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/kirillmc/platform_common/pkg/db"
	"github.com/kirillmc/trainings-server/internal/model"
)

func (r *repo) UpdateProgram(ctx context.Context, req *model.TrainProgramToUpdate) error {
	builder := sq.Update(trainProgramsTable).
		PlaceholderFormat(sq.Dollar).Where(sq.Eq{idColumm: req.Id})

	if req.ProgramName != nil {
		builder = builder.Set(programName, req.ProgramName.Value)
	}

	if req.Description != nil {
		builder = builder.Set(description, req.Description.Value)
	}

	if req.IsPublic != nil {
		builder = builder.Set(isPublic, req.IsPublic.Value)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "training_repository.UpdateProgram",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) UpdateTrainDay(ctx context.Context, req *model.TrainDayToUpdate) error {
	builder := sq.Update(trainDaysTable).
		PlaceholderFormat(sq.Dollar).Where(sq.Eq{idColumm: req.Id})

	if req.DayName != nil {
		builder = builder.Set(dayName, req.DayName.Value)
	}

	if req.Description != nil {
		builder = builder.Set(description, req.Description.Value)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "training_repository.UpdateTrainDay",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) UpdateExercise(ctx context.Context, req *model.ExerciseToUpdate) error {
	builder := sq.Update(exercisesTable).
		PlaceholderFormat(sq.Dollar).Where(sq.Eq{idColumm: req.Id})

	if req.ExerciseName != nil {
		builder = builder.Set(exerciseName, req.ExerciseName.Value)
	}

	if req.Pictures != nil {
		builder = builder.Set(pictures, req.Pictures.Value)
	}

	if req.Description != nil {
		builder = builder.Set(description, req.Description.Value)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "training_repository.UpdateExercise",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) UpdateSet(ctx context.Context, req *model.SetToUpdate) error {
	builder := sq.Update(setsTable).
		PlaceholderFormat(sq.Dollar).Where(sq.Eq{idColumm: req.Id})

	if req.Quantity != nil {
		builder = builder.Set(quantity, req.Quantity.Value)
	}

	if req.Weight != nil {
		builder = builder.Set(weight, req.Weight.Value)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "training_repository.UpdateSet",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) UpdateStatistic(ctx context.Context, req *model.StatisticToUpdate) error {
	builder := sq.Update(statisticsTable).
		PlaceholderFormat(sq.Dollar).Where(sq.Eq{idColumm: req.Id})

	if req.Quantity != nil {
		builder = builder.Set(quantity, req.Quantity.Value)
	}

	if req.Weight != nil {
		builder = builder.Set(weight, req.Weight.Value)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "training_repository.UpdateStatistic",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
