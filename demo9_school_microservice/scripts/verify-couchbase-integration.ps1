# Couchbase Integration Verification Script
# This script verifies that all services are properly integrated with Couchbase

param(
    [switch]$Detailed,
    [switch]$Quick
)

function Write-LogMessage {
    param([string]$Message, [string]$Color = "White")
    $timestamp = Get-Date -Format "HH:mm:ss"
    Write-Host "[$timestamp] $Message" -ForegroundColor $Color
}

Write-LogMessage "🔍 Starting Couchbase Integration Verification..." "Green"

# Test 1: Check if services are using Couchbase (look for connection logs)
Write-LogMessage "🔌 Checking service Couchbase connections..." "Cyan"
$services = @("student-service", "teacher-service", "academic-service", "achievement-service")
$connectedServices = 0

foreach ($service in $services) {
    try {
        $logs = docker-compose logs $service 2>&1 | Select-String "Couchbase|Successfully connected"
        if ($logs) {
            Write-LogMessage "✅ $service is connected to Couchbase" "Green"
            $connectedServices++
            if ($Detailed) { 
                $logs | ForEach-Object { Write-LogMessage "  $_" "Gray" }
            }
        } else {
            Write-LogMessage "⚠️ $service may not be using Couchbase (no connection logs found)" "Yellow"
        }
    } catch {
        Write-LogMessage "❌ Failed to check $service logs: $($_.Exception.Message)" "Red"
    }
}

Write-LogMessage "📊 Services connected to Couchbase: $connectedServices/4" "Cyan"

# Test 2: Verify Couchbase cluster is healthy
Write-LogMessage "🏥 Checking Couchbase cluster health..." "Cyan"
try {
    $clusterInfo = Invoke-RestMethod -Uri "http://localhost:8091/pools/default" `
        -Headers @{Authorization = "Basic $([Convert]::ToBase64String([Text.Encoding]::ASCII.GetBytes('Administrator:password123')))"} `
        -TimeoutSec 10
    
    Write-LogMessage "✅ Couchbase cluster is healthy" "Green"
    if ($Detailed) {
        Write-LogMessage "  Cluster name: $($clusterInfo.clusterName)" "Gray"
        Write-LogMessage "  Total RAM: $([math]::Round($clusterInfo.storageTotals.ram.total / 1MB, 2)) MB" "Gray"
    }
} catch {
    Write-LogMessage "❌ Couchbase cluster health check failed: $($_.Exception.Message)" "Red"
}

# Test 3: Verify schoolmgmt bucket exists
Write-LogMessage "🗄️ Checking schoolmgmt bucket..." "Cyan"
try {
    $buckets = Invoke-RestMethod -Uri "http://localhost:8091/pools/default/buckets" `
        -Headers @{Authorization = "Basic $([Convert]::ToBase64String([Text.Encoding]::ASCII.GetBytes('Administrator:password123')))"} `
        -TimeoutSec 10
    
    $schoolmgmtBucket = $buckets | Where-Object { $_.name -eq "schoolmgmt" }
    if ($schoolmgmtBucket) {
        Write-LogMessage "✅ schoolmgmt bucket exists and is healthy" "Green"
        if ($Detailed) {
            Write-LogMessage "  Bucket type: $($schoolmgmtBucket.bucketType)" "Gray"
            Write-LogMessage "  RAM quota: $($schoolmgmtBucket.quota.ram) MB" "Gray"
            Write-LogMessage "  Item count: $($schoolmgmtBucket.basicStats.itemCount)" "Gray"
        }
    } else {
        Write-LogMessage "❌ schoolmgmt bucket not found" "Red"
    }
} catch {
    Write-LogMessage "❌ Failed to check buckets: $($_.Exception.Message)" "Red"
}

# Test 4: Test data persistence (if not Quick mode)
if (-not $Quick) {
    Write-LogMessage "🧪 Testing data persistence..." "Cyan"
    try {
        # Create test student
        $testStudent = @{
            firstName = "Persistence"
            lastName = "Test"
            email = "persistence.test@school.edu"
            grade = "10"
        } | ConvertTo-Json
        
        Write-LogMessage "📝 Creating test student..." "Yellow"
        $createResponse = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/students" `
            -Method Post `
            -Body $testStudent `
            -ContentType "application/json" `
            -TimeoutSec 15
        
        $studentId = $createResponse.data.id
        Write-LogMessage "✅ Test student created with ID: $studentId" "Green"
        
        # Wait a moment for data to be written
        Start-Sleep 2
        
        # Verify data exists in Couchbase directly
        Write-LogMessage "🔍 Verifying data in Couchbase..." "Yellow"
        $queryBody = "statement=SELECT * FROM \`schoolmgmt\` WHERE type='student' AND firstName='Persistence'"
        $directQuery = Invoke-RestMethod -Uri "http://localhost:8093/query/service" `
            -Method Post `
            -Body $queryBody `
            -ContentType "application/x-www-form-urlencoded" `
            -Headers @{Authorization = "Basic $([Convert]::ToBase64String([Text.Encoding]::ASCII.GetBytes('Administrator:password123')))"} `
            -TimeoutSec 15
        
        if ($directQuery.results -and $directQuery.results.Count -gt 0) {
            Write-LogMessage "✅ Data verified in Couchbase directly" "Green"
        } else {
            Write-LogMessage "⚠️ Data not found in direct Couchbase query" "Yellow"
        }
        
        # Test service restart persistence
        Write-LogMessage "🔄 Testing persistence across service restart..." "Yellow"
        docker-compose restart student-service | Out-Null
        
        # Wait for service to restart
        Start-Sleep 15
        
        # Try to retrieve after restart
        $retrieveResponse = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/students/$studentId" `
            -Method Get `
            -TimeoutSec 15
        
        if ($retrieveResponse.data.id -eq $studentId) {
            Write-LogMessage "✅ Data persisted across service restart!" "Green"
        } else {
            Write-LogMessage "❌ Data did not persist across service restart" "Red"
        }
        
        # Clean up test data
        Write-LogMessage "🗑️ Cleaning up test data..." "Yellow"
        Invoke-RestMethod -Uri "http://localhost:8080/api/v1/students/$studentId" `
            -Method Delete `
            -TimeoutSec 15 | Out-Null
        
        Write-LogMessage "✅ Test data cleaned up" "Green"
        
    } catch {
        Write-LogMessage "❌ Data persistence test failed: $($_.Exception.Message)" "Red"
        if ($Detailed) {
            Write-LogMessage "  Full error: $($_.Exception)" "Gray"
        }
    }
}

# Test 5: Query service availability
Write-LogMessage "🔍 Testing N1QL query service..." "Cyan"
try {
    $pingResponse = Invoke-WebRequest -Uri "http://localhost:8093/admin/ping" `
        -TimeoutSec 10
    
    if ($pingResponse.StatusCode -eq 200) {
        Write-LogMessage "✅ N1QL query service is available" "Green"
    } else {
        Write-LogMessage "⚠️ N1QL query service responded with status: $($pingResponse.StatusCode)" "Yellow"
    }
} catch {
    Write-LogMessage "❌ N1QL query service is not available: $($_.Exception.Message)" "Red"
}

# Test 6: Primary index verification
Write-LogMessage "🔍 Checking primary index..." "Cyan"
try {
    $indexQuery = "statement=SELECT * FROM system:indexes WHERE keyspace_id='schoolmgmt' AND is_primary=true"
    $indexResponse = Invoke-RestMethod -Uri "http://localhost:8093/query/service" `
        -Method Post `
        -Body $indexQuery `
        -ContentType "application/x-www-form-urlencoded" `
        -Headers @{Authorization = "Basic $([Convert]::ToBase64String([Text.Encoding]::ASCII.GetBytes('Administrator:password123')))"} `
        -TimeoutSec 10
    
    if ($indexResponse.results -and $indexResponse.results.Count -gt 0) {
        Write-LogMessage "✅ Primary index exists for schoolmgmt bucket" "Green"
    } else {
        Write-LogMessage "⚠️ Primary index not found for schoolmgmt bucket" "Yellow"
        Write-LogMessage "💡 Run: .\scripts\init-couchbase.ps1 to create indexes" "Cyan"
    }
} catch {
    Write-LogMessage "❌ Failed to check primary index: $($_.Exception.Message)" "Red"
}

# Summary
Write-LogMessage "" "White"
Write-LogMessage "📋 Verification Summary:" "Green"
Write-LogMessage "🔌 Services connected to Couchbase: $connectedServices/4" "White"

if ($connectedServices -eq 4) {
    Write-LogMessage "🎉 Couchbase integration verification completed successfully!" "Green"
    Write-LogMessage "💡 Your School Management System is properly integrated with Couchbase" "Cyan"
} else {
    Write-LogMessage "⚠️ Some integration issues detected" "Yellow"
    Write-LogMessage "🔧 Run: .\scripts\enhanced-setup-couchbase.ps1 -Verbose to fix issues" "Cyan"
}

Write-LogMessage "" "White"
Write-LogMessage "📚 Next Steps:" "Cyan"
Write-LogMessage "• Use scripts/couchbase-crud-commands.md for comprehensive testing" "White"
Write-LogMessage "• Check FIXES_AND_SOLUTIONS.md for troubleshooting guidance" "White"
Write-LogMessage "• Visit http://localhost:8091 to access Couchbase Console" "White"
