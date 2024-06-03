package trainer

import (
	"github.com/kirillmc/trainings-server/internal/service"
	desc "github.com/kirillmc/trainings-server/pkg/trainer_v1"
)

type Implementation struct {
	desc.UnimplementedTrainerV1Server
	trainerService service.TrainerService
}

func NewImplementation(trainerService service.TrainerService) *Implementation {
	return &Implementation{
		trainerService: trainerService,
	}
}
