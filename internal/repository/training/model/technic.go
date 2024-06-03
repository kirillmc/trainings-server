package model

type Technic struct {
	TechnicId  int64 `db:"technic_id"`
	ExerciseId int64 `db:"exercise_id"`
}

type Positions struct {
	Id     int64 `db:"id"`
	Number int64 `db:"number"`
}

type Dots struct {
	Id          int64   `db:"id"`
	Number      int64   `db:"number"`
	XCoordinate float64 `db:"x_coordinate"`
	YCoordinate float64 `db:"y_coordinate"`
}
