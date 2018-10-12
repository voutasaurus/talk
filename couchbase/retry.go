import (
	"context"
	"time"
)

func NewDB(ctx context.Context, c *Config) (*DB, error) {
	db, err := newDB(c)
	if err == nil {
		return db, nil
	}

	// If no timeout is set, exit early
	_, ok := ctx.Deadline()
	if !ok {
		return nil, err
	}

	// Wait for DB to become available (or timeout)
	for {
		time.Sleep(100 * time.Millisecond)
		db, err := newDB(c)
		if err == nil {
			return db, nil
		}
		select {
		case <-ctx.Done():
			return nil, err
		default:
		}
	}
}
