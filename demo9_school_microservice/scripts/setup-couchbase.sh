#!/bin/bash

# Couchbase Setup Script
# This script initializes Couchbase with the required bucket and indexes

echo "Setting up Couchbase for School Microservice..."

# Wait for Couchbase to be ready
echo "Waiting for Couchbase to be ready..."
sleep 30

# Initialize cluster
echo "Initializing Couchbase cluster..."
docker exec couchbase couchbase-cli cluster-init \
    --cluster couchbase://localhost \
    --cluster-username Administrator \
    --cluster-password password \
    --cluster-name school-cluster \
    --cluster-ramsize 1024 \
    --cluster-index-ramsize 512 \
    --services data,index,query

# Create bucket
echo "Creating school bucket..."
docker exec couchbase couchbase-cli bucket-create \
    --cluster couchbase://localhost \
    --username Administrator \
    --password password \
    --bucket school \
    --bucket-type couchbase \
    --bucket-ramsize 512 \
    --bucket-replica 0

# Wait for bucket to be ready
echo "Waiting for bucket to be ready..."
sleep 10

# Create indexes for better query performance
echo "Creating indexes..."
docker exec couchbase cbq -u Administrator -p password -s="CREATE PRIMARY INDEX ON \`school\`"
docker exec couchbase cbq -u Administrator -p password -s="CREATE INDEX idx_type ON \`school\`(type)"
docker exec couchbase cbq -u Administrator -p password -s="CREATE INDEX idx_student_id ON \`school\`(studentId) WHERE type IN ['academic', 'achievement']"
docker exec couchbase cbq -u Administrator -p password -s="CREATE INDEX idx_teacher_id ON \`school\`(teacherId)"
docker exec couchbase cbq -u Administrator -p password -s="CREATE INDEX idx_class_id ON \`school\`(classId)"

echo "Couchbase setup completed successfully!"
echo "You can access Couchbase Web Console at: http://localhost:8091"
echo "Username: Administrator"
echo "Password: password"
