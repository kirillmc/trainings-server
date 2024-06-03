package trainer

import (
	"context"

	"github.com/kirillmc/trainings-server/internal/model"
)

func (s *serv) BlockTrainer(ctx context.Context, trainerToBlock *model.TrainerClient) error {
	err := s.trainerRepository.BlockTrainer(ctx, trainerToBlock)
	if err != nil {
		return err
	}

	return nil
}

func (s *serv) BlockClient(ctx context.Context, clientToBlock *model.TrainerClient) error {
	err := s.trainerRepository.BlockClient(ctx, clientToBlock)
	if err != nil {
		return err
	}

	return nil
}
