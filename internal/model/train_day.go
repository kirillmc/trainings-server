package model

import "google.golang.org/protobuf/types/known/wrapperspb"

type TrainDay struct {
	Id          int64
	DayName     string
	Description string
	ProgramId   int64
}

type TrainDayToCreate struct {
	DayName     string
	Description string
	ProgramId   int64
}

type TrainDayToUpdate struct {
	Id          int64
	DayName     *wrapperspb.StringValue
	Description *wrapperspb.StringValue
}
