package log

import (
	"context"
	"github.com/sirupsen/logrus"
	"sync"
)

type logKey struct{}

var (
	defaultEntry = logrus.NewEntry(logrus.StandardLogger())

	mutex = sync.RWMutex{}

	// C gives us a logger with context
	C = LoggerFromContext
	// D gives us a default logger
	D = DefaultLogger
)

// LoggerFromContext retrieves the current logger from the context.
func LoggerFromContext(ctx context.Context) *logrus.Entry {
	mutex.RLock()
	defer mutex.RUnlock()
	entry := ctx.Value(logKey{})
	if entry == nil {
		entry = defaultEntry
	}
	return copyEntry(entry.(*logrus.Entry))
}

// DefaultLogger returns the default logger
func DefaultLogger() *logrus.Entry {
	return LoggerFromContext(context.Background())
}

func copyEntry(entry *logrus.Entry) *logrus.Entry {
	entryData := make(logrus.Fields, len(entry.Data))
	for k, v := range entry.Data {
		entryData[k] = v
	}

	newEntry := logrus.NewEntry(entry.Logger)
	newEntry.Level = entry.Level
	newEntry.Data = entryData
	newEntry.Time = entry.Time
	newEntry.Message = entry.Message
	newEntry.Buffer = entry.Buffer

	return newEntry
}