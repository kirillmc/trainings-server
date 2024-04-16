package training

import (
	"context"

	"github.com/kirillmc/trainings-server/internal/model"
)

func (r *repo) GetProgram(ctx context.Context, id int64) (*model.TrainProgram, error) {
	//err := i.userService.Update(ctx, converter.ToUserModelUpdateFromDesc(req))
	//if err != nil {
	//	return nil, err
	//}

	return &model.TrainProgram{}, nil
}

func (r *repo) GetTrainDay(ctx context.Context, id int64) (*model.TrainDay, error) {
	//err := i.userService.Update(ctx, converter.ToUserModelUpdateFromDesc(req))
	//if err != nil {
	//	return nil, err
	//}

	return &model.TrainDay{}, nil
}

func (r *repo) GetExercise(ctx context.Context, id int64) (*model.Exercise, error) {
	//err := i.userService.Update(ctx, converter.ToUserModelUpdateFromDesc(req))
	//if err != nil {
	//	return nil, err
	//}

	return &model.Exercise{}, nil

}

func (r *repo) GetSet(ctx context.Context, id int64) (*model.Set, error) {
	//err := i.userService.Update(ctx, converter.ToUserModelUpdateFromDesc(req))
	//if err != nil {
	//	return nil, err
	//}

	return &model.Set{}, nil
}

func (r *repo) GetStatistic(ctx context.Context, id int64) (*model.Statistic, error) {
	//err := i.userService.Update(ctx, converter.ToUserModelUpdateFromDesc(req))
	//if err != nil {
	//	return nil, err
	//}

	return &model.Statistic{}, nil
}
