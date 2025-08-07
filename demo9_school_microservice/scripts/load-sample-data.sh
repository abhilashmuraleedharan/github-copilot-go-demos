#!/bin/bash

# School Management System - Sample Data Loader
# This script loads sample data into all collections for testing and demonstration

# Configuration
COUCHBASE_URL="${COUCHBASE_URL:-http://localhost:8091}"
QUERY_URL="${QUERY_URL:-http://localhost:8093}"
USERNAME="${USERNAME:-Administrator}"
PASSWORD="${PASSWORD:-password123}"

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

# Helper functions
log_success() { echo -e "${GREEN}✅ $1${NC}"; }
log_error() { echo -e "${RED}❌ $1${NC}"; }
log_info() { echo -e "${CYAN}ℹ️  $1${NC}"; }
log_warning() { echo -e "${YELLOW}⚠️  $1${NC}"; }

echo -e "${GREEN}📊 Loading Sample Data into School Management System${NC}"
echo -e "${GREEN}=================================================${NC}"

# Load Students
log_info "Loading sample students..."

students=(
    '{"id": "student-001", "first_name": "John", "last_name": "Doe", "email": "john.doe@school.edu", "grade": "10", "age": 16, "enrollment_date": "2024-08-01", "status": "active", "created_at": "2024-08-05T10:30:00Z"}'
    '{"id": "student-002", "first_name": "Jane", "last_name": "Smith", "email": "jane.smith@school.edu", "grade": "11", "age": 17, "enrollment_date": "2023-08-01", "status": "active", "created_at": "2024-08-05T10:31:00Z"}'
    '{"id": "student-003", "first_name": "Michael", "last_name": "Johnson", "email": "michael.johnson@school.edu", "grade": "12", "age": 18, "enrollment_date": "2022-08-01", "status": "active", "created_at": "2024-08-05T10:32:00Z"}'
    '{"id": "student-004", "first_name": "Emily", "last_name": "Davis", "email": "emily.davis@school.edu", "grade": "9", "age": 15, "enrollment_date": "2024-08-01", "status": "active", "created_at": "2024-08-05T10:33:00Z"}'
    '{"id": "student-005", "first_name": "David", "last_name": "Wilson", "email": "david.wilson@school.edu", "grade": "10", "age": 16, "enrollment_date": "2024-08-01", "status": "active", "created_at": "2024-08-05T10:34:00Z"}'
)

counter=1
for student in "${students[@]}"; do
    id="student-$(printf "%03d" $counter)"
    curl -s -X POST "$QUERY_URL/query/service" \
        -u "$USERNAME:$PASSWORD" \
        -H "Content-Type: application/json" \
        -d "{\"statement\": \"INSERT INTO schoolmgmt.school.students (KEY, VALUE) VALUES (\\\"$id\\\", $student)\"}" > /dev/null
    
    if [ $? -eq 0 ]; then
        name=$(echo "$student" | grep -o '"first_name": "[^"]*"' | cut -d '"' -f 4)
        surname=$(echo "$student" | grep -o '"last_name": "[^"]*"' | cut -d '"' -f 4)
        log_success "Created student: $name $surname ($id)"
    else
        log_warning "Failed to create student $id"
    fi
    ((counter++))
done

# Load Teachers
log_info "Loading sample teachers..."

teachers=(
    '{"id": "teacher-001", "first_name": "Dr. Emma", "last_name": "Wilson", "email": "emma.wilson@school.edu", "department": "Mathematics", "subjects": ["Algebra", "Geometry", "Calculus"], "experience": 8, "hire_date": "2020-01-15", "status": "active", "created_at": "2024-08-05T10:45:00Z"}'
    '{"id": "teacher-002", "first_name": "Prof. Robert", "last_name": "Brown", "email": "robert.brown@school.edu", "department": "Science", "subjects": ["Physics", "Chemistry"], "experience": 12, "hire_date": "2018-08-20", "status": "active", "created_at": "2024-08-05T10:46:00Z"}'
    '{"id": "teacher-003", "first_name": "Ms. Sarah", "last_name": "Johnson", "email": "sarah.johnson@school.edu", "department": "English", "subjects": ["Literature", "Writing"], "experience": 6, "hire_date": "2021-09-01", "status": "active", "created_at": "2024-08-05T10:47:00Z"}'
    '{"id": "teacher-004", "first_name": "Mr. James", "last_name": "Davis", "email": "james.davis@school.edu", "department": "History", "subjects": ["World History", "American History"], "experience": 10, "hire_date": "2019-01-10", "status": "active", "created_at": "2024-08-05T10:48:00Z"}'
    '{"id": "teacher-005", "first_name": "Dr. Lisa", "last_name": "Garcia", "email": "lisa.garcia@school.edu", "department": "Science", "subjects": ["Biology", "Environmental Science"], "experience": 9, "hire_date": "2020-08-15", "status": "active", "created_at": "2024-08-05T10:49:00Z"}'
)

counter=1
for teacher in "${teachers[@]}"; do
    id="teacher-$(printf "%03d" $counter)"
    curl -s -X POST "$QUERY_URL/query/service" \
        -u "$USERNAME:$PASSWORD" \
        -H "Content-Type: application/json" \
        -d "{\"statement\": \"INSERT INTO schoolmgmt.school.teachers (KEY, VALUE) VALUES (\\\"$id\\\", $teacher)\"}" > /dev/null
    
    if [ $? -eq 0 ]; then
        name=$(echo "$teacher" | grep -o '"first_name": "[^"]*"' | cut -d '"' -f 4)
        surname=$(echo "$teacher" | grep -o '"last_name": "[^"]*"' | cut -d '"' -f 4)
        dept=$(echo "$teacher" | grep -o '"department": "[^"]*"' | cut -d '"' -f 4)
        log_success "Created teacher: $name $surname - $dept ($id)"
    else
        log_warning "Failed to create teacher $id"
    fi
    ((counter++))
done

# Load Academic Records
log_info "Loading sample academic records..."

academics=(
    '{"id": "academic-001", "student_id": "student-001", "teacher_id": "teacher-001", "subject": "Algebra", "grade": "A", "semester": "Spring 2024", "academic_year": "2023-2024", "max_marks": 100, "obtained_marks": 95, "percentage": 95.0, "status": "pass", "created_at": "2024-08-05T11:00:00Z"}'
    '{"id": "academic-002", "student_id": "student-001", "teacher_id": "teacher-002", "subject": "Physics", "grade": "B+", "semester": "Spring 2024", "academic_year": "2023-2024", "max_marks": 100, "obtained_marks": 87, "percentage": 87.0, "status": "pass", "created_at": "2024-08-05T11:01:00Z"}'
    '{"id": "academic-003", "student_id": "student-002", "teacher_id": "teacher-001", "subject": "Geometry", "grade": "A-", "semester": "Spring 2024", "academic_year": "2023-2024", "max_marks": 100, "obtained_marks": 92, "percentage": 92.0, "status": "pass", "created_at": "2024-08-05T11:02:00Z"}'
    '{"id": "academic-004", "student_id": "student-002", "teacher_id": "teacher-003", "subject": "Literature", "grade": "A", "semester": "Spring 2024", "academic_year": "2023-2024", "max_marks": 100, "obtained_marks": 96, "percentage": 96.0, "status": "pass", "created_at": "2024-08-05T11:03:00Z"}'
    '{"id": "academic-005", "student_id": "student-003", "teacher_id": "teacher-004", "subject": "World History", "grade": "B", "semester": "Spring 2024", "academic_year": "2023-2024", "max_marks": 100, "obtained_marks": 85, "percentage": 85.0, "status": "pass", "created_at": "2024-08-05T11:04:00Z"}'
    '{"id": "academic-006", "student_id": "student-004", "teacher_id": "teacher-005", "subject": "Biology", "grade": "A+", "semester": "Spring 2024", "academic_year": "2023-2024", "max_marks": 100, "obtained_marks": 98, "percentage": 98.0, "status": "pass", "created_at": "2024-08-05T11:05:00Z"}'
    '{"id": "academic-007", "student_id": "student-005", "teacher_id": "teacher-001", "subject": "Algebra", "grade": "B+", "semester": "Spring 2024", "academic_year": "2023-2024", "max_marks": 100, "obtained_marks": 88, "percentage": 88.0, "status": "pass", "created_at": "2024-08-05T11:06:00Z"}'
)

counter=1
for academic in "${academics[@]}"; do
    id="academic-$(printf "%03d" $counter)"
    curl -s -X POST "$QUERY_URL/query/service" \
        -u "$USERNAME:$PASSWORD" \
        -H "Content-Type: application/json" \
        -d "{\"statement\": \"INSERT INTO schoolmgmt.school.academics (KEY, VALUE) VALUES (\\\"$id\\\", $academic)\"}" > /dev/null
    
    if [ $? -eq 0 ]; then
        subject=$(echo "$academic" | grep -o '"subject": "[^"]*"' | cut -d '"' -f 4)
        grade=$(echo "$academic" | grep -o '"grade": "[^"]*"' | cut -d '"' -f 4)
        student_id=$(echo "$academic" | grep -o '"student_id": "[^"]*"' | cut -d '"' -f 4)
        log_success "Created academic record: $student_id - $subject ($grade)"
    else
        log_warning "Failed to create academic record $id"
    fi
    ((counter++))
done

# Load Achievements
log_info "Loading sample achievements..."

achievements=(
    '{"id": "achievement-001", "student_id": "student-001", "title": "Math Excellence Award", "description": "Outstanding performance in mathematics", "category": "academic", "points": 100, "date": "2024-04-20", "status": "active", "created_at": "2024-08-05T11:30:00Z"}'
    '{"id": "achievement-002", "student_id": "student-002", "title": "Honor Roll", "description": "Achieved honor roll status for Spring 2024", "category": "academic", "points": 75, "date": "2024-05-15", "status": "active", "created_at": "2024-08-05T11:31:00Z"}'
    '{"id": "achievement-003", "student_id": "student-003", "title": "Leadership Award", "description": "Exceptional leadership in student council", "category": "leadership", "points": 80, "date": "2024-03-10", "status": "active", "created_at": "2024-08-05T11:32:00Z"}'
    '{"id": "achievement-004", "student_id": "student-001", "title": "Science Fair Winner", "description": "First place in regional science fair", "category": "science", "points": 120, "date": "2024-02-28", "status": "active", "created_at": "2024-08-05T11:33:00Z"}'
    '{"id": "achievement-005", "student_id": "student-004", "title": "Perfect Attendance", "description": "Perfect attendance for the academic year", "category": "attendance", "points": 50, "date": "2024-06-01", "status": "active", "created_at": "2024-08-05T11:34:00Z"}'
    '{"id": "achievement-006", "student_id": "student-002", "title": "Debate Champion", "description": "Won the inter-school debate competition", "category": "extracurricular", "points": 90, "date": "2024-04-05", "status": "active", "created_at": "2024-08-05T11:35:00Z"}'
    '{"id": "achievement-007", "student_id": "student-005", "title": "Community Service Award", "description": "100+ hours of community service", "category": "service", "points": 85, "date": "2024-05-20", "status": "active", "created_at": "2024-08-05T11:36:00Z"}'
)

counter=1
for achievement in "${achievements[@]}"; do
    id="achievement-$(printf "%03d" $counter)"
    curl -s -X POST "$QUERY_URL/query/service" \
        -u "$USERNAME:$PASSWORD" \
        -H "Content-Type: application/json" \
        -d "{\"statement\": \"INSERT INTO schoolmgmt.school.achievements (KEY, VALUE) VALUES (\\\"$id\\\", $achievement)\"}" > /dev/null
    
    if [ $? -eq 0 ]; then
        title=$(echo "$achievement" | grep -o '"title": "[^"]*"' | cut -d '"' -f 4)
        student_id=$(echo "$achievement" | grep -o '"student_id": "[^"]*"' | cut -d '"' -f 4)
        points=$(echo "$achievement" | grep -o '"points": [0-9]*' | cut -d ':' -f 2 | tr -d ' ')
        log_success "Created achievement: $title - $student_id ($points pts)"
    else
        log_warning "Failed to create achievement $id"
    fi
    ((counter++))
done

# Show summary statistics
log_info "📊 Loading completed! Database statistics:"

collections=("students" "teachers" "academics" "achievements")

for collection in "${collections[@]}"; do
    response=$(curl -s -X POST "$QUERY_URL/query/service" \
        -u "$USERNAME:$PASSWORD" \
        -H "Content-Type: application/json" \
        -d "{\"statement\": \"SELECT COUNT(*) as count FROM schoolmgmt.school.$collection\"}")
    
    count=$(echo "$response" | jq '.results[0].count' 2>/dev/null || echo "0")
    log_success "📈 $collection: $count records"
done

echo ""
log_success "🎉 Sample data loading completed successfully!"
echo ""
log_info "You can now test the API with the loaded data:"
echo "  curl http://localhost:8080/api/v1/students"
echo "  curl http://localhost:8080/api/v1/teachers"
echo "  curl http://localhost:8080/api/v1/academics"
echo "  curl http://localhost:8080/api/v1/achievements"
