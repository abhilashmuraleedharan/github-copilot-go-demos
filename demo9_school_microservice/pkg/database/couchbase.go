// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-20
package database

import (
	"fmt"
	"log"
	"time"

	"github.com/couchbase/gocb/v2"
	"school-microservice/pkg/config"
)

// CouchbaseClient wraps the Couchbase cluster and collection
type CouchbaseClient struct {
	Cluster    *gocb.Cluster
	Collection *gocb.Collection
}

// NewCouchbaseClient creates a new Couchbase client
func NewCouchbaseClient(cfg *config.Config, collectionName string) (*CouchbaseClient, error) {
	// Connect to Couchbase cluster
	connectionString := fmt.Sprintf("couchbase://%s", cfg.CouchbaseHost)
	cluster, err := gocb.Connect(connectionString, gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: cfg.CouchbaseUsername,
			Password: cfg.CouchbasePassword,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Couchbase: %w", err)
	}

	// Wait for cluster to be ready
	err = cluster.WaitUntilReady(10*time.Second, nil)
	if err != nil {
		return nil, fmt.Errorf("cluster not ready: %w", err)
	}

	// Get bucket
	bucket := cluster.Bucket(cfg.CouchbaseBucket)

	// Wait for bucket to be ready
	err = bucket.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		return nil, fmt.Errorf("bucket not ready: %w", err)
	}

	// Get collection
	collection := bucket.DefaultCollection()
	// For simplicity in this demo, we'll use the default collection for all services
	// In production, you might want to use separate collections per service

	log.Printf("Connected to Couchbase cluster at %s, bucket: %s, collection: %s", 
		cfg.CouchbaseHost, cfg.CouchbaseBucket, collectionName)

	return &CouchbaseClient{
		Cluster:    cluster,
		Collection: collection,
	}, nil
}

// Close closes the Couchbase connection
func (c *CouchbaseClient) Close() error {
	if c.Cluster != nil {
		return c.Cluster.Close(nil)
	}
	return nil
}
