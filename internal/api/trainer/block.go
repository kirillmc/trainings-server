package trainer

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/kirillmc/trainings-server/internal/converter"
	desc "github.com/kirillmc/trainings-server/pkg/trainer_v1"
)

func (i *Implementation) BlockTrainer(ctx context.Context, req *desc.TrainerClientRequest) (*emptypb.Empty, error) {
	err := i.trainerService.BlockTrainer(ctx, converter.ToTrainerClientFromDesc(req))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (i *Implementation) BlockClient(ctx context.Context, req *desc.TrainerClientRequest) (*emptypb.Empty, error) {
	err := i.trainerService.BlockClient(ctx, converter.ToTrainerClientFromDesc(req))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
