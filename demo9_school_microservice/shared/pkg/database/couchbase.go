package database

import (
	"fmt"
	"time"

	"github.com/couchbase/gocb/v2"
	"github.com/sirupsen/logrus"
	"schoolmgmt/shared/pkg/config"
)

type CouchbaseClient struct {
	Cluster    *gocb.Cluster
	Bucket     *gocb.Bucket
	Collection *gocb.Collection
}

func NewCouchbaseClient(cfg *config.Config) (*CouchbaseClient, error) {
	connectionString := fmt.Sprintf("couchbase://%s", cfg.CouchbaseHost)
	
	logrus.Infof("🔌 Attempting to connect to Couchbase at: %s", connectionString)
	logrus.Infof("📝 Using credentials - User: %s, Bucket: %s", cfg.CouchbaseUser, cfg.CouchbaseBucket)
	
	// Connect to cluster
	cluster, err := gocb.Connect(connectionString, gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: cfg.CouchbaseUser,
			Password: cfg.CouchbasePass,
		},
		TimeoutsConfig: gocb.TimeoutsConfig{
			ConnectTimeout:    10 * time.Second,
			KVTimeout:         5 * time.Second,
			QueryTimeout:      15 * time.Second,
			AnalyticsTimeout:  15 * time.Second,
			SearchTimeout:     15 * time.Second,
			ManagementTimeout: 15 * time.Second,
		},
	})
	if err != nil {
		logrus.Errorf("❌ Failed to connect to Couchbase cluster: %v", err)
		return nil, fmt.Errorf("failed to connect to Couchbase cluster: %w", err)
	}
	
	logrus.Infof("✅ Connected to Couchbase cluster, waiting for readiness...")

	// Wait for cluster to be ready
	err = cluster.WaitUntilReady(30*time.Second, nil)
	if err != nil {
		logrus.Errorf("❌ Cluster not ready: %v", err)
		return nil, fmt.Errorf("cluster not ready: %w", err)
	}
	
	logrus.Infof("✅ Cluster is ready, accessing bucket: %s", cfg.CouchbaseBucket)

	// Get bucket
	bucket := cluster.Bucket(cfg.CouchbaseBucket)
	
	// Wait for bucket to be ready
	err = bucket.WaitUntilReady(30*time.Second, nil)
	if err != nil {
		logrus.Warnf("⚠️ Bucket %s not ready, will create it: %v", cfg.CouchbaseBucket, err)
		
		// Try to create bucket if it doesn't exist
		bucketManager := cluster.Buckets()
		bucketSettings := gocb.CreateBucketSettings{
			BucketSettings: gocb.BucketSettings{
				Name:                 cfg.CouchbaseBucket,
				RAMQuotaMB:           256,
				BucketType:           gocb.CouchbaseBucketType,
				MaxTTL:               0,
				CompressionMode:      gocb.CompressionModePassive,
			},
		}
		
		logrus.Infof("🏗️ Creating bucket: %s with %d MB RAM", cfg.CouchbaseBucket, 256)
		err = bucketManager.CreateBucket(bucketSettings, nil)
		if err != nil {
			logrus.Warnf("⚠️ Failed to create bucket (might already exist): %v", err)
		} else {
			logrus.Infof("✅ Created bucket: %s", cfg.CouchbaseBucket)
		}
		
		// Get bucket again
		bucket = cluster.Bucket(cfg.CouchbaseBucket)
		logrus.Infof("⏳ Waiting for newly created bucket to be ready...")
		err = bucket.WaitUntilReady(30*time.Second, nil)
		if err != nil {
			logrus.Warnf("⚠️ Bucket still not ready after creation: %v", err)
		} else {
			logrus.Infof("✅ Bucket is now ready!")
		}
	} else {
		logrus.Infof("✅ Bucket %s is ready!", cfg.CouchbaseBucket)
	}

	// Get default collection
	collection := bucket.DefaultCollection()
	logrus.Infof("✅ Connected to default collection")

	logrus.Infof("🎉 Successfully connected to Couchbase bucket: %s", cfg.CouchbaseBucket)

	return &CouchbaseClient{
		Cluster:    cluster,
		Bucket:     bucket,
		Collection: collection,
	}, nil
}

func (c *CouchbaseClient) Close() {
	if c.Cluster != nil {
		c.Cluster.Close(nil)
	}
}
