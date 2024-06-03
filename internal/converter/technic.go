package converter

import (
	"github.com/kirillmc/trainings-server/internal/model"
	desc "github.com/kirillmc/trainings-server/pkg/training_v1"
)

func ToTechnicsFromDesc(technics []*desc.Technic) []*model.Technic {
	var modelTechnics []*model.Technic

	for _, technic := range technics {
		modelTechnics = append(modelTechnics, &model.Technic{
			ExerciseId: technic.ExerciseId,
			Positions:  ToPositionsFromDesc(technic.Positions),
		})
	}

	return modelTechnics
}

func ToPositionsFromDesc(positions []*desc.Position) []model.Position {
	var positionsModel []model.Position

	for _, position := range positions {
		positionsModel = append(positionsModel, model.Position{
			Number: position.PositionNumber,
			Dots:   ToDotsFromDesc(position.Dots),
		})
	}

	return positionsModel
}

func ToDotsFromDesc(dots []*desc.Dots) []model.Dot {
	var dotModel []model.Dot
	for _, dot := range dots {
		dotModel = append(dotModel, model.Dot{
			Number:      dot.DotNumber,
			XCoordinate: dot.XCoordinate,
			YCoordinate: dot.YCoordinate,
		})
	}

	return dotModel
}

func ToDescFromTechnic(technic *model.TechnicToGet) *desc.GetTechnicResponse {
	return &desc.GetTechnicResponse{
		TechnicId:  technic.TechnicId,
		ExerciseId: technic.ExerciseId,
		Positions:  ToDescFromPositions(technic.PositionsSet),
	}
}

func ToDescFromPositions(positions []*model.PositionsToGet) []*desc.Position {
	var positionsDesc []*desc.Position
	for _, position := range positions {
		positionsDesc = append(positionsDesc, &desc.Position{
			PositionNumber: position.Number,
			Dots:           ToDescFromDots(position.DotsSet),
		})
	}

	return positionsDesc
}

func ToDescFromDots(dots []*model.DotsToGet) []*desc.Dots {
	var dotDesc []*desc.Dots
	for _, dot := range dots {
		dotDesc = append(dotDesc, &desc.Dots{
			DotNumber:   dot.Number,
			XCoordinate: dot.XCoordinate,
			YCoordinate: dot.YCoordinate,
		})
	}

	return dotDesc
}
