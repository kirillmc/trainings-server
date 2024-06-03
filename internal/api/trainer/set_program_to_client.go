package trainer

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/kirillmc/trainings-server/internal/converter"
	desc "github.com/kirillmc/trainings-server/pkg/trainer_v1"
)

func (i *Implementation) SetProgramToClient(ctx context.Context, req *desc.TrainerClientRequest) (*empty.Empty, error) {
	err := i.trainerService.SetProgramToClient(ctx, converter.ToTrainerClientFromDesc(req))
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
