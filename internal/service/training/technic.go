package training

import (
	"context"
	"log"

	"github.com/kirillmc/trainings-server/internal/model"
)

func (s *serv) SetTechnic(ctx context.Context, technics []*model.Technic) error {
	for _, technic := range technics {
		technicId, err := s.trainingRepository.SetTechnic(ctx, technic.ExerciseId)
		if err != nil {
			return err
		}

		for _, position := range technic.Positions {
			positionId, err := s.trainingRepository.SetPosition(ctx, position.Number, technicId)
			if err != nil {
				return err
			}
			for _, dot := range position.Dots {
				err = s.trainingRepository.SetDot(ctx, &dot, positionId)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
func (s *serv) GetTechnic(ctx context.Context, exerciseId int64) (*model.TechnicToGet, error) {
	technicId, err := s.trainingRepository.GetTechnicId(ctx, exerciseId)
	if err != nil {
		return nil, err
	}
	var positions []*model.PositionsToGet
	positionsIds, err := s.trainingRepository.GetPositionsIds(ctx, technicId)
	for positionId := range positionsIds {
		dotsIds, err := s.trainingRepository.GetDotsIds(ctx, int64(positionId))
		if err != nil {
			return nil, err
		}
		log.Printf("\nIDS:%v\n", dotsIds)
		var dots []*model.DotsToGet
		for _, dotId := range dotsIds {
			log.Printf("\ndotsIds:%v\n", dotsIds)
			log.Printf("\nID:%v\n", dotId)

			dot, err := s.trainingRepository.GetDot(ctx, int64(dotId))
			if err != nil {
				return nil, err
			}

			dots = append(dots, dot)
		}

		number, err := s.trainingRepository.GetNumber(ctx, int64(positionId))
		positions = append(positions, &model.PositionsToGet{Id: int64(positionId), Number: number, DotsSet: dots})
	}

	return &model.TechnicToGet{
		TechnicId:    technicId,
		ExerciseId:   exerciseId,
		PositionsSet: positions,
	}, nil
}

type a interface {
	SetDot(ctx context.Context, req *model.Dot, positionId int64) error
	SetPosition(ctx context.Context, number int64, technicId int64) error
	SetTechnic(ctx context.Context, exerciseId int64) (int64, error)
}
