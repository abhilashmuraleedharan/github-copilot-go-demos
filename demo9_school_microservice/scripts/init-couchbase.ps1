# Couchbase initialization script for Windows
# This script initializes the Couchbase cluster and creates necessary buckets

param(
    [string]$CouchbaseHost = "localhost",
    [string]$CouchbaseUser = "Administrator", 
    [string]$CouchbasePass = "password123",
    [string]$CouchbaseBucket = "schoolmgmt"
)

Write-Host "🚀 Starting Couchbase initialization..." -ForegroundColor Green
Write-Host "📍 Host: $CouchbaseHost" -ForegroundColor Cyan
Write-Host "👤 User: $CouchbaseUser" -ForegroundColor Cyan
Write-Host "🗄️ Bucket: $CouchbaseBucket" -ForegroundColor Cyan

# Wait for Couchbase to be accessible
Write-Host "⏳ Waiting for Couchbase to be ready..." -ForegroundColor Yellow
do {
    try {
        $response = Invoke-WebRequest -Uri "http://${CouchbaseHost}:8091/pools" -TimeoutSec 5 -ErrorAction Stop
        Write-Host "✅ Couchbase is accessible!" -ForegroundColor Green
        break
    } catch {
        if ($_.Exception.Response.StatusCode -eq 401) {
            Write-Host "✅ Couchbase is accessible (needs setup)!" -ForegroundColor Green
            break
        }
        Write-Host "⏳ Couchbase not ready yet, waiting..." -ForegroundColor Yellow
        Start-Sleep 5
    }
} while ($true)

# Initialize cluster
Write-Host "🏗️ Initializing Couchbase cluster..." -ForegroundColor Yellow
try {
    $clusterBody = "memoryQuota=512&indexMemoryQuota=256"
    Invoke-RestMethod -Uri "http://${CouchbaseHost}:8091/pools/default" `
        -Method Post `
        -Body $clusterBody `
        -ContentType "application/x-www-form-urlencoded" `
        -ErrorAction Stop
    Write-Host "✅ Cluster initialized successfully!" -ForegroundColor Green
} catch {
    Write-Host "⚠️ Cluster may already be initialized: $($_.Exception.Message)" -ForegroundColor Yellow
}

# Set up administrator
Write-Host "👥 Setting up administrator user..." -ForegroundColor Yellow
try {
    $adminBody = "username=$CouchbaseUser" + "&password=$CouchbasePass" + "&port=SAME"
    Invoke-RestMethod -Uri "http://${CouchbaseHost}:8091/settings/web" `
        -Method Post `
        -Body $adminBody `
        -ContentType "application/x-www-form-urlencoded" `
        -ErrorAction Stop
    Write-Host "✅ Administrator user set up successfully!" -ForegroundColor Green
} catch {
    Write-Host "⚠️ User may already exist: $($_.Exception.Message)" -ForegroundColor Yellow
}

# Wait a bit for settings to take effect
Start-Sleep 5

# Create bucket
Write-Host "🗄️ Creating bucket: $CouchbaseBucket" -ForegroundColor Yellow
try {
    $bucketBody = "name=$CouchbaseBucket" + "&ramQuotaMB=256&bucketType=membase"
    $credentials = [System.Convert]::ToBase64String([System.Text.Encoding]::ASCII.GetBytes("$CouchbaseUser" + ":" + "$CouchbasePass"))
    $headers = @{ Authorization = "Basic $credentials" }
    
    Invoke-RestMethod -Uri "http://${CouchbaseHost}:8091/pools/default/buckets" `
        -Method Post `
        -Body $bucketBody `
        -ContentType "application/x-www-form-urlencoded" `
        -Headers $headers `
        -ErrorAction Stop
    Write-Host "✅ Bucket created successfully!" -ForegroundColor Green
} catch {
    Write-Host "⚠️ Bucket may already exist: $($_.Exception.Message)" -ForegroundColor Yellow
}

# Wait for bucket to be ready
Write-Host "⏳ Waiting for bucket to be operational..." -ForegroundColor Yellow
Start-Sleep 10

# Create a primary index for N1QL queries
Write-Host "🔍 Creating primary index..." -ForegroundColor Yellow
try {
    $indexBody = "statement=CREATE PRIMARY INDEX ON \`$CouchbaseBucket\`"
    $credentials = [System.Convert]::ToBase64String([System.Text.Encoding]::ASCII.GetBytes("$CouchbaseUser" + ":" + "$CouchbasePass"))
    $headers = @{ Authorization = "Basic $credentials" }
    
    Invoke-RestMethod -Uri "http://${CouchbaseHost}:8093/query/service" `
        -Method Post `
        -Body $indexBody `
        -ContentType "application/x-www-form-urlencoded" `
        -Headers $headers `
        -ErrorAction Stop
    Write-Host "✅ Primary index created successfully!" -ForegroundColor Green
} catch {
    Write-Host "⚠️ Index may already exist: $($_.Exception.Message)" -ForegroundColor Yellow
}

Write-Host "🎉 Couchbase initialization completed!" -ForegroundColor Green

# Test connection
Write-Host "🧪 Testing connection..." -ForegroundColor Yellow
try {
    $testBody = "statement=SELECT COUNT(*) as doc_count FROM \`$CouchbaseBucket\`"
    $credentials = [System.Convert]::ToBase64String([System.Text.Encoding]::ASCII.GetBytes("$CouchbaseUser" + ":" + "$CouchbasePass"))
    $headers = @{ Authorization = "Basic $credentials" }
    
    $result = Invoke-RestMethod -Uri "http://${CouchbaseHost}:8093/query/service" `
        -Method Post `
        -Body $testBody `
        -ContentType "application/x-www-form-urlencoded" `
        -Headers $headers `
        -ErrorAction Stop
    
    Write-Host "✅ Connection test successful!" -ForegroundColor Green
    Write-Host "📊 Query result: $($result | ConvertTo-Json)" -ForegroundColor Cyan
} catch {
    Write-Host "❌ Connection test failed: $($_.Exception.Message)" -ForegroundColor Red
}

Write-Host "✅ Couchbase is ready for use!" -ForegroundColor Green
