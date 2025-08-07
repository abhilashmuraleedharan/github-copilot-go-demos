# Simple Couchbase initialization script
param(
    [string]$CouchbaseHost = "localhost",
    [string]$CouchbaseUser = "Administrator", 
    [string]$CouchbasePass = "password123",
    [string]$CouchbaseBucket = "schoolmgmt"
)

Write-Host "🚀 Starting Couchbase initialization..." -ForegroundColor Green

# Wait for Couchbase to be accessible
Write-Host "⏳ Waiting for Couchbase to be ready..." -ForegroundColor Yellow
do {
    try {
        $response = Invoke-WebRequest -Uri "http://$CouchbaseHost`:8091/pools" -TimeoutSec 5 -ErrorAction Stop
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
    $clusterBody = "memoryQuota=512" + "&indexMemoryQuota=256"
    Invoke-RestMethod -Uri "http://$CouchbaseHost`:8091/pools/default" -Method Post -Body $clusterBody -ContentType "application/x-www-form-urlencoded" -ErrorAction Stop
    Write-Host "✅ Cluster initialized!" -ForegroundColor Green
} catch {
    Write-Host "⚠️ Cluster may already be initialized" -ForegroundColor Yellow
}

# Set up administrator
Write-Host "👥 Setting up administrator..." -ForegroundColor Yellow
try {
    $adminBody = "username=$CouchbaseUser" + "&password=$CouchbasePass" + "&port=SAME"
    Invoke-RestMethod -Uri "http://$CouchbaseHost`:8091/settings/web" -Method Post -Body $adminBody -ContentType "application/x-www-form-urlencoded" -ErrorAction Stop
    Write-Host "✅ Administrator set up!" -ForegroundColor Green
} catch {
    Write-Host "⚠️ User may already exist" -ForegroundColor Yellow
}

Start-Sleep 5

# Create bucket
Write-Host "🗄️ Creating bucket: $CouchbaseBucket" -ForegroundColor Yellow
try {
    $bucketBody = "name=$CouchbaseBucket" + "&ramQuotaMB=256" + "&bucketType=membase"
    $auth = [System.Convert]::ToBase64String([System.Text.Encoding]::ASCII.GetBytes("$CouchbaseUser`:$CouchbasePass"))
    $headers = @{ Authorization = "Basic $auth" }
    
    Invoke-RestMethod -Uri "http://$CouchbaseHost`:8091/pools/default/buckets" -Method Post -Body $bucketBody -ContentType "application/x-www-form-urlencoded" -Headers $headers -ErrorAction Stop
    Write-Host "✅ Bucket created!" -ForegroundColor Green
} catch {
    Write-Host "⚠️ Bucket may already exist" -ForegroundColor Yellow
}

Start-Sleep 10

# Create primary index
Write-Host "🔍 Creating primary index..." -ForegroundColor Yellow
try {
    $indexBody = "statement=CREATE PRIMARY INDEX ON \`$CouchbaseBucket\`"
    $auth = [System.Convert]::ToBase64String([System.Text.Encoding]::ASCII.GetBytes("$CouchbaseUser`:$CouchbasePass"))
    $headers = @{ Authorization = "Basic $auth" }
    
    Invoke-RestMethod -Uri "http://$CouchbaseHost`:8093/query/service" -Method Post -Body $indexBody -ContentType "application/x-www-form-urlencoded" -Headers $headers -ErrorAction Stop
    Write-Host "✅ Primary index created!" -ForegroundColor Green
} catch {
    Write-Host "⚠️ Index may already exist" -ForegroundColor Yellow
}

Write-Host "🎉 Couchbase initialization completed!" -ForegroundColor Green
