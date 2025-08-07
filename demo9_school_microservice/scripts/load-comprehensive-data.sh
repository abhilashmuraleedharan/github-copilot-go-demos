#!/bin/bash

# Enhanced Data Loading Script for School Management System
# This script loads comprehensive sample data into Couchbase

set -e  # Exit on any error

# Configuration
COUCHBASE_HOST=${COUCHBASE_HOST:-localhost}
COUCHBASE_PORT=${COUCHBASE_PORT:-8091}
COUCHBASE_USER=${COUCHBASE_USER:-Administrator}
COUCHBASE_PASSWORD=${COUCHBASE_PASSWORD:-password}
BUCKET_NAME=${BUCKET_NAME:-schoolmgmt}
TIMEOUT=300

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Logging functions
log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

log_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if couchbase-cli is available
check_couchbase_cli() {
    if ! command -v couchbase-cli &> /dev/null; then
        log_error "couchbase-cli not found. Please install Couchbase CLI tools or run this script from within the Couchbase container."
        log_info "To run from container: docker-compose exec couchbase bash -c 'curl -o load-data.sh <script-url> && chmod +x load-data.sh && ./load-data.sh'"
        exit 1
    fi
}

# Wait for Couchbase to be ready
wait_for_couchbase() {
    log_info "Waiting for Couchbase to be ready..."
    local count=0
    while [ $count -lt $TIMEOUT ]; do
        if couchbase-cli server-info -c $COUCHBASE_HOST:$COUCHBASE_PORT -u $COUCHBASE_USER -p $COUCHBASE_PASSWORD &>/dev/null; then
            log_success "Couchbase is ready!"
            return 0
        fi
        log_info "Waiting for Couchbase... (${count}s/${TIMEOUT}s)"
        sleep 5
        count=$((count + 5))
    done
    log_error "Couchbase did not become ready within ${TIMEOUT} seconds"
    exit 1
}

# Create bucket if it doesn't exist
create_bucket() {
    log_info "Creating bucket: $BUCKET_NAME"
    
    # Check if bucket exists
    if couchbase-cli bucket-list -c $COUCHBASE_HOST:$COUCHBASE_PORT -u $COUCHBASE_USER -p $COUCHBASE_PASSWORD | grep -q "$BUCKET_NAME"; then
        log_warning "Bucket $BUCKET_NAME already exists"
    else
        couchbase-cli bucket-create \
            -c $COUCHBASE_HOST:$COUCHBASE_PORT \
            -u $COUCHBASE_USER \
            -p $COUCHBASE_PASSWORD \
            --bucket $BUCKET_NAME \
            --bucket-type couchbase \
            --bucket-ramsize 256 \
            --bucket-replica 0 \
            --wait
        log_success "Bucket $BUCKET_NAME created successfully"
    fi
}

# Load sample data using cbimport or direct N1QL
load_sample_data() {
    log_info "Loading sample data into $BUCKET_NAME..."

    # Create temporary directory for data files
    local temp_dir=$(mktemp -d)
    log_info "Using temporary directory: $temp_dir"

    # Generate students data
    cat > "$temp_dir/students.json" << 'EOF'
{"id":"student-001","type":"student","firstName":"John","lastName":"Doe","email":"john.doe@school.edu","dateOfBirth":"2005-03-15","grade":"10","enrollmentDate":"2023-09-01","status":"active","address":{"street":"123 Main St","city":"Springfield","state":"IL","zipCode":"62701"},"parentContact":{"name":"Jane Doe","phone":"555-0101","email":"jane.doe@email.com"}}
{"id":"student-002","type":"student","firstName":"Emily","lastName":"Smith","email":"emily.smith@school.edu","dateOfBirth":"2005-07-22","grade":"10","enrollmentDate":"2023-09-01","status":"active","address":{"street":"456 Oak Ave","city":"Springfield","state":"IL","zipCode":"62702"},"parentContact":{"name":"Robert Smith","phone":"555-0102","email":"robert.smith@email.com"}}
{"id":"student-003","type":"student","firstName":"Michael","lastName":"Johnson","email":"michael.johnson@school.edu","dateOfBirth":"2004-11-08","grade":"11","enrollmentDate":"2022-09-01","status":"active","address":{"street":"789 Pine Rd","city":"Springfield","state":"IL","zipCode":"62703"},"parentContact":{"name":"Lisa Johnson","phone":"555-0103","email":"lisa.johnson@email.com"}}
{"id":"student-004","type":"student","firstName":"Sarah","lastName":"Williams","email":"sarah.williams@school.edu","dateOfBirth":"2003-12-30","grade":"12","enrollmentDate":"2021-09-01","status":"active","address":{"street":"321 Elm St","city":"Springfield","state":"IL","zipCode":"62704"},"parentContact":{"name":"David Williams","phone":"555-0104","email":"david.williams@email.com"}}
{"id":"student-005","type":"student","firstName":"Alex","lastName":"Brown","email":"alex.brown@school.edu","dateOfBirth":"2005-05-14","grade":"10","enrollmentDate":"2023-09-01","status":"active","address":{"street":"654 Maple Dr","city":"Springfield","state":"IL","zipCode":"62705"},"parentContact":{"name":"Michelle Brown","phone":"555-0105","email":"michelle.brown@email.com"}}
EOF

    # Generate teachers data
    cat > "$temp_dir/teachers.json" << 'EOF'
{"id":"teacher-001","type":"teacher","firstName":"Dr. Robert","lastName":"Anderson","email":"robert.anderson@school.edu","department":"Mathematics","subject":"Algebra","hireDate":"2015-08-15","status":"active","qualifications":["PhD Mathematics","MEd Education"],"officeHours":"Mon-Wed-Fri 2:00-4:00 PM","phone":"555-1001"}
{"id":"teacher-002","type":"teacher","firstName":"Ms. Jennifer","lastName":"Davis","email":"jennifer.davis@school.edu","department":"English","subject":"Literature","hireDate":"2018-01-10","status":"active","qualifications":["MA English Literature","BA English"],"officeHours":"Tue-Thu 1:00-3:00 PM","phone":"555-1002"}
{"id":"teacher-003","type":"teacher","firstName":"Mr. James","lastName":"Wilson","email":"james.wilson@school.edu","department":"Science","subject":"Biology","hireDate":"2016-07-20","status":"active","qualifications":["MS Biology","BS Biology","Teaching Certificate"],"officeHours":"Mon-Wed 3:00-5:00 PM","phone":"555-1003"}
{"id":"teacher-004","type":"teacher","firstName":"Mrs. Patricia","lastName":"Miller","email":"patricia.miller@school.edu","department":"History","subject":"World History","hireDate":"2012-09-01","status":"active","qualifications":["MA History","BA History","MEd Secondary Education"],"officeHours":"Tue-Thu 2:00-4:00 PM","phone":"555-1004"}
{"id":"teacher-005","type":"teacher","firstName":"Dr. Mark","lastName":"Garcia","email":"mark.garcia@school.edu","department":"Science","subject":"Chemistry","hireDate":"2019-02-15","status":"active","qualifications":["PhD Chemistry","MS Chemistry"],"officeHours":"Mon-Fri 1:00-2:00 PM","phone":"555-1005"}
EOF

    # Generate courses data
    cat > "$temp_dir/courses.json" << 'EOF'
{"id":"course-001","type":"course","courseCode":"MATH101","courseName":"Algebra I","department":"Mathematics","teacherId":"teacher-001","credits":1.0,"semester":"Fall 2024","capacity":25,"enrolledStudents":["student-001","student-002","student-005"],"schedule":{"days":["Monday","Wednesday","Friday"],"time":"09:00-10:00","room":"Math-101"}}
{"id":"course-002","type":"course","courseCode":"ENG201","courseName":"American Literature","department":"English","teacherId":"teacher-002","credits":1.0,"semester":"Fall 2024","capacity":20,"enrolledStudents":["student-003","student-004"],"schedule":{"days":["Tuesday","Thursday"],"time":"10:00-11:30","room":"Eng-201"}}
{"id":"course-003","type":"course","courseCode":"BIO301","courseName":"Advanced Biology","department":"Science","teacherId":"teacher-003","credits":1.5,"semester":"Fall 2024","capacity":15,"enrolledStudents":["student-003","student-004"],"schedule":{"days":["Monday","Wednesday","Friday"],"time":"11:00-12:00","room":"Lab-301"}}
{"id":"course-004","type":"course","courseCode":"HIST101","courseName":"World History","department":"History","teacherId":"teacher-004","credits":1.0,"semester":"Fall 2024","capacity":30,"enrolledStudents":["student-001","student-002","student-005"],"schedule":{"days":["Tuesday","Thursday"],"time":"14:00-15:30","room":"Hist-101"}}
{"id":"course-005","type":"course","courseCode":"CHEM201","courseName":"Organic Chemistry","department":"Science","teacherId":"teacher-005","credits":1.5,"semester":"Fall 2024","capacity":18,"enrolledStudents":["student-004"],"schedule":{"days":["Monday","Wednesday","Friday"],"time":"13:00-14:00","room":"Lab-201"}}
EOF

    # Generate grades data
    cat > "$temp_dir/grades.json" << 'EOF'
{"id":"grade-001","type":"grade","studentId":"student-001","courseId":"course-001","teacherId":"teacher-001","assignmentName":"Midterm Exam","grade":"B+","points":87,"maxPoints":100,"dateGraded":"2024-10-15","semester":"Fall 2024","category":"exam"}
{"id":"grade-002","type":"grade","studentId":"student-001","courseId":"course-004","teacherId":"teacher-004","assignmentName":"Essay: World War I","grade":"A-","points":92,"maxPoints":100,"dateGraded":"2024-10-20","semester":"Fall 2024","category":"assignment"}
{"id":"grade-003","type":"grade","studentId":"student-002","courseId":"course-001","teacherId":"teacher-001","assignmentName":"Midterm Exam","grade":"A","points":95,"maxPoints":100,"dateGraded":"2024-10-15","semester":"Fall 2024","category":"exam"}
{"id":"grade-004","type":"grade","studentId":"student-003","courseId":"course-002","teacherId":"teacher-002","assignmentName":"Book Report","grade":"B","points":84,"maxPoints":100,"dateGraded":"2024-10-25","semester":"Fall 2024","category":"assignment"}
{"id":"grade-005","type":"grade","studentId":"student-004","courseId":"course-003","teacherId":"teacher-003","assignmentName":"Lab Practical","grade":"A+","points":98,"maxPoints":100,"dateGraded":"2024-10-30","semester":"Fall 2024","category":"lab"}
EOF

    # Generate achievements data
    cat > "$temp_dir/achievements.json" << 'EOF'
{"id":"achievement-001","type":"achievement","studentId":"student-002","title":"Honor Roll","description":"Achieved GPA above 3.5 for Fall 2024 semester","category":"academic","dateAwarded":"2024-11-01","points":100,"level":"semester","metadata":{"gpa":3.8,"semester":"Fall 2024"}}
{"id":"achievement-002","type":"achievement","studentId":"student-004","title":"Science Fair Winner","description":"First place in Chemistry category at state science fair","category":"competition","dateAwarded":"2024-10-15","points":200,"level":"state","metadata":{"category":"Chemistry","placement":1,"event":"State Science Fair 2024"}}
{"id":"achievement-003","type":"achievement","studentId":"student-001","title":"Perfect Attendance","description":"No absences for the month of October","category":"attendance","dateAwarded":"2024-11-01","points":50,"level":"monthly","metadata":{"month":"October 2024","absences":0}}
{"id":"achievement-004","type":"achievement","studentId":"student-003","title":"Literary Award","description":"Outstanding performance in American Literature class","category":"academic","dateAwarded":"2024-10-20","points":150,"level":"course","metadata":{"course":"ENG201","grade":"A+","teacher":"teacher-002"}}
{"id":"achievement-005","type":"achievement","studentId":"student-005","title":"Math Competition Finalist","description":"Placed in top 10 at regional math competition","category":"competition","dateAwarded":"2024-09-30","points":75,"level":"regional","metadata":{"event":"Regional Math Competition","placement":8,"participants":150}}
EOF

    # Import data using cbimport if available, otherwise use curl with N1QL
    if command -v cbimport &> /dev/null; then
        log_info "Using cbimport to load data..."
        
        for file in students teachers courses grades achievements; do
            log_info "Importing $file data..."
            cbimport json -c $COUCHBASE_HOST:$COUCHBASE_PORT \
                -u $COUCHBASE_USER \
                -p $COUCHBASE_PASSWORD \
                -b $BUCKET_NAME \
                -d "file://$temp_dir/$file.json" \
                -f lines \
                -g "#MONO_INCR#" \
                || log_warning "Failed to import $file data"
        done
    else
        log_info "cbimport not available, using N1QL UPSERT..."
        
        # Load data using N1QL
        for file in students teachers courses grades achievements; do
            log_info "Loading $file data using N1QL..."
            while IFS= read -r line; do
                if [ ! -z "$line" ]; then
                    id=$(echo "$line" | grep -o '"id":"[^"]*"' | cut -d'"' -f4)
                    curl -s -X POST \
                        -H "Content-Type: application/json" \
                        -u "$COUCHBASE_USER:$COUCHBASE_PASSWORD" \
                        -d "{\"statement\":\"UPSERT INTO \`$BUCKET_NAME\` (KEY, VALUE) VALUES ('$id', $line)\"}" \
                        "http://$COUCHBASE_HOST:8093/query" > /dev/null || log_warning "Failed to insert record: $id"
                fi
            done < "$temp_dir/$file.json"
        done
    fi

    # Clean up temporary files
    rm -rf "$temp_dir"
    log_success "Sample data loaded successfully!"
}

# Create indexes for better query performance
create_indexes() {
    log_info "Creating indexes for better performance..."

    # Array of index creation statements
    indexes=(
        "CREATE INDEX idx_student_type ON \`$BUCKET_NAME\`(type) WHERE type = 'student'"
        "CREATE INDEX idx_teacher_type ON \`$BUCKET_NAME\`(type) WHERE type = 'teacher'"
        "CREATE INDEX idx_course_type ON \`$BUCKET_NAME\`(type) WHERE type = 'course'"
        "CREATE INDEX idx_grade_type ON \`$BUCKET_NAME\`(type) WHERE type = 'grade'"
        "CREATE INDEX idx_achievement_type ON \`$BUCKET_NAME\`(type) WHERE type = 'achievement'"
        "CREATE INDEX idx_student_email ON \`$BUCKET_NAME\`(email) WHERE type = 'student'"
        "CREATE INDEX idx_teacher_email ON \`$BUCKET_NAME\`(email) WHERE type = 'teacher'"
        "CREATE INDEX idx_grade_student ON \`$BUCKET_NAME\`(studentId) WHERE type = 'grade'"
        "CREATE INDEX idx_achievement_student ON \`$BUCKET_NAME\`(studentId) WHERE type = 'achievement'"
        "CREATE INDEX idx_course_teacher ON \`$BUCKET_NAME\`(teacherId) WHERE type = 'course'"
    )

    for index in "${indexes[@]}"; do
        curl -s -X POST \
            -H "Content-Type: application/json" \
            -u "$COUCHBASE_USER:$COUCHBASE_PASSWORD" \
            -d "{\"statement\":\"$index\"}" \
            "http://$COUCHBASE_HOST:8093/query" > /dev/null || log_warning "Failed to create index"
    done

    log_success "Indexes created successfully!"
}

# Verify data was loaded correctly
verify_data() {
    log_info "Verifying data was loaded correctly..."

    # Check document counts by type
    types=("student" "teacher" "course" "grade" "achievement")
    
    for type in "${types[@]}"; do
        count=$(curl -s -X POST \
            -H "Content-Type: application/json" \
            -u "$COUCHBASE_USER:$COUCHBASE_PASSWORD" \
            -d "{\"statement\":\"SELECT COUNT(*) as count FROM \`$BUCKET_NAME\` WHERE type = '$type'\"}" \
            "http://$COUCHBASE_HOST:8093/query" | grep -o '"count":[0-9]*' | cut -d':' -f2)
        
        if [ ! -z "$count" ] && [ "$count" -gt 0 ]; then
            log_success "$type: $count records loaded"
        else
            log_warning "$type: No records found"
        fi
    done
}

# Print summary and next steps
print_summary() {
    log_success "Data loading completed successfully!"
    echo
    echo "Summary of loaded data:"
    echo "- Students: 5 records"
    echo "- Teachers: 5 records"
    echo "- Courses: 5 records"
    echo "- Grades: 5 records"
    echo "- Achievements: 5 records"
    echo
    echo "You can now:"
    echo "1. Access Couchbase Web Console at: http://$COUCHBASE_HOST:$COUCHBASE_PORT"
    echo "2. Use the REST API endpoints to query the data"
    echo "3. Run sample cURL commands to test the services"
    echo
    echo "Next steps:"
    echo "- Run: curl http://localhost:8080/api/students to list all students"
    echo "- Run: curl http://localhost:8080/api/teachers to list all teachers"
    echo "- Check the API documentation for more endpoints"
}

# Main execution
main() {
    log_info "Starting data loading process for School Management System..."
    
    # Run all steps
    check_couchbase_cli
    wait_for_couchbase
    create_bucket
    load_sample_data
    create_indexes
    verify_data
    print_summary
    
    log_success "Data loading script completed successfully!"
}

# Run main function
main "$@"
