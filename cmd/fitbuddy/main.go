package main

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/FitBuddy/internal/domain"
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
	"github.com/FitBuddy/pkg/persistence"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/rs/cors"
	"net/http"
	"os"
	"time"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Current working directory:", cwd)

	ctx := context.TODO()
	databaseCfg := persistence.DatabaseConfig{
		User:               "postgres",
		Password:           "pgsql@12345",
		Host:               "127.0.0.1",
		Port:               "5432",
		Name:               "fitbuddy",
		SSLMode:            "disable",
		MaxOpenConnections: 10,
		MaxIdleConnections: 10,
		ConnMaxLifetime:    30 * time.Second,
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
	exerciseRepository := exercise.NewRepository(db)
	workoutRepository := workout.NewRepository(db)
	sleepRepository := sleep.NewRepository(db)
	forumRepository := forum.NewRepository(db)
	goalRepository := goal.NewRepository(db)
	nutritionRepository := nutrition.NewRepository(db)
	leaderboardRepository := leaderboard.NewRepository(db)

	// Create services for each aggregate
	userService := user.NewService(userRepository)
	coachService := coach.NewService(coachRepository)
	exerciseService := exercise.NewService(exerciseRepository)
	workoutService := workout.NewService(workoutRepository)
	sleepService := sleep.NewService(sleepRepository)
	forumService := forum.NewService(forumRepository)
	goalService := goal.NewService(goalRepository)
	nutritionService := nutrition.NewService(nutritionRepository)
	leaderboardService := leaderboard.NewService(leaderboardRepository)

	// Create resolvers for each aggregate
	userResolver := user.NewResolver(userService, leaderboardService)
	coachResolver := coach.NewResolver(coachService)
	exerciseResolver := exercise.NewResolver(exerciseService)
	workoutResolver := workout.NewResolver(workoutService)
	sleepResolver := sleep.NewResolver(sleepService)
	forumResolver := forum.NewResolver(forumService)
	goalResolver := goal.NewResolver(goalService, leaderboardService)
	nutritionResolver := nutrition.NewNutritionResolver(nutritionService)
	leaderboardResolver := leaderboard.NewLeaderboardResolver(leaderboardService)

	// Create a new mux router
	mainRouter := mux.NewRouter()

	PlaygroundAPIEndpoint := "/graphql"
	mainRouter.HandleFunc("/", playground.Handler("Dataloader", PlaygroundAPIEndpoint))

	rootResolver := domain.NewRootResolver(userResolver, coachResolver, exerciseResolver, forumResolver, goalResolver, leaderboardResolver, nutritionResolver, sleepResolver, workoutResolver)

	gqlCfg := graphql.Config{
		Resolvers: rootResolver,
	}

	executableSchema := graphql.NewExecutableSchema(gqlCfg)

	gqlAPIEndpoint := "/graphql"
	gqlAPIRouter := mainRouter.PathPrefix(gqlAPIEndpoint).Subrouter()

	gqlServ := handler.NewDefaultServer(executableSchema)
	gqlServ.SetRecoverFunc(func(ctx context.Context, err interface{}) (userMessage error) {
		errText := fmt.Sprintf("%+v", err)

		return errors.New(errText)
	})

	gqlAPIRouter.HandleFunc("", gqlServ.ServeHTTP)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:63342"}, // Add your frontend origin
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	// Wrap your existing router or handler with the CORS handler
	corsHandler := c.Handler(mainRouter)

	// Serve the API on a specified port
	ServerTimeout := time.Second * 30
	runMainSrv, shutdownMainSrv := createServer(ctx, "localhost:8080", corsHandler, "main", ServerTimeout)

	go func() {
		<-ctx.Done()
		// Interrupt signal received - shut down the server
		shutdownMainSrv()
	}()

	runMainSrv()
}

func exitOnError(err error, context string) {
	if err != nil {
		wrappedError := errors.Wrap(err, context)
		log.D().Fatal(wrappedError)
	}
}

func createServer(ctx context.Context, address string, handler http.Handler, name string, timeout time.Duration) (func(), func()) {
	srv := &http.Server{
		Addr:              address,
		Handler:           handler,
		ReadHeaderTimeout: timeout,
	}

	runFn := func() {
		log.C(ctx).Infof("Running %s server on %s...", name, address)
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.C(ctx).WithError(err).Errorf("An error has occurred with %s HTTP server when ListenAndServe: %v", name, err)
		}
	}

	shutdownFn := func() {
		log.C(ctx).Infof("Shutting down %s server...", name)
		if err := srv.Shutdown(context.Background()); err != nil {
			log.C(ctx).WithError(err).Errorf("An error has occurred while shutting down HTTP server %s: %v", name, err)
		}
	}

	return runFn, shutdownFn
}
