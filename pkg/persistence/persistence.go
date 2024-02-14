package persistence

import (
	"context"
	"database/sql"
	"github.com/FitBuddy/pkg/log"
	"time"

	_ "github.com/lib/pq"
)

const (
	// RetryCount is a number of retries when trying to open the database
	RetryCount int = 50
)

// Configure returns the instance of the database
func Configure(context context.Context, conf *DatabaseConfig) (*sql.DB, func() error, error) {
	db, closeFunc, err := waitForPersistence(context, conf, RetryCount)

	return db, closeFunc, err
}

func waitForPersistence(ctx context.Context, conf *DatabaseConfig, retryCount int) (*sql.DB, func() error, error) {
	var sqlDB *sql.DB
	var err error

	for i := 0; i < retryCount; i++ {
		if i > 0 {
			time.Sleep(5 * time.Second)
		}
		log.C(ctx).Info("Trying to connect to DB...")

		sqlDB, err = sql.Open("postgres", conf.GetConnString())
		if err != nil {
			return nil, nil, err
		}
		ctxWithTimeout, cancelFunc := context.WithTimeout(ctx, time.Second)
		err = sqlDB.PingContext(ctxWithTimeout)
		cancelFunc()
		if err != nil {
			log.C(ctx).Infof("Got error on pinging DB: %v", err)
			continue
		}

		log.C(ctx).Infof("Configuring MaxOpenConnections: [%d], MaxIdleConnections: [%d], ConnectionMaxLifetime: [%s]", conf.MaxOpenConnections, conf.MaxIdleConnections, conf.ConnMaxLifetime.String())
		sqlDB.SetMaxOpenConns(conf.MaxOpenConnections)
		sqlDB.SetMaxIdleConns(conf.MaxIdleConnections)
		sqlDB.SetConnMaxLifetime(conf.ConnMaxLifetime)
		return sqlDB, sqlDB.Close, nil
	}

	return nil, nil, err
}
