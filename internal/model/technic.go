package model

type TechnicToGet struct {
	TechnicId    int64
	ExerciseId   int64
	PositionsSet []*PositionsToGet
}

type PositionsToGet struct {
	Id      int64
	Number  int64
	DotsSet []*DotsToGet
}

type DotsToGet struct {
	Number      int64
	XCoordinate float64
	YCoordinate float64
}

type Technic struct {
	ExerciseId int64
	Positions  []Position
}

type Position struct {
	Number int64
	Dots   []Dot
}

type Dot struct {
	Number      int64
	XCoordinate float64
	YCoordinate float64
}
