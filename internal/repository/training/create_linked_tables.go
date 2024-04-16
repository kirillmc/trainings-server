package training

import "context"

func (r *repo) CreateUsersPrograms(ctx context.Context, user_id, program_id int64) error {
	return nil
}
func (r *repo) CreateDaysPrograms(ctx context.Context, program_id, day_id int64) error {
	return nil
}
func (r *repo) CreateExercisesDays(ctx context.Context, day_id, exercise_id int64) error {
	return nil
}
func (r *repo) CreateSetsExercises(ctx context.Context, exercise_id, set_id int64) error {
	return nil
}
func (r *repo) CreateStatisticsSets(ctx context.Context, set_id, statistic_id int64) error {
	return nil
}
