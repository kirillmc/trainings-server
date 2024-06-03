package converter

import (
	"github.com/kirillmc/trainings-server/internal/model"
	desc "github.com/kirillmc/trainings-server/pkg/trainer_v1"
)

func ToTrainerClientFromDesc(blockTrainerRequest *desc.TrainerClientRequest) *model.TrainerClient {

	return &model.TrainerClient{
		TrainerId: blockTrainerRequest.GetTrainerId(),
		ClientId:  blockTrainerRequest.GetClientId(),
		ProgramId: blockTrainerRequest.GetProgramId(),
	}
}

func ToClientDescFromClientsTrainingProgram(clientsTrainingProgram []*model.ClientsTrainingProgram) *desc.GetClientsProgramsResponse {
	clientsPrograms := []*desc.ClientsTrainingProgram{}
	for i, _ := range clientsTrainingProgram {
		clientsProgram := &desc.ClientsTrainingProgram{
			ProgramId:   clientsTrainingProgram[i].ProgramId,
			ClientId:    clientsTrainingProgram[i].ClientId,
			ProgramName: clientsTrainingProgram[i].ProgramName,
			Description: clientsTrainingProgram[i].Description,
			Status:      desc.Status(clientsTrainingProgram[i].Status),
		}

		clientsPrograms = append(clientsPrograms, clientsProgram)
	}

	return &desc.GetClientsProgramsResponse{ClientsPrograms: clientsPrograms}
}

func ToTrainerDescFromClientsTrainingProgram(trainersTrainingProgram []*model.TrainersTrainingProgram) *desc.GetProgramsWithTrainersResponse {
	trainersPrograms := []*desc.TrainersTrainingProgram{}
	for i, _ := range trainersTrainingProgram {
		trainersProgram := &desc.TrainersTrainingProgram{
			ProgramId:   trainersTrainingProgram[i].ProgramId,
			TrainerId:   trainersTrainingProgram[i].TrainerId,
			ProgramName: trainersTrainingProgram[i].ProgramName,
			Description: trainersTrainingProgram[i].Description,
			Status:      desc.Status(trainersTrainingProgram[i].Status),
		}

		trainersPrograms = append(trainersPrograms, trainersProgram)
	}

	return &desc.GetProgramsWithTrainersResponse{TrainersPrograms: trainersPrograms}
}

func ToGetClientsToAllowDescFromModel(clients []int64) *desc.GetClientsToAllowResponse {

	return &desc.GetClientsToAllowResponse{
		ClientId: clients,
	}
}

func ToGetTrainersToAllowDescFromModel(trainers []int64) *desc.GetTrainersToAllowResponse {

	return &desc.GetTrainersToAllowResponse{
		TrainerId: trainers,
	}
}
