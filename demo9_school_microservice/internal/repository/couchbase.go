// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-09-24
package repository

import (
	"context"
	"fmt"
	"log"
	"school-microservice/internal/config"
	"strings"
	"time"

	"github.com/couchbase/gocb/v2"
)

// CouchbaseDB wraps the Couchbase cluster and collection
type CouchbaseDB struct {
	cluster    *gocb.Cluster
	bucket     *gocb.Bucket
	collection *gocb.Collection
}

// NewCouchbaseDB creates a new Couchbase database connection
func NewCouchbaseDB(cfg *config.CouchbaseConfig) (*CouchbaseDB, error) {
	// Configure cluster options
	opts := gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: cfg.Username,
			Password: cfg.Password,
		},
		TimeoutsConfig: gocb.TimeoutsConfig{
			ConnectTimeout: cfg.ConnectTimeout,
			KVTimeout:      cfg.KVTimeout,
		},
	}

	// Connect to cluster
	cluster, err := gocb.Connect(cfg.ConnectionString, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Couchbase cluster: %w", err)
	}

	// Wait until cluster is ready
	err = cluster.WaitUntilReady(cfg.ConnectTimeout, nil)
	if err != nil {
		return nil, fmt.Errorf("cluster not ready: %w", err)
	}

	// Open bucket
	bucket := cluster.Bucket(cfg.BucketName)

	// Wait until bucket is ready
	err = bucket.WaitUntilReady(cfg.ConnectTimeout, nil)
	if err != nil {
		return nil, fmt.Errorf("bucket not ready: %w", err)
	}

	// Get collection
	scope := bucket.Scope(cfg.ScopeName)
	collection := scope.Collection(cfg.CollectionName)

	log.Printf("Successfully connected to Couchbase cluster: %s, bucket: %s", cfg.ConnectionString, cfg.BucketName)

	return &CouchbaseDB{
		cluster:    cluster,
		bucket:     bucket,
		collection: collection,
	}, nil
}

// Close closes the Couchbase connection
func (db *CouchbaseDB) Close() error {
	if db.cluster != nil {
		return db.cluster.Close(nil)
	}
	return nil
}

// CreateIndexes creates necessary indexes for better query performance
func (db *CouchbaseDB) CreateIndexes(ctx context.Context) error {
	indexes := []struct {
		name   string
		fields []string
		where  string
	}{
		{"idx_student_email", []string{"email"}, "type = 'student'"},
		{"idx_student_grade", []string{"grade"}, "type = 'student'"},
		{"idx_teacher_email", []string{"email"}, "type = 'teacher'"},
		{"idx_teacher_department", []string{"department"}, "type = 'teacher'"},
		{"idx_teacher_subject", []string{"subject"}, "type = 'teacher'"},
		{"idx_class_teacher", []string{"teacher_id"}, "type = 'class'"},
		{"idx_class_grade", []string{"grade"}, "type = 'class'"},
		{"idx_class_subject", []string{"subject"}, "type = 'class'"},
		{"idx_academic_student", []string{"student_id"}, "type = 'academic'"},
		{"idx_academic_class", []string{"class_id"}, "type = 'academic'"},
		{"idx_academic_student_subject", []string{"student_id", "subject"}, "type = 'academic'"},
		{"idx_achievement_student", []string{"student_id"}, "type = 'achievement'"},
		{"idx_achievement_category", []string{"category"}, "type = 'achievement'"},
		{"idx_achievement_level", []string{"level"}, "type = 'achievement'"},
		{"idx_student_class", []string{"student_id", "class_id"}, "type = 'student_class'"},
	}

	for _, idx := range indexes {
		// Create index query manually
		var whereClause string
		if idx.where != "" {
			whereClause = fmt.Sprintf(" WHERE %s", idx.where)
		}
		
		fieldsStr := fmt.Sprintf("`%s`", strings.Join(idx.fields, "`, `"))
		indexQuery := fmt.Sprintf("CREATE INDEX `%s` ON `%s`(%s)%s", 
			idx.name, db.bucket.Name(), fieldsStr, whereClause)

		_, err := db.cluster.Query(indexQuery, &gocb.QueryOptions{
			Timeout: 30 * time.Second,
		})
		if err != nil {
			// Check if index already exists
			if strings.Contains(err.Error(), "already exists") {
				log.Printf("Index %s already exists, skipping", idx.name)
				continue
			}
			log.Printf("Failed to create index %s: %v", idx.name, err)
			continue
		}

		log.Printf("Created index: %s", idx.name)
	}

	return nil
}

// Query executes a N1QL query
func (db *CouchbaseDB) Query(ctx context.Context, statement string, options *gocb.QueryOptions) (*gocb.QueryResult, error) {
	return db.cluster.Query(statement, options)
}

// Get retrieves a document by key
func (db *CouchbaseDB) Get(ctx context.Context, key string, valuePtr interface{}) error {
	result, err := db.collection.Get(key, nil)
	if err != nil {
		return err
	}
	return result.Content(valuePtr)
}

// Insert inserts a new document
func (db *CouchbaseDB) Insert(ctx context.Context, key string, doc interface{}) error {
	_, err := db.collection.Insert(key, doc, nil)
	return err
}

// Upsert inserts or updates a document
func (db *CouchbaseDB) Upsert(ctx context.Context, key string, doc interface{}) error {
	_, err := db.collection.Upsert(key, doc, nil)
	return err
}

// Remove removes a document
func (db *CouchbaseDB) Remove(ctx context.Context, key string) error {
	_, err := db.collection.Remove(key, nil)
	return err
}

// Exists checks if a document exists
func (db *CouchbaseDB) Exists(ctx context.Context, key string) (bool, error) {
	result, err := db.collection.Exists(key, nil)
	if err != nil {
		return false, err
	}
	return result.Exists(), nil
}

// generateID generates a unique ID for documents
func generateID(prefix string) string {
	return fmt.Sprintf("%s::%d", prefix, time.Now().UnixNano())
}