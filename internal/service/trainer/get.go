package trainer

import (
	"context"

	"github.com/kirillmc/trainings-server/internal/model"
)

func (s *serv) GetClientsPrograms(ctx context.Context, trainerId int64) ([]*model.ClientsTrainingProgram, error) {

	return nil, nil
}

func (s *serv) GetProgramsWithTrainers(ctx context.Context, clientId int64) ([]*model.TrainersTrainingProgram, error) {

	return nil, nil
}

func (s *serv) GetClientsToAllow(ctx context.Context, trainerId int64) ([]int64, error) {

	return nil, nil
}

func (s *serv) GetTrainersToAllow(ctx context.Context, clientId int64) ([]int64, error) {

	return nil, nil
}
