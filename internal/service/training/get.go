package training

import (
	"context"

	"github.com/kirillmc/trainings-server/internal/model"
)

func (s *serv) GetProgram(ctx context.Context, id int64) (*model.TrainProgram, error) {
	program, err := s.trainingRepository.GetProgram(ctx, id)
	if err != nil {
		return nil, err
	}

	return program, nil
}

func (s *serv) GetTrainDay(ctx context.Context, id int64) (*model.TrainDay, error) {
	trainDay, err := s.trainingRepository.GetTrainDay(ctx, id)
	if err != nil {
		return nil, err
	}

	return trainDay, nil
}

func (s *serv) GetExercise(ctx context.Context, id int64) (*model.Exercise, error) {
	exercise, err := s.trainingRepository.GetExercise(ctx, id)
	if err != nil {
		return nil, err
	}

	return exercise, nil
}

func (s *serv) GetSet(ctx context.Context, id int64) (*model.Set, error) {
	set, err := s.trainingRepository.GetSet(ctx, id)
	if err != nil {
		return nil, err
	}

	return set, nil
}

func (s *serv) GetStatistic(ctx context.Context, id int64) (*model.Statistic, error) {
	statistic, err := s.trainingRepository.GetStatistic(ctx, id)
	if err != nil {
		return nil, err
	}

	return statistic, nil
}
