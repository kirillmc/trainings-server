package converter

import (
	"github.com/kirillmc/trainings-server/internal/model"
	desc "github.com/kirillmc/trainings-server/pkg/training_v1"
)

func ToTrainingProgramFromDesc(trainingProgram *desc.CreateTrainingProgramRequest) *model.TrainProgramToCreate {
	return &model.TrainProgramToCreate{
		ProgramName: trainingProgram.Info.ProgramName,
		UserId:      trainingProgram.Info.UserId,
		Description: trainingProgram.Info.Description,
		IsPublic:    trainingProgram.Info.IsPublic,
	}
}

func ToGetTrainingProgramResponseFromService(trainingProgram *model.TrainProgram) *desc.TrainingProgram {
	return &desc.TrainingProgram{
		Id: trainingProgram.Id,
		Info: &desc.TrainingProgramInfo{
			ProgramName: trainingProgram.ProgramName,
			UserId:      trainingProgram.UserId,
			Description: trainingProgram.Description,
			IsPublic:    trainingProgram.IsPublic,
		},
	}
}

func ToGetTrainingProgramsResponseFromService(trainingPrograms []*model.TrainProgram) *desc.GetTrainingProgramsResponse {
	var trainingProgramsList []*desc.TrainingProgram
	for _, elem := range trainingPrograms {
		trainingProgramsList = append(trainingProgramsList, ToGetTrainingProgramResponseFromService(elem))
	}
	return &desc.GetTrainingProgramsResponse{TraingPrograms: trainingProgramsList}
}

func ToTrainingProgramUpdateFromDesc(trainingProgram *desc.UpdateTrainingProgramRequest) *model.TrainProgramToUpdate {
	return &model.TrainProgramToUpdate{
		Id:          trainingProgram.GetId(),
		ProgramName: trainingProgram.Info.GetProgramName(),
		Description: trainingProgram.Info.GetDescription(),
		IsPublic:    trainingProgram.Info.GetIsPublic(),
	}
}
