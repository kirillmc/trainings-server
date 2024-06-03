package trainer

import (
	"github.com/kirillmc/trainings-server/internal/repository"
	def "github.com/kirillmc/trainings-server/internal/service"
)

var _ def.TrainerService = (*serv)(nil)

type serv struct {
	trainerRepository repository.TrainerRepository
}

func NewService(trainerRepository repository.TrainerRepository) *serv {
	return &serv{trainerRepository: trainerRepository}
}
