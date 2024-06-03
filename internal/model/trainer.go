package model

type TrainerClient struct {
	TrainerId int64
	ClientId  int64
	ProgramId int64
}

type ClientsTrainingProgram struct {
	ProgramId   int64
	ClientId    int64
	ProgramName string
	Description string
	Status      Status
}

type TrainersTrainingProgram struct {
	ProgramId   int64
	TrainerId   int64
	ProgramName string
	Description string
	Status      Status
}
