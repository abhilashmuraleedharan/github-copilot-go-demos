#!/bin/bash
# [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-09-24
# Couchbase initialization script

echo "Waiting for Couchbase to start..."
sleep 30

# Setup cluster
couchbase-cli cluster-init \
    --cluster localhost \
    --cluster-username Administrator \
    --cluster-password password \
    --services data,index,query \
    --cluster-ramsize 1024 \
    --cluster-index-ramsize 512

# Create bucket
couchbase-cli bucket-create \
    --cluster localhost \
    --username Administrator \
    --password password \
    --bucket school \
    --bucket-type couchbase \
    --bucket-ramsize 512

echo "Couchbase initialization completed"