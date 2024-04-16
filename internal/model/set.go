package model

import "google.golang.org/protobuf/types/known/wrapperspb"

type Set struct {
	Id         int64
	Quantity   int64
	Weight     float64
	ExerciseId int64
}

type SetToCreate struct {
	Quantity   int64
	Weight     float64
	ExerciseId int64
}

type SetToUpdate struct {
	Id       int64
	Quantity *wrapperspb.Int64Value
	Weight   *wrapperspb.DoubleValue
}
