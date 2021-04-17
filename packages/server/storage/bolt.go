package storage

import (
	"log"

	"github.com/boltdb/bolt"
	"github.com/jecrocker/nats-baggage/packages/server/logging"
	"go.uber.org/zap"
)

// BaggageBolt defines the BoldDB implementation of the store
type BaggageBolt struct {
	db     *bolt.DB
	logger *zap.SugaredLogger
}

func New() BaggageBolt {
	conn, err := bolt.Open("baggage.db", 0600, nil)
	if err != nil {
		log.Fatalf("Could not open DB: %v", err)
	}

	return BaggageBolt{
		db:     conn,
		logger: logging.New("storage"),
	}
}

func (b *BaggageBolt) Store(topic string, content string) {
	b.logger.Infow("Store Message", "topic", topic, "content-length", len(content))
	// Prevent messages that are too stale from being retrieved later on
	b.Tidy()
}

func (b *BaggageBolt) Get(topic string, timestamp string) {
	b.logger.Infow("Get Messages", "topic", topic, "from", timestamp)
	// Prevent messages that are too stale from being retrieved now
	b.Tidy()
}

func (b *BaggageBolt) Tidy() {
	b.logger.Infow("Tidying DB")
}
