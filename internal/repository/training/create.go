package training

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/kirillmc/platform_common/pkg/db"
	"github.com/kirillmc/trainings-server/internal/model"
)

func (r *repo) CreateProgram(ctx context.Context, req *model.TrainProgramToCreate) (int64, error) {
	builder := sq.Insert(trainProgramsTable).PlaceholderFormat(sq.Dollar).
		Columns(programName, description, isPublic).
		Values(req.ProgramName, req.Description, req.IsPublic).
		Suffix(returnId)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "training_repository.CreateProgram",
		QueryRaw: query,
	}

	var id int64

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil

}

func (r *repo) CreateTrainDay(ctx context.Context, req *model.TrainDayToCreate) (int64, error) {
	//err := i.userService.Update(ctx, converter.ToUserModelUpdateFromDesc(req))
	//if err != nil {
	//	return nil, err
	//}

	return 0, nil

}

func (r *repo) CreateExercise(ctx context.Context, req *model.ExerciseToCreate) (int64, error) {
	//err := i.userService.Update(ctx, converter.ToUserModelUpdateFromDesc(req))
	//if err != nil {
	//	return nil, err
	//}

	return 0, nil

}

func (r *repo) CreateSet(ctx context.Context, req *model.SetToCreate) (int64, error) {
	//err := i.userService.Update(ctx, converter.ToUserModelUpdateFromDesc(req))
	//if err != nil {
	//	return nil, err
	//}

	return 0, nil

}

func (r *repo) CreateStatistic(ctx context.Context, req *model.StatisticToCreate) (int64, error) {
	//err := i.userService.Update(ctx, converter.ToUserModelUpdateFromDesc(req))
	//if err != nil {
	//	return nil, err
	//}

	return 0, nil

}
