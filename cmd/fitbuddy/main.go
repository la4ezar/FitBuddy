package main

import (
	"context"
	"fmt"
	"github.com/FitBuddy/internal/domain/coach"
	"github.com/FitBuddy/internal/domain/forum"
	"github.com/FitBuddy/internal/domain/goal"
	"github.com/FitBuddy/internal/domain/leaderboard"
	"github.com/FitBuddy/internal/domain/nutrition"
	"github.com/FitBuddy/internal/domain/sleep"
	"github.com/FitBuddy/internal/domain/user"
	"github.com/FitBuddy/internal/domain/workout"
	"github.com/FitBuddy/pkg/log"
	"github.com/FitBuddy/pkg/persistence"
	"github.com/pkg/errors"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	ctx := context.TODO()

	databaseCfg := persistence.DatabaseConfig{
		User:               "",
		Password:           "",
		Host:               "",
		Port:               "",
		Name:               "",
		SSLMode:            "",
		MaxOpenConnections: 0,
		MaxIdleConnections: 0,
		ConnMaxLifetime:    0,
	}
	db, closeFunc, err := persistence.Configure(ctx, databaseCfg)
	exitOnError(err, "Error while establishing the connection to the database")
	defer func() {
		err := closeFunc()
		exitOnError(err, "Error while closing the connection to the database")
	}()

	// Create repositories for each aggregate
	userRepository := user.NewRepository(db)
	coachRepository := coach.NewRepository(db)
	workoutRepository := workout.NewRepository(db)
	sleepRepository := sleep.NewRepository(db)
	forumRepository := forum.NewRepository(db)
	goalRepository := goal.NewRepository(db)
	nutritionRepository := nutrition.NewRepository(db)
	leaderboardRepository := leaderboard.NewRepository(db)

	// Create services for each aggregate
	userService := user.NewService(userRepository)
	coachService := coach.NewService(coachRepository)
	workoutService := workout.NewService(workoutRepository)
	sleepService := sleep.NewService(sleepRepository)
	forumService := forum.NewService(forumRepository)
	goalService := goal.NewService(goalRepository)
	nutritionService := nutrition.NewService(nutritionRepository)
	leaderboardService := leaderboard.NewService(leaderboardRepository)

	// Create resolvers for each aggregate
	userResolver := user.NewResolver(userService)
	coachResolver := coach.NewResolver(coachService)
	workoutResolver := workout.NewResolver(workoutService)
	sleepResolver := sleep.NewResolver(sleepService)
	forumResolver := forum.NewResolver(forumService)
	goalResolver := goal.NewResolver(goalService)
	nutritionResolver := nutrition.NewNutritionResolver(nutritionService)
	leaderboardResolver := leaderboard.NewLeaderboardResolver(leaderboardService)

	// Create a new mux router
	router := mux.NewRouter()

	// Register GraphQL handlers
	server.RegisterGraphQLHandlers(router, userResolver, coachResolver, workoutResolver, sleepResolver, forumResolver, goalResolver, nutritionResolver, leaderboardResolver)

	// Serve the API on a specified port
	port := 8080
	address := fmt.Sprintf(":%d", port)
	log.Printf("Server is running on http://localhost%s", address)
	log.Fatal(http.ListenAndServe(address, router))
}

func exitOnError(err error, context string) {
	if err != nil {
		wrappedError := errors.Wrap(err, context)
		log.D().Fatal(wrappedError)
	}
}
