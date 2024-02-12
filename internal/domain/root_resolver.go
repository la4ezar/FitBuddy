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
	"github.com/FitBuddy/pkg/log"
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

func (m mutationResolver) LoginUser(ctx context.Context, email string, password string) (*graphql.User, error) {
	log.C(ctx).Infof("Logging user with email %q...", email)
	u, err := m.userResolver.LoginUserMutation(ctx, email, password)
	if err != nil {
		return nil, err
	}
	log.C(ctx).Infof("Successfully logged user with email %q...", email)

	gqlUser := &graphql.User{
		ID:    u.ID,
		Email: u.Email,
	}
	return gqlUser, nil
}

func (m mutationResolver) LogoutUser(ctx context.Context, email string) (*graphql.User, error) {
	log.C(ctx).Infof("Logging out user with email %q...", email)
	u, err := m.userResolver.LogoutUserMutation(ctx, email)
	if err != nil {
		return nil, err
	}
	log.C(ctx).Infof("Successfully logging out user with email %q...", email)

	gqlUser := &graphql.User{
		ID:    u.ID,
		Email: u.Email,
	}
	return gqlUser, nil
}

func (m mutationResolver) CreateUser(ctx context.Context, email string, password string) (*graphql.User, error) {
	log.C(ctx).Infof("Creating User with email %q...", email)
	input := user.CreateUserInput{
		Email:    email,
		Password: password,
		Logged:   false,
	}

	u, err := m.userResolver.CreateUserMutation(ctx, input)
	if err != nil {
		return nil, err
	}
	log.C(ctx).Info("Successfully created user with email %q", email)

	gqlUser := &graphql.User{
		ID:    u.ID,
		Email: u.Email,
	}
	return gqlUser, nil
}

func (m mutationResolver) CreateCoach(ctx context.Context, name string, specialty string) (*graphql.Coach, error) {
	//TODO implement me
	panic("implement me")
}

func (m mutationResolver) CreateWorkoutLog(ctx context.Context, userID string, exercise string, sets int, reps int, weight float64) (*graphql.WorkoutLog, error) {
	//TODO implement me
	panic("implement me")
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

func (q queryResolver) GetUserByID(ctx context.Context, userID string) (*graphql.User, error) {
	//TODO implement me
	panic("implement me")
}

func (q queryResolver) GetCoachByID(ctx context.Context, coachID string) (*graphql.Coach, error) {
	//TODO implement me
	panic("implement me")
}

func (q queryResolver) GetWorkoutLogByID(ctx context.Context, workoutLogID string) (*graphql.WorkoutLog, error) {
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
