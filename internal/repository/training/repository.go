package training

import (
	"github.com/kirillmc/platform_common/pkg/db"
	"github.com/kirillmc/trainings-server/internal/repository"
)

const (
	usersProgramsTable  = "users_programs"
	trainProgramsTable  = "train_programs"
	trainDaysTable      = "train_days"
	daysProgramsTable   = "days_programs"
	exercisesTable      = "exercises"
	exercisesDaysTable  = "exercises_days"
	setsTable           = "sets"
	setsExercisesTable  = "sets_exercises"
	statisticsTable     = "statistics"
	statisticsSetsTable = "statistics_sets"

	idColumm   = "id"
	userId     = "user_id"
	programId  = "program_id"
	exerciseId = "exercise_id"
	setId      = "set_id"

	description  = "description"
	programName  = "program_name"
	dayName      = "day_name"
	exerciseName = "exercise_name"
	isPublic     = "is_public"

	pictures = "pictures"
	quantity = "quantity"
	weight   = "weight"
	time     = "time"

	returnId = "RETURNING id"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.TrainingRepository {
	return &repo{db: db}
}
