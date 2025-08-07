# Couchbase Setup and Demo Script for School Management System
# PowerShell Version
# Run this script to initialize Couchbase and demonstrate CRUD operations

param(
    [string]$Action = "all",
    [string]$CouchbaseUrl = "http://localhost:8091",
    [string]$QueryUrl = "http://localhost:8093",
    [string]$Username = "Administrator",
    [string]$Password = "password123"
)

# Color output functions
function Write-Success { param($Message) Write-Host $Message -ForegroundColor Green }
function Write-Error { param($Message) Write-Host $Message -ForegroundColor Red }
function Write-Info { param($Message) Write-Host $Message -ForegroundColor Cyan }
function Write-Warning { param($Message) Write-Host $Message -ForegroundColor Yellow }

function Test-CouchbaseConnection {
    Write-Info "Testing Couchbase connection..."
    try {
        $response = Invoke-WebRequest -Uri "$CouchbaseUrl/pools" -Method GET -TimeoutSec 10
        if ($response.StatusCode -eq 200) {
            Write-Success "✅ Couchbase is accessible"
            return $true
        }
    } catch {
        Write-Error "❌ Couchbase is not accessible at $CouchbaseUrl"
        Write-Warning "Make sure Docker Compose is running: docker-compose up -d"
        return $false
    }
}

function Initialize-Cluster {
    Write-Info "🔧 Initializing Couchbase cluster..."
    
    try {
        # Initialize cluster
        $initResponse = Invoke-WebRequest -Uri "$CouchbaseUrl/pools/default" -Method POST -Body "memoryQuota=512&indexMemoryQuota=256" -ContentType "application/x-www-form-urlencoded"
        Write-Success "Cluster initialized"
        
        # Setup administrator
        $adminResponse = Invoke-WebRequest -Uri "$CouchbaseUrl/settings/web" -Method POST -Body "username=$Username&password=$Password&port=SAME" -ContentType "application/x-www-form-urlencoded"
        Write-Success "Administrator account created"
        
        # Wait for cluster to stabilize
        Start-Sleep -Seconds 5
        
        # Create bucket
        $bucketBody = "name=schoolmgmt&ramQuotaMB=512&bucketType=membase"
        $bucketResponse = Invoke-WebRequest -Uri "$CouchbaseUrl/pools/default/buckets" -Method POST -Body $bucketBody -ContentType "application/x-www-form-urlencoded" -Headers @{Authorization = "Basic $([Convert]::ToBase64String([Text.Encoding]::ASCII.GetBytes("$($Username):$($Password)")))"}
        Write-Success "Bucket 'schoolmgmt' created"
        
    } catch {
        Write-Warning "Cluster may already be initialized: $($_.Exception.Message)"
    }
}

function Create-Collections {
    Write-Info "📁 Creating collections and scopes..."
    
    $auth = [Convert]::ToBase64String([Text.Encoding]::ASCII.GetBytes("$($Username):$($Password)"))
    $headers = @{Authorization = "Basic $auth"}
    
    try {
        # Create scope
        $scopeResponse = Invoke-WebRequest -Uri "$CouchbaseUrl/pools/default/buckets/schoolmgmt/scopes" -Method POST -Body "name=school" -ContentType "application/x-www-form-urlencoded" -Headers $headers
        Write-Success "Scope 'school' created"
        
        # Wait for scope to be available
        Start-Sleep -Seconds 3
        
        # Create collections
        $collections = @("students", "teachers", "academics", "achievements")
        foreach ($collection in $collections) {
            try {
                $collectionResponse = Invoke-WebRequest -Uri "$CouchbaseUrl/pools/default/buckets/schoolmgmt/scopes/school/collections" -Method POST -Body "name=$collection" -ContentType "application/x-www-form-urlencoded" -Headers $headers
                Write-Success "Collection '$collection' created"
                Start-Sleep -Seconds 1
            } catch {
                Write-Warning "Collection '$collection' may already exist"
            }
        }
        
    } catch {
        Write-Warning "Scope/Collections may already exist: $($_.Exception.Message)"
    }
}

function Create-Indexes {
    Write-Info "🏗️ Creating database indexes..."
    
    $auth = [Convert]::ToBase64String([Text.Encoding]::ASCII.GetBytes("$($Username):$($Password)"))
    $headers = @{
        Authorization = "Basic $auth"
        "Content-Type" = "application/json"
    }
    
    $indexes = @(
        "CREATE PRIMARY INDEX ON schoolmgmt.school.students",
        "CREATE PRIMARY INDEX ON schoolmgmt.school.teachers", 
        "CREATE PRIMARY INDEX ON schoolmgmt.school.academics",
        "CREATE PRIMARY INDEX ON schoolmgmt.school.achievements",
        "CREATE INDEX idx_student_grade ON schoolmgmt.school.students(grade)",
        "CREATE INDEX idx_teacher_department ON schoolmgmt.school.teachers(department)",
        "CREATE INDEX idx_academic_student ON schoolmgmt.school.academics(student_id, academic_year)"
    )
    
    foreach ($indexStatement in $indexes) {
        try {
            $queryBody = @{statement = $indexStatement} | ConvertTo-Json
            $response = Invoke-WebRequest -Uri "$QueryUrl/query/service" -Method POST -Body $queryBody -Headers $headers
            Write-Success "Index created: $($indexStatement.Split(' ')[-1])"
            Start-Sleep -Seconds 1
        } catch {
            Write-Warning "Index may already exist: $($indexStatement.Split(' ')[-1])"
        }
    }
}

function Demo-StudentCRUD {
    Write-Info "👨‍🎓 Demonstrating Student CRUD operations..."
    
    $auth = [Convert]::ToBase64String([Text.Encoding]::ASCII.GetBytes("$($Username):$($Password)"))
    $headers = @{
        Authorization = "Basic $auth"
        "Content-Type" = "application/json"
    }
    
    # CREATE Student
    Write-Info "Creating student..."
    $createQuery = @{
        statement = 'INSERT INTO schoolmgmt.school.students (KEY, VALUE) VALUES ("student-001", {
            "id": "student-001",
            "first_name": "John",
            "last_name": "Doe", 
            "email": "john.doe@school.edu",
            "grade": "10",
            "age": 16,
            "enrollment_date": "2024-08-01",
            "status": "active",
            "created_at": "2024-08-05T10:30:00Z"
        })'
    } | ConvertTo-Json
    
    try {
        $response = Invoke-WebRequest -Uri "$QueryUrl/query/service" -Method POST -Body $createQuery -Headers $headers
        Write-Success "✅ Student created successfully"
    } catch {
        Write-Warning "Student may already exist"
    }
    
    # READ Student
    Write-Info "Reading student..."
    $readQuery = @{
        statement = 'SELECT * FROM schoolmgmt.school.students WHERE META().id = "student-001"'
    } | ConvertTo-Json
    
    try {
        $response = Invoke-WebRequest -Uri "$QueryUrl/query/service" -Method POST -Body $readQuery -Headers $headers
        $result = ($response.Content | ConvertFrom-Json)
        if ($result.results.Count -gt 0) {
            Write-Success "✅ Student retrieved: $($result.results[0].students.first_name) $($result.results[0].students.last_name)"
        }
    } catch {
        Write-Error "Failed to read student: $($_.Exception.Message)"
    }
    
    # UPDATE Student
    Write-Info "Updating student..."
    $updateQuery = @{
        statement = 'UPDATE schoolmgmt.school.students SET grade = "11", age = 17 WHERE META().id = "student-001"'
    } | ConvertTo-Json
    
    try {
        $response = Invoke-WebRequest -Uri "$QueryUrl/query/service" -Method POST -Body $updateQuery -Headers $headers
        Write-Success "✅ Student updated successfully"
    } catch {
        Write-Error "Failed to update student: $($_.Exception.Message)"
    }
    
    # LIST Students
    Write-Info "Listing all students..."
    $listQuery = @{
        statement = 'SELECT META().id, * FROM schoolmgmt.school.students ORDER BY last_name'
    } | ConvertTo-Json
    
    try {
        $response = Invoke-WebRequest -Uri "$QueryUrl/query/service" -Method POST -Body $listQuery -Headers $headers
        $result = ($response.Content | ConvertFrom-Json)
        Write-Success "✅ Found $($result.results.Count) students"
        foreach ($student in $result.results) {
            Write-Host "  - $($student.id): $($student.students.first_name) $($student.students.last_name) (Grade: $($student.students.grade))"
        }
    } catch {
        Write-Error "Failed to list students: $($_.Exception.Message)"
    }
}

function Demo-TeacherCRUD {
    Write-Info "👩‍🏫 Demonstrating Teacher CRUD operations..."
    
    $auth = [Convert]::ToBase64String([Text.Encoding]::ASCII.GetBytes("$($Username):$($Password)"))
    $headers = @{
        Authorization = "Basic $auth"
        "Content-Type" = "application/json"
    }
    
    # CREATE Teacher
    Write-Info "Creating teacher..."
    $createQuery = @{
        statement = 'INSERT INTO schoolmgmt.school.teachers (KEY, VALUE) VALUES ("teacher-001", {
            "id": "teacher-001",
            "first_name": "Dr. Emma",
            "last_name": "Wilson",
            "email": "emma.wilson@school.edu",
            "department": "Mathematics",
            "subjects": ["Algebra", "Geometry"],
            "experience": 8,
            "hire_date": "2020-01-15",
            "status": "active",
            "created_at": "2024-08-05T10:45:00Z"
        })'
    } | ConvertTo-Json
    
    try {
        $response = Invoke-WebRequest -Uri "$QueryUrl/query/service" -Method POST -Body $createQuery -Headers $headers
        Write-Success "✅ Teacher created successfully"
    } catch {
        Write-Warning "Teacher may already exist"
    }
    
    # READ Teacher
    Write-Info "Reading teacher..."
    $readQuery = @{
        statement = 'SELECT * FROM schoolmgmt.school.teachers WHERE department = "Mathematics"'
    } | ConvertTo-Json
    
    try {
        $response = Invoke-WebRequest -Uri "$QueryUrl/query/service" -Method POST -Body $readQuery -Headers $headers
        $result = ($response.Content | ConvertFrom-Json)
        if ($result.results.Count -gt 0) {
            Write-Success "✅ Found $($result.results.Count) Mathematics teachers"
        }
    } catch {
        Write-Error "Failed to read teacher: $($_.Exception.Message)"
    }
}

function Demo-AcademicCRUD {
    Write-Info "📚 Demonstrating Academic Records CRUD operations..."
    
    $auth = [Convert]::ToBase64String([Text.Encoding]::ASCII.GetBytes("$($Username):$($Password)"))
    $headers = @{
        Authorization = "Basic $auth"
        "Content-Type" = "application/json"
    }
    
    # CREATE Academic Record
    Write-Info "Creating academic record..."
    $createQuery = @{
        statement = 'INSERT INTO schoolmgmt.school.academics (KEY, VALUE) VALUES ("academic-001", {
            "id": "academic-001",
            "student_id": "student-001",
            "teacher_id": "teacher-001",
            "subject": "Algebra",
            "grade": "A",
            "semester": "Spring 2024",
            "max_marks": 100,
            "obtained_marks": 95,
            "percentage": 95.0,
            "status": "pass",
            "created_at": "2024-08-05T11:00:00Z"
        })'
    } | ConvertTo-Json
    
    try {
        $response = Invoke-WebRequest -Uri "$QueryUrl/query/service" -Method POST -Body $createQuery -Headers $headers
        Write-Success "✅ Academic record created successfully"
    } catch {
        Write-Warning "Academic record may already exist"
    }
    
    # Query student performance
    Write-Info "Querying student performance..."
    $performanceQuery = @{
        statement = 'SELECT * FROM schoolmgmt.school.academics WHERE student_id = "student-001"'
    } | ConvertTo-Json
    
    try {
        $response = Invoke-WebRequest -Uri "$QueryUrl/query/service" -Method POST -Body $performanceQuery -Headers $headers
        $result = ($response.Content | ConvertFrom-Json)
        Write-Success "✅ Found $($result.results.Count) academic records for student"
        foreach ($record in $result.results) {
            Write-Host "  - $($record.academics.subject): $($record.academics.grade) ($($record.academics.percentage)%)"
        }
    } catch {
        Write-Error "Failed to query academic records: $($_.Exception.Message)"
    }
}

function Demo-AchievementCRUD {
    Write-Info "🏆 Demonstrating Achievement CRUD operations..."
    
    $auth = [Convert]::ToBase64String([Text.Encoding]::ASCII.GetBytes("$($Username):$($Password)"))
    $headers = @{
        Authorization = "Basic $auth"
        "Content-Type" = "application/json"
    }
    
    # CREATE Achievement
    Write-Info "Creating achievement..."
    $createQuery = @{
        statement = 'INSERT INTO schoolmgmt.school.achievements (KEY, VALUE) VALUES ("achievement-001", {
            "id": "achievement-001",
            "student_id": "student-001",
            "title": "Math Excellence Award",
            "description": "Outstanding performance in mathematics",
            "category": "academic",
            "points": 100,
            "date": "2024-04-20",
            "status": "active",
            "created_at": "2024-08-05T11:30:00Z"
        })'
    } | ConvertTo-Json
    
    try {
        $response = Invoke-WebRequest -Uri "$QueryUrl/query/service" -Method POST -Body $createQuery -Headers $headers
        Write-Success "✅ Achievement created successfully"
    } catch {
        Write-Warning "Achievement may already exist"
    }
    
    # Query leaderboard
    Write-Info "Creating leaderboard..."
    $leaderboardQuery = @{
        statement = 'SELECT student_id, SUM(points) as total_points FROM schoolmgmt.school.achievements GROUP BY student_id ORDER BY total_points DESC'
    } | ConvertTo-Json
    
    try {
        $response = Invoke-WebRequest -Uri "$QueryUrl/query/service" -Method POST -Body $leaderboardQuery -Headers $headers
        $result = ($response.Content | ConvertFrom-Json)
        Write-Success "✅ Leaderboard created with $($result.results.Count) students"
        foreach ($entry in $result.results) {
            Write-Host "  - Student $($entry.student_id): $($entry.total_points) points"
        }
    } catch {
        Write-Error "Failed to create leaderboard: $($_.Exception.Message)"
    }
}

function Show-DatabaseStats {
    Write-Info "📊 Showing database statistics..."
    
    $auth = [Convert]::ToBase64String([Text.Encoding]::ASCII.GetBytes("$($Username):$($Password)"))
    $headers = @{
        Authorization = "Basic $auth"
        "Content-Type" = "application/json"
    }
    
    $collections = @("students", "teachers", "academics", "achievements")
    
    foreach ($collection in $collections) {
        $countQuery = @{
            statement = "SELECT COUNT(*) as count FROM schoolmgmt.school.$collection"
        } | ConvertTo-Json
        
        try {
            $response = Invoke-WebRequest -Uri "$QueryUrl/query/service" -Method POST -Body $countQuery -Headers $headers
            $result = ($response.Content | ConvertFrom-Json)
            $count = $result.results[0].count
            Write-Success "📈 $collection`: $count records"
        } catch {
            Write-Warning "Could not get count for $collection"
        }
    }
}

function Show-Menu {
    Write-Host "`n🎓 Couchbase School Management Demo Script" -ForegroundColor Yellow
    Write-Host "==========================================" -ForegroundColor Yellow
    Write-Host "1. Test Connection"
    Write-Host "2. Initialize Cluster"
    Write-Host "3. Create Collections"
    Write-Host "4. Create Indexes"
    Write-Host "5. Demo Student CRUD"
    Write-Host "6. Demo Teacher CRUD"
    Write-Host "7. Demo Academic CRUD"
    Write-Host "8. Demo Achievement CRUD"
    Write-Host "9. Show Database Stats"
    Write-Host "10. Run All Setup (1-4)"
    Write-Host "11. Run All Demos (5-9)"
    Write-Host "12. Full Demo (All)"
    Write-Host "0. Exit"
    Write-Host ""
}

# Main execution
Write-Host "🎓 Couchbase School Management System Demo" -ForegroundColor Green
Write-Host "=============================================" -ForegroundColor Green

if ($Action -eq "all") {
    # Interactive mode
    do {
        Show-Menu
        $choice = Read-Host "Enter your choice (0-12)"
        
        switch ($choice) {
            "1" { Test-CouchbaseConnection }
            "2" { Initialize-Cluster }
            "3" { Create-Collections }
            "4" { Create-Indexes }
            "5" { Demo-StudentCRUD }
            "6" { Demo-TeacherCRUD }
            "7" { Demo-AcademicCRUD }
            "8" { Demo-AchievementCRUD }
            "9" { Show-DatabaseStats }
            "10" { 
                Test-CouchbaseConnection
                Initialize-Cluster
                Create-Collections
                Create-Indexes
                Write-Success "🎉 Setup completed!"
            }
            "11" {
                Demo-StudentCRUD
                Demo-TeacherCRUD
                Demo-AcademicCRUD
                Demo-AchievementCRUD
                Show-DatabaseStats
                Write-Success "🎉 All demos completed!"
            }
            "12" {
                Test-CouchbaseConnection
                Initialize-Cluster
                Create-Collections
                Create-Indexes
                Demo-StudentCRUD
                Demo-TeacherCRUD
                Demo-AcademicCRUD
                Demo-AchievementCRUD
                Show-DatabaseStats
                Write-Success "🎉 Full demo completed!"
            }
            "0" { 
                Write-Success "Goodbye!"
                break
            }
            default { Write-Warning "Invalid choice. Please try again." }
        }
        
        if ($choice -ne "0") {
            Write-Host "`nPress any key to continue..." -ForegroundColor Gray
            $null = $Host.UI.RawUI.ReadKey("NoEcho,IncludeKeyDown")
        }
    } while ($choice -ne "0")
} else {
    # Command line mode
    switch ($Action.ToLower()) {
        "setup" {
            Test-CouchbaseConnection
            Initialize-Cluster
            Create-Collections
            Create-Indexes
        }
        "demo" {
            Demo-StudentCRUD
            Demo-TeacherCRUD
            Demo-AcademicCRUD
            Demo-AchievementCRUD
            Show-DatabaseStats
        }
        "test" { Test-CouchbaseConnection }
        default {
            Write-Host "Usage: .\couchbase-demo.ps1 [-Action setup|demo|test|all]"
            Write-Host "  setup - Initialize cluster and create collections"
            Write-Host "  demo  - Run CRUD demonstrations"
            Write-Host "  test  - Test connection only"
            Write-Host "  all   - Interactive mode (default)"
        }
    }
}
