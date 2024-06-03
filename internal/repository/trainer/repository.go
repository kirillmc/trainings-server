package trainer

import (
	"github.com/kirillmc/platform_common/pkg/db"
	"github.com/kirillmc/trainings-server/internal/repository"
)

const (
	usersProgramsTable         = "users_programs"
	trainersProgramsTable      = "trainers_programs"
	usersTrainersProgramsTable = "users_trainers_programs"
	trainProgramsTable         = "train_programs"
	trainDaysTable             = "train_days"
	daysProgramsTable          = "days_programs"
	exercisesTable             = "exercises"
	exercisesDaysTable         = "exercises_days"
	setsTable                  = "sets"
	setsExercisesTable         = "sets_exercises"
	exercisesPicturesTable     = "exercises_pictures"
	exercisesTechnicsTable     = "exercises_technics"
	technicsPositionsTable     = "technics_positions"
	positionsDotsTable         = "positions_dots"
	positionsTable             = "positions"
	dotsTable                  = "dots"
	picturesTable              = "pictures"
	statisticsTable            = "statistics"
	statisticsSetsTable        = "statistics_sets"

	idColumm             = "id"
	userIdColumn         = "user_id"
	trainerIdColumn      = "trainer_id"
	technicIdColumn      = "technic_id"
	positionIdColumn     = "position_id"
	picturesIdColumn     = "pictures_id"
	dotIdColumn          = "dot_id"
	programIdColumn      = "program_id"
	trainsDayIdColumn    = "trains_day_id"
	exerciseIdColumn     = "exercise_id"
	setIdColumn          = "set_id"
	statisticIdColumn    = "statistic_id"
	isUserAgreeColumn    = "is_user_agree"
	isTrainerAgreeColumn = "is_trainer_agree"

	xCoordinate  = "x_coordinate"
	yCoordinate  = "y_coordinate"
	numberColumn = "number"
	description  = "description"
	programName  = "program_name"
	dayName      = "day_name"
	exerciseName = "exercise_name"
	status       = "status"

	picture  = "picture"
	quantity = "quantity"
	weight   = "weight"
	time     = "time"

	publicStatusValue = 3

	returnId        = "RETURNING id"
	returnTechnicId = "RETURNING technic_id"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.TrainerRepository {
	return &repo{
		db: db,
	}
}
