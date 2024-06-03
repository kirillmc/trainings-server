package training

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/kirillmc/trainings-server/internal/converter"
	desc "github.com/kirillmc/trainings-server/pkg/training_v1"
)

func (i *Implementation) SetTechnic(ctx context.Context, req *desc.SetTechnicRequest) (*emptypb.Empty, error) {
	err := i.trainingService.SetTechnic(ctx, converter.ToTechnicsFromDesc(req.GetTechnics()))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (i *Implementation) GetTechnic(ctx context.Context, req *desc.GetTechnicRequest) (*desc.GetTechnicResponse, error) {
	technic, err := i.trainingService.GetTechnic(ctx, req.GetExerciseId())
	if err != nil {
		return nil, err
	}

	return converter.ToDescFromTechnic(technic), nil

}
