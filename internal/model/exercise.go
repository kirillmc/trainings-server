package model

import "google.golang.org/protobuf/types/known/wrapperspb"

type Exercise struct {
	Id           int64
	ExerciseName string
	Description  string
	DayId        int64
}

type ExerciseToCreate struct {
	ExerciseName string
	Description  string
	DayId        int64
}

type ExerciseToUpdate struct {
	Id           int64
	ExerciseName *wrapperspb.StringValue
	Description  *wrapperspb.StringValue
}
