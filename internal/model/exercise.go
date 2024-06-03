package model

import (
	"github.com/kirillmc/platform_common/pkg/nillable"
)

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
	ExerciseName nillable.NilString
	Description  nillable.NilString
}
