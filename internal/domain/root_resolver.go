package domain

import (
	"context"
	"github.com/FitBuddy/internal/domain/coach"
	"github.com/FitBuddy/internal/domain/exercise"
	"github.com/FitBuddy/internal/domain/forum"
	"github.com/FitBuddy/internal/domain/goal"
	"github.com/FitBuddy/internal/domain/leaderboard"
	"github.com/FitBuddy/internal/domain/nutrition"
	"github.com/FitBuddy/internal/domain/sleep"
	"github.com/FitBuddy/internal/domain/user"
	"github.com/FitBuddy/internal/domain/workout"
	"github.com/FitBuddy/pkg/graphql"
	"time"
)

var _ graphql.ResolverRoot = &RootResolver{}

type RootResolver struct {
	userResolver        *user.Resolver
	coachResolver       *coach.Resolver
	exerciseResolver    *exercise.Resolver
	forumResolver       *forum.Resolver
	goalResolver        *goal.Resolver
	leaderboardResolver *leaderboard.Resolver
	nutritionResolver   *nutrition.Resolver
	sleepResolver       *sleep.Resolver
	workoutResolver     *workout.Resolver
}

func NewRootResolver(userResolver *user.Resolver, coachResolver *coach.Resolver, exerciseResolver *exercise.Resolver, forumResolver *forum.Resolver, goalResolver *goal.Resolver, leaderboardResolver *leaderboard.Resolver, nutritionResolver *nutrition.Resolver, sleepResolver *sleep.Resolver, workoutResolver *workout.Resolver) *RootResolver {
	return &RootResolver{
		userResolver:        userResolver,
		coachResolver:       coachResolver,
		exerciseResolver:    exerciseResolver,
		forumResolver:       forumResolver,
		goalResolver:        goalResolver,
		leaderboardResolver: leaderboardResolver,
		nutritionResolver:   nutritionResolver,
		sleepResolver:       sleepResolver,
		workoutResolver:     workoutResolver,
	}
}

type mutationResolver struct {
	*RootResolver
}

func (m mutationResolver) CreateGoal(ctx context.Context, name string, description string, startDate string, endDate string, email string) (*graphql.Goal, error) {
	parsedStartDate, err := time.Parse(time.RFC3339, startDate)
	if err != nil {
		return nil, err
	}
	parsedEndDate, err := time.Parse(time.RFC3339, endDate)
	if err != nil {
		return nil, err
	}

	return m.goalResolver.CreateGoal(ctx, email, name, description, parsedStartDate, parsedEndDate)
}

func (m mutationResolver) DeleteGoal(ctx context.Context, goalID string) (bool, error) {
	return m.goalResolver.DeleteGoal(ctx, goalID)
}

func (m mutationResolver) CreateSleepLog(ctx context.Context, userEmail, sleepTime, wakeTime, date string) (*graphql.SleepLog, error) {
	parsedSleepTime, err := time.Parse(time.RFC3339, sleepTime)
	if err != nil {
		return nil, err
	}
	parsedWakeTime, err := time.Parse(time.RFC3339, wakeTime)
	if err != nil {
		return nil, err
	}
	parsedDateTime, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return nil, err
	}

	return m.sleepResolver.CreateSleepLog(ctx, userEmail, parsedSleepTime, parsedWakeTime, parsedDateTime)
}

func (m mutationResolver) DeleteSleepLog(ctx context.Context, sleepLogID string) (bool, error) {
	return m.sleepResolver.DeleteSleepLog(ctx, sleepLogID)
}

func (m mutationResolver) BookCoach(ctx context.Context, email string, coachName string) (bool, error) {
	return m.coachResolver.BookCoach(ctx, email, coachName)
}

func (m mutationResolver) UnbookCoach(ctx context.Context, email string, coachName string) (bool, error) {
	return m.coachResolver.UnbookCoach(ctx, email, coachName)

}

func (m mutationResolver) CreatePost(ctx context.Context, title, content, email string) (*graphql.Post, error) {
	return m.forumResolver.CreatePost(ctx, title, content, email)
}

func (m mutationResolver) DeletePost(ctx context.Context, postID string) (bool, error) {
	return m.forumResolver.DeletePost(ctx, postID)
}

func (m mutationResolver) LoginUser(ctx context.Context, email string, password string) (*graphql.User, error) {
	return m.userResolver.LoginUser(ctx, email, password)
}

func (m mutationResolver) LogoutUser(ctx context.Context, email string) (*graphql.User, error) {
	return m.userResolver.LogoutUser(ctx, email)
}

func (m mutationResolver) CreateUser(ctx context.Context, email string, password string) (*graphql.User, error) {
	return m.userResolver.CreateUser(ctx, email, password)
}

func (m mutationResolver) CreateCoach(ctx context.Context, name string, specialty string) (*graphql.Coach, error) {
	//TODO implement me
	panic("implement me")
}

func (m mutationResolver) CreateWorkout(ctx context.Context, email, exercise, date string, sets, reps int, weight float64) (*graphql.Workout, error) {
	return m.workoutResolver.CreateWorkout(ctx, email, exercise, date, sets, reps, weight)
}

func (m mutationResolver) CreateNutritionLog(ctx context.Context, userID string, description string, calories int) (*graphql.NutritionLog, error) {
	//TODO implement me
	panic("implement me")
}

// Mutation missing godoc
func (r *RootResolver) Mutation() graphql.MutationResolver {
	return &mutationResolver{r}
}

type queryResolver struct {
	*RootResolver
}

func (q queryResolver) GetSleepLogByEmailAndDate(ctx context.Context, userEmail string, date string) ([]*graphql.SleepLog, error) {
	parsedDateTime, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return nil, err
	}

	return q.sleepResolver.GetSleepLogByEmailAndDate(ctx, userEmail, parsedDateTime)
}

func (q queryResolver) IsCoachBookedByUser(ctx context.Context, coachName, userEmail string) (bool, error) {
	return q.coachResolver.IsCoachBookedByUser(ctx, coachName, userEmail)
}

func (q queryResolver) IsCoachBooked(ctx context.Context, coachName string) (bool, error) {
	return q.coachResolver.IsCoachBooked(ctx, coachName)
}

func (q queryResolver) GetAllCoaches(ctx context.Context) ([]*graphql.Coach, error) {
	return q.coachResolver.GetAllCoaches(ctx)
}

func (q queryResolver) GetAllPosts(ctx context.Context) ([]*graphql.Post, error) {
	return q.forumResolver.GetAllPosts(ctx)
}

func (q queryResolver) GetAllWorkoutsByEmailAndDate(ctx context.Context, email, date string) ([]*graphql.Workout, error) {
	return q.workoutResolver.GetAllWorkouts(ctx, email, date)
}

func (q queryResolver) GetAllExercises(ctx context.Context) ([]*graphql.Exercise, error) {
	return q.exerciseResolver.GetAllExercises(ctx)
}

func (q queryResolver) GetGoals(ctx context.Context, userEmail string) ([]*graphql.Goal, error) {
	return q.goalResolver.GetGoals(ctx, userEmail)
}

func (q queryResolver) GetUserByID(ctx context.Context, userID string) (*graphql.User, error) {
	//TODO implement me
	panic("implement me")
}

func (q queryResolver) GetCoachByID(ctx context.Context, coachID string) (*graphql.Coach, error) {
	//TODO implement me
	panic("implement me")
}

func (q queryResolver) GetNutritionLogByID(ctx context.Context, nutritionLogID string) (*graphql.NutritionLog, error) {
	//TODO implement me
	panic("implement me")
}

// Query missing godoc
func (r *RootResolver) Query() graphql.QueryResolver {
	return &queryResolver{r}
}
