# Load Sample Data PowerShell Script
# Enhanced data loading script for School Management System on Windows

param(
    [string]$CouchbaseHost = "localhost",
    [int]$CouchbasePort = 8091,
    [string]$CouchbaseUser = "Administrator",
    [string]$CouchbasePassword = "password",
    [string]$BucketName = "schoolmgmt",
    [int]$Timeout = 300
)

# Function to write colored output
function Write-ColorOutput {
    param(
        [string]$Message,
        [string]$Color = "White"
    )
    
    $colorMap = @{
        "Red" = "Red"
        "Green" = "Green"
        "Yellow" = "Yellow"
        "Blue" = "Blue"
        "Cyan" = "Cyan"
        "White" = "White"
    }
    
    Write-Host $Message -ForegroundColor $colorMap[$Color]
}

function Write-Info {
    param([string]$Message)
    Write-ColorOutput "[INFO] $Message" "Blue"
}

function Write-Success {
    param([string]$Message)
    Write-ColorOutput "[SUCCESS] $Message" "Green"
}

function Write-Warning {
    param([string]$Message)
    Write-ColorOutput "[WARNING] $Message" "Yellow"
}

function Write-Error {
    param([string]$Message)
    Write-ColorOutput "[ERROR] $Message" "Red"
}

# Wait for Couchbase to be ready
function Wait-ForCouchbase {
    Write-Info "Waiting for Couchbase to be ready..."
    $count = 0
    
    while ($count -lt $Timeout) {
        try {
            $response = Invoke-WebRequest -Uri "http://${CouchbaseHost}:${CouchbasePort}/pools" -TimeoutSec 5 -ErrorAction Stop
            if ($response.StatusCode -eq 200) {
                Write-Success "Couchbase is ready!"
                return $true
            }
        }
        catch {
            # Continue waiting
        }
        
        Write-Info "Waiting for Couchbase... (${count}s/${Timeout}s)"
        Start-Sleep -Seconds 5
        $count += 5
    }
    
    Write-Error "Couchbase did not become ready within ${Timeout} seconds"
    return $false
}

# Create bucket if it doesn't exist
function New-CouchbaseBucket {
    Write-Info "Creating bucket: $BucketName"
    
    try {
        # Check if bucket exists
        $bucketListUri = "http://${CouchbaseHost}:${CouchbasePort}/pools/default/buckets"
        $auth = [System.Convert]::ToBase64String([System.Text.Encoding]::ASCII.GetBytes("${CouchbaseUser}:${CouchbasePassword}"))
        $headers = @{ Authorization = "Basic $auth" }
        
        $buckets = Invoke-RestMethod -Uri $bucketListUri -Headers $headers -Method Get
        
        $bucketExists = $false
        foreach ($bucket in $buckets) {
            if ($bucket.name -eq $BucketName) {
                $bucketExists = $true
                break
            }
        }
        
        if ($bucketExists) {
            Write-Warning "Bucket $BucketName already exists"
        }
        else {
            # Create bucket
            $createBucketUri = "http://${CouchbaseHost}:${CouchbasePort}/pools/default/buckets"
            $body = @{
                name = $BucketName
                bucketType = "couchbase"
                ramQuotaMB = 256
                replicaNumber = 0
            }
            
            $response = Invoke-RestMethod -Uri $createBucketUri -Headers $headers -Method Post -Body $body
            Write-Success "Bucket $BucketName created successfully"
        }
    }
    catch {
        Write-Error "Failed to create bucket: $($_.Exception.Message)"
        return $false
    }
    
    return $true
}

# Load sample data using N1QL
function Import-SampleData {
    Write-Info "Loading sample data into $BucketName..."
    
    # Create temporary directory for data files
    $tempDir = New-TemporaryFile | Split-Path
    $dataDir = Join-Path $tempDir "sms_data"
    New-Item -ItemType Directory -Path $dataDir -Force | Out-Null
    
    Write-Info "Using temporary directory: $dataDir"
    
    # Sample data arrays
    $students = @(
        @{
            id = "student-001"
            type = "student"
            firstName = "John"
            lastName = "Doe"
            email = "john.doe@school.edu"
            dateOfBirth = "2005-03-15"
            grade = "10"
            enrollmentDate = "2023-09-01"
            status = "active"
            address = @{
                street = "123 Main St"
                city = "Springfield"
                state = "IL"
                zipCode = "62701"
            }
            parentContact = @{
                name = "Jane Doe"
                phone = "555-0101"
                email = "jane.doe@email.com"
            }
        },
        @{
            id = "student-002"
            type = "student"
            firstName = "Emily"
            lastName = "Smith"
            email = "emily.smith@school.edu"
            dateOfBirth = "2005-07-22"
            grade = "10"
            enrollmentDate = "2023-09-01"
            status = "active"
            address = @{
                street = "456 Oak Ave"
                city = "Springfield"
                state = "IL"
                zipCode = "62702"
            }
            parentContact = @{
                name = "Robert Smith"
                phone = "555-0102"
                email = "robert.smith@email.com"
            }
        }
    )
    
    $teachers = @(
        @{
            id = "teacher-001"
            type = "teacher"
            firstName = "Dr. Robert"
            lastName = "Anderson"
            email = "robert.anderson@school.edu"
            department = "Mathematics"
            subject = "Algebra"
            hireDate = "2015-08-15"
            status = "active"
            qualifications = @("PhD Mathematics", "MEd Education")
            officeHours = "Mon-Wed-Fri 2:00-4:00 PM"
            phone = "555-1001"
        },
        @{
            id = "teacher-002"
            type = "teacher"
            firstName = "Ms. Jennifer"
            lastName = "Davis"
            email = "jennifer.davis@school.edu"
            department = "English"
            subject = "Literature"
            hireDate = "2018-01-10"
            status = "active"
            qualifications = @("MA English Literature", "BA English")
            officeHours = "Tue-Thu 1:00-3:00 PM"
            phone = "555-1002"
        }
    )
    
    $courses = @(
        @{
            id = "course-001"
            type = "course"
            courseCode = "MATH101"
            courseName = "Algebra I"
            department = "Mathematics"
            teacherId = "teacher-001"
            credits = 1.0
            semester = "Fall 2024"
            capacity = 25
            enrolledStudents = @("student-001", "student-002")
            schedule = @{
                days = @("Monday", "Wednesday", "Friday")
                time = "09:00-10:00"
                room = "Math-101"
            }
        }
    )
    
    $grades = @(
        @{
            id = "grade-001"
            type = "grade"
            studentId = "student-001"
            courseId = "course-001"
            teacherId = "teacher-001"
            assignmentName = "Midterm Exam"
            grade = "B+"
            points = 87
            maxPoints = 100
            dateGraded = "2024-10-15"
            semester = "Fall 2024"
            category = "exam"
        }
    )
    
    $achievements = @(
        @{
            id = "achievement-001"
            type = "achievement"
            studentId = "student-002"
            title = "Honor Roll"
            description = "Achieved GPA above 3.5 for Fall 2024 semester"
            category = "academic"
            dateAwarded = "2024-11-01"
            points = 100
            level = "semester"
            metadata = @{
                gpa = 3.8
                semester = "Fall 2024"
            }
        }
    )
    
    # Combine all data
    $allData = @{
        students = $students
        teachers = $teachers
        courses = $courses
        grades = $grades
        achievements = $achievements
    }
    
    # Load data using N1QL
    $n1qlUri = "http://${CouchbaseHost}:8093/query"
    $auth = [System.Convert]::ToBase64String([System.Text.Encoding]::ASCII.GetBytes("${CouchbaseUser}:${CouchbasePassword}"))
    $headers = @{ 
        Authorization = "Basic $auth"
        "Content-Type" = "application/json"
    }
    
    foreach ($dataType in $allData.Keys) {
        Write-Info "Loading $dataType data using N1QL..."
        
        foreach ($item in $allData[$dataType]) {
            try {
                $jsonData = $item | ConvertTo-Json -Depth 10 -Compress
                $statement = "UPSERT INTO ``$BucketName`` (KEY, VALUE) VALUES ('$($item.id)', $jsonData)"
                
                $body = @{
                    statement = $statement
                } | ConvertTo-Json
                
                $response = Invoke-RestMethod -Uri $n1qlUri -Headers $headers -Method Post -Body $body
                
                if ($response.status -eq "success") {
                    Write-Info "Inserted record: $($item.id)"
                }
                else {
                    Write-Warning "Failed to insert record: $($item.id)"
                }
            }
            catch {
                Write-Warning "Failed to insert record $($item.id): $($_.Exception.Message)"
            }
        }
    }
    
    # Clean up temporary files
    Remove-Item -Path $dataDir -Recurse -Force -ErrorAction SilentlyContinue
    Write-Success "Sample data loaded successfully!"
}

# Create indexes for better query performance
function New-CouchbaseIndexes {
    Write-Info "Creating indexes for better performance..."
    
    $indexes = @(
        "CREATE INDEX idx_student_type ON ``$BucketName``(type) WHERE type = 'student'",
        "CREATE INDEX idx_teacher_type ON ``$BucketName``(type) WHERE type = 'teacher'",
        "CREATE INDEX idx_course_type ON ``$BucketName``(type) WHERE type = 'course'",
        "CREATE INDEX idx_grade_type ON ``$BucketName``(type) WHERE type = 'grade'",
        "CREATE INDEX idx_achievement_type ON ``$BucketName``(type) WHERE type = 'achievement'",
        "CREATE INDEX idx_student_email ON ``$BucketName``(email) WHERE type = 'student'",
        "CREATE INDEX idx_teacher_email ON ``$BucketName``(email) WHERE type = 'teacher'",
        "CREATE INDEX idx_grade_student ON ``$BucketName``(studentId) WHERE type = 'grade'",
        "CREATE INDEX idx_achievement_student ON ``$BucketName``(studentId) WHERE type = 'achievement'",
        "CREATE INDEX idx_course_teacher ON ``$BucketName``(teacherId) WHERE type = 'course'"
    )
    
    $n1qlUri = "http://${CouchbaseHost}:8093/query"
    $auth = [System.Convert]::ToBase64String([System.Text.Encoding]::ASCII.GetBytes("${CouchbaseUser}:${CouchbasePassword}"))
    $headers = @{ 
        Authorization = "Basic $auth"
        "Content-Type" = "application/json"
    }
    
    foreach ($index in $indexes) {
        try {
            $body = @{
                statement = $index
            } | ConvertTo-Json
            
            $response = Invoke-RestMethod -Uri $n1qlUri -Headers $headers -Method Post -Body $body
            Write-Info "Created index successfully"
        }
        catch {
            Write-Warning "Failed to create index: $($_.Exception.Message)"
        }
    }
    
    Write-Success "Indexes created successfully!"
}

# Verify data was loaded correctly
function Test-DataLoad {
    Write-Info "Verifying data was loaded correctly..."
    
    $types = @("student", "teacher", "course", "grade", "achievement")
    
    $n1qlUri = "http://${CouchbaseHost}:8093/query"
    $auth = [System.Convert]::ToBase64String([System.Text.Encoding]::ASCII.GetBytes("${CouchbaseUser}:${CouchbasePassword}"))
    $headers = @{ 
        Authorization = "Basic $auth"
        "Content-Type" = "application/json"
    }
    
    foreach ($type in $types) {
        try {
            $statement = "SELECT COUNT(*) as count FROM ``$BucketName`` WHERE type = '$type'"
            $body = @{
                statement = $statement
            } | ConvertTo-Json
            
            $response = Invoke-RestMethod -Uri $n1qlUri -Headers $headers -Method Post -Body $body
            
            if ($response.status -eq "success" -and $response.results) {
                $count = $response.results[0].count
                if ($count -gt 0) {
                    Write-Success "$type`: $count records loaded"
                }
                else {
                    Write-Warning "$type`: No records found"
                }
            }
        }
        catch {
            Write-Warning "$type`: Could not verify record count"
        }
    }
}

# Print summary and next steps
function Show-Summary {
    Write-Success "Data loading completed successfully!"
    Write-Host ""
    Write-Host "Summary of loaded data:"
    Write-Host "- Students: 2 records"
    Write-Host "- Teachers: 2 records"
    Write-Host "- Courses: 1 record"
    Write-Host "- Grades: 1 record"
    Write-Host "- Achievements: 1 record"
    Write-Host ""
    Write-Host "You can now:"
    Write-Host "1. Access Couchbase Web Console at: http://${CouchbaseHost}:${CouchbasePort}"
    Write-Host "2. Use the REST API endpoints to query the data"
    Write-Host "3. Run sample cURL commands to test the services"
    Write-Host ""
    Write-Host "Next steps:"
    Write-Host "- Run: curl http://localhost:8080/api/students"
    Write-Host "- Run: curl http://localhost:8080/api/teachers"
    Write-Host "- Check the API documentation for more endpoints"
}

# Main execution
function Main {
    Write-Info "Starting data loading process for School Management System..."
    
    try {
        # Run all steps
        if (-not (Wait-ForCouchbase)) {
            return
        }
        
        if (-not (New-CouchbaseBucket)) {
            return
        }
        
        Import-SampleData
        New-CouchbaseIndexes
        Test-DataLoad
        Show-Summary
        
        Write-Success "Data loading script completed successfully!"
    }
    catch {
        Write-Error "Script failed: $($_.Exception.Message)"
    }
}

# Run main function
Main
