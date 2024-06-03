package trainer

import (
	"context"

	"github.com/kirillmc/trainings-server/internal/converter"
	desc "github.com/kirillmc/trainings-server/pkg/trainer_v1"
)

func (i *Implementation) GetClientsPrograms(ctx context.Context, req *desc.GetClientsProgramsRequest) (*desc.GetClientsProgramsResponse, error) {
	clientsTrainingProgram, err := i.trainerService.GetClientsPrograms(ctx, req.GetTrainerId())
	if err != nil {
		return nil, err
	}

	return converter.ToClientDescFromClientsTrainingProgram(clientsTrainingProgram), nil
}

func (i *Implementation) GetProgramsWithTrainers(ctx context.Context, req *desc.GetProgramsWithTrainersRequest) (*desc.GetProgramsWithTrainersResponse, error) {
	trainerTrainingProgram, err := i.trainerService.GetProgramsWithTrainers(ctx, req.GetClientId())
	if err != nil {
		return nil, err
	}

	return converter.ToTrainerDescFromClientsTrainingProgram(trainerTrainingProgram), nil
}

func (i *Implementation) GetClientsToAllow(ctx context.Context, req *desc.GetClientsToAllowRequest) (*desc.GetClientsToAllowResponse, error) {
	clients, err := i.trainerService.GetClientsToAllow(ctx, req.GetTrainerId())
	if err != nil {
		return nil, err
	}

	return converter.ToGetClientsToAllowDescFromModel(clients), nil
}

func (i *Implementation) GetTrainersToAllow(ctx context.Context, req *desc.GetTrainersToAllowRequest) (*desc.GetTrainersToAllowResponse, error) {
	trainers, err := i.trainerService.GetTrainersToAllow(ctx, req.GetClientId())
	if err != nil {
		return nil, err
	}

	return converter.ToGetTrainersToAllowDescFromModel(trainers), nil
}
