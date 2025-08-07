#!/bin/bash

# Couchbase initialization script
# This script initializes the Couchbase cluster and creates necessary buckets

set -e

CB_HOST=${COUCHBASE_HOST:-localhost}
CB_USER=${COUCHBASE_USERNAME:-Administrator}
CB_PASS=${COUCHBASE_PASSWORD:-password123}
CB_BUCKET=${COUCHBASE_BUCKET:-schoolmgmt}

echo "🚀 Starting Couchbase initialization..."
echo "📍 Host: $CB_HOST"
echo "👤 User: $CB_USER"
echo "🗄️ Bucket: $CB_BUCKET"

# Wait for Couchbase to be accessible
echo "⏳ Waiting for Couchbase to be ready..."
until curl -sf http://$CB_HOST:8091/pools > /dev/null 2>&1; do
    echo "⏳ Couchbase not ready yet, waiting..."
    sleep 5
done

echo "✅ Couchbase is accessible!"

# Initialize cluster
echo "🏗️ Initializing Couchbase cluster..."
curl -v -X POST http://$CB_HOST:8091/pools/default \
    -d 'memoryQuota=512&indexMemoryQuota=256' || echo "⚠️ Cluster may already be initialized"

# Set up administrator
echo "👥 Setting up administrator user..."
curl -v -X POST http://$CB_HOST:8091/settings/web \
    -d "username=$CB_USER&password=$CB_PASS&port=SAME" || echo "⚠️ User may already exist"

# Wait a bit for settings to take effect
sleep 5

# Create bucket
echo "🗄️ Creating bucket: $CB_BUCKET"
curl -v -X POST http://$CB_HOST:8091/pools/default/buckets \
    -u "$CB_USER:$CB_PASS" \
    -d "name=$CB_BUCKET&ramQuotaMB=256&bucketType=membase" || echo "⚠️ Bucket may already exist"

# Wait for bucket to be ready
echo "⏳ Waiting for bucket to be operational..."
sleep 10

# Create a primary index for N1QL queries
echo "🔍 Creating primary index..."
curl -v -X POST http://$CB_HOST:8093/query/service \
    -u "$CB_USER:$CB_PASS" \
    -d "statement=CREATE PRIMARY INDEX ON \`$CB_BUCKET\`" || echo "⚠️ Index may already exist"

echo "🎉 Couchbase initialization completed!"

# Test connection
echo "🧪 Testing connection..."
curl -v -X POST http://$CB_HOST:8093/query/service \
    -u "$CB_USER:$CB_PASS" \
    -d "statement=SELECT COUNT(*) as doc_count FROM \`$CB_BUCKET\`"

echo "✅ Couchbase is ready for use!"
