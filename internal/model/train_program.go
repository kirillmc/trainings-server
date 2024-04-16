package model

import "google.golang.org/protobuf/types/known/wrapperspb"

type TrainProgram struct {
	Id          int64
	ProgramName string
	UserId      int64
	Description string
	IsPublic    bool
}

type TrainProgramToCreate struct {
	ProgramName string
	UserId      int64
	Description string
	IsPublic    bool
}

type TrainProgramToUpdate struct {
	Id          int64
	ProgramName *wrapperspb.StringValue
	Description *wrapperspb.StringValue
	IsPublic    *wrapperspb.BoolValue
}
