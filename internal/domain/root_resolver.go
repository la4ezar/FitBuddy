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

// CreateGoal creates a goal with provided arguments.
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

// CompleteGoal marks goal as completed
func (m mutationResolver) CompleteGoal(ctx context.Context, userEmail, goalID string) (bool, error) {
	return m.goalResolver.CompleteGoal(ctx, userEmail, goalID)
}

// DeleteGoal deletes a goal by its id.
func (m mutationResolver) DeleteGoal(ctx context.Context, goalID string) (bool, error) {
	return m.goalResolver.DeleteGoal(ctx, goalID)
}

// CreateSleepLog creates a sleep log with provided arguments.
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

// DeleteSleepLog deletes a sleep log by its id.
func (m mutationResolver) DeleteSleepLog(ctx context.Context, sleepLogID string) (bool, error) {
	return m.sleepResolver.DeleteSleepLog(ctx, sleepLogID)
}

// BookCoach books a coach
func (m mutationResolver) BookCoach(ctx context.Context, email string, coachName string) (bool, error) {
	return m.coachResolver.BookCoach(ctx, email, coachName)
}

// UnbookCoach unbooks a coach
func (m mutationResolver) UnbookCoach(ctx context.Context, email string, coachName string) (bool, error) {
	return m.coachResolver.UnbookCoach(ctx, email, coachName)

}

// CreatePost creates a post with provided arguments.
func (m mutationResolver) CreatePost(ctx context.Context, title, content, email string) (*graphql.Post, error) {
	return m.forumResolver.CreatePost(ctx, title, content, email)
}

// DeletePost deletes a post by its id.
func (m mutationResolver) DeletePost(ctx context.Context, postID string) (bool, error) {
	return m.forumResolver.DeletePost(ctx, postID)
}

// LoginUser logins an user
func (m mutationResolver) LoginUser(ctx context.Context, email string, password string) (*graphql.User, error) {
	return m.userResolver.LoginUser(ctx, email, password)
}

// LogoutUser logs out an user
func (m mutationResolver) LogoutUser(ctx context.Context, email string) (*graphql.User, error) {
	return m.userResolver.LogoutUser(ctx, email)
}

// CreateUser creates user with provided arguments.
func (m mutationResolver) CreateUser(ctx context.Context, email string, password string) (*graphql.User, error) {
	return m.userResolver.CreateUser(ctx, email, password)
}

// CreateWorkout creates workout with provided arguments.
func (m mutationResolver) CreateWorkout(ctx context.Context, email, exercise, date string, sets, reps int, weight float64) (*graphql.Workout, error) {
	return m.workoutResolver.CreateWorkout(ctx, email, exercise, date, sets, reps, weight)
}

// CreateNutrition creates nutrition with provided arguments.
func (m mutationResolver) CreateNutrition(ctx context.Context, email, meal, date string, servingSize, numberOfServings int) (*graphql.Nutrition, error) {
	return m.nutritionResolver.CreateNutrition(ctx, email, meal, date, servingSize, numberOfServings)
}

// Mutation missing godoc
func (r *RootResolver) Mutation() graphql.MutationResolver {
	return &mutationResolver{r}
}

type queryResolver struct {
	*RootResolver
}

// GetSleepLogByEmailAndDate get sleep log by user email and date
func (q queryResolver) GetSleepLogByEmailAndDate(ctx context.Context, userEmail string, date string) ([]*graphql.SleepLog, error) {
	parsedDateTime, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return nil, err
	}

	return q.sleepResolver.GetSleepLogByEmailAndDate(ctx, userEmail, parsedDateTime)
}

// IsCoachBookedByUser checks if a coach is booked by user
func (q queryResolver) IsCoachBookedByUser(ctx context.Context, coachName, userEmail string) (bool, error) {
	return q.coachResolver.IsCoachBookedByUser(ctx, coachName, userEmail)
}

// IsCoachBooked checks if a coach is booked
func (q queryResolver) IsCoachBooked(ctx context.Context, coachName string) (bool, error) {
	return q.coachResolver.IsCoachBooked(ctx, coachName)
}

// GetAllCoaches gets all coaches
func (q queryResolver) GetAllCoaches(ctx context.Context) ([]*graphql.Coach, error) {
	return q.coachResolver.GetAllCoaches(ctx)
}

// GetAllPosts gets all posts
func (q queryResolver) GetAllPosts(ctx context.Context) ([]*graphql.Post, error) {
	return q.forumResolver.GetAllPosts(ctx)
}

// GetAllWorkoutsByEmailAndDate gets all workouts by user email and date
func (q queryResolver) GetAllWorkoutsByEmailAndDate(ctx context.Context, email, date string) ([]*graphql.Workout, error) {
	return q.workoutResolver.GetAllWorkouts(ctx, email, date)
}

// GetAllExercises gets all exercises
func (q queryResolver) GetAllExercises(ctx context.Context) ([]*graphql.Exercise, error) {
	return q.exerciseResolver.GetAllExercises(ctx)
}

// GetAllNutritionsByEmailAndDate gets nutritions by user email and date
func (q queryResolver) GetAllNutritionsByEmailAndDate(ctx context.Context, email, date string) ([]*graphql.Nutrition, error) {
	return q.nutritionResolver.GetAllNutritions(ctx, email, date)
}

// GetAllMeals gets all meals
func (q queryResolver) GetAllMeals(ctx context.Context) ([]*graphql.Meal, error) {
	return q.nutritionResolver.GetAllMeals(ctx)
}

// GetGoals gets goals by user email
func (q queryResolver) GetGoals(ctx context.Context, userEmail string) ([]*graphql.Goal, error) {
	return q.goalResolver.GetGoals(ctx, userEmail)
}

// GetLeaderboardUsers gets leaderboard users
func (q queryResolver) GetLeaderboardUsers(ctx context.Context) ([]*graphql.LeaderboardUser, error) {
	return q.leaderboardResolver.GetLeaderboardUsers(ctx)
}

// Query missing godoc
func (r *RootResolver) Query() graphql.QueryResolver {
	return &queryResolver{r}
}
