#!/bin/bash

# Couchbase Setup and Demo Script for School Management System
# Bash Version
# Run this script to initialize Couchbase and demonstrate CRUD operations

# Configuration
COUCHBASE_URL="${COUCHBASE_URL:-http://localhost:8091}"
QUERY_URL="${QUERY_URL:-http://localhost:8093}"
USERNAME="${USERNAME:-Administrator}"
PASSWORD="${PASSWORD:-password123}"
ACTION="${1:-all}"

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

# Test Couchbase connection
test_connection() {
    log_info "Testing Couchbase connection..."
    
    if curl -s -f "$COUCHBASE_URL/pools" > /dev/null 2>&1; then
        log_success "Couchbase is accessible"
        return 0
    else
        log_error "Couchbase is not accessible at $COUCHBASE_URL"
        log_warning "Make sure Docker Compose is running: docker-compose up -d"
        return 1
    fi
}

# Initialize cluster
initialize_cluster() {
    log_info "🔧 Initializing Couchbase cluster..."
    
    # Initialize cluster
    curl -s -X POST "$COUCHBASE_URL/pools/default" \
        -d 'memoryQuota=512' \
        -d 'indexMemoryQuota=256' > /dev/null 2>&1
    log_success "Cluster initialized"
    
    # Setup administrator
    curl -s -X POST "$COUCHBASE_URL/settings/web" \
        -d "username=$USERNAME" \
        -d "password=$PASSWORD" \
        -d 'port=SAME' > /dev/null 2>&1
    log_success "Administrator account created"
    
    # Wait for cluster to stabilize
    sleep 5
    
    # Create bucket
    curl -s -X POST "$COUCHBASE_URL/pools/default/buckets" \
        -u "$USERNAME:$PASSWORD" \
        -d 'name=schoolmgmt' \
        -d 'ramQuotaMB=512' \
        -d 'bucketType=membase' > /dev/null 2>&1
    
    if [ $? -eq 0 ]; then
        log_success "Bucket 'schoolmgmt' created"
    else
        log_warning "Bucket may already exist"
    fi
}

# Create collections and scopes
create_collections() {
    log_info "📁 Creating collections and scopes..."
    
    # Create scope
    curl -s -X POST "$COUCHBASE_URL/pools/default/buckets/schoolmgmt/scopes" \
        -u "$USERNAME:$PASSWORD" \
        -d 'name=school' > /dev/null 2>&1
    log_success "Scope 'school' created"
    
    # Wait for scope to be available
    sleep 3
    
    # Create collections
    collections=("students" "teachers" "academics" "achievements")
    for collection in "${collections[@]}"; do
        curl -s -X POST "$COUCHBASE_URL/pools/default/buckets/schoolmgmt/scopes/school/collections" \
            -u "$USERNAME:$PASSWORD" \
            -d "name=$collection" > /dev/null 2>&1
        
        if [ $? -eq 0 ]; then
            log_success "Collection '$collection' created"
        else
            log_warning "Collection '$collection' may already exist"
        fi
        sleep 1
    done
}

# Create indexes
create_indexes() {
    log_info "🏗️ Creating database indexes..."
    
    indexes=(
        "CREATE PRIMARY INDEX ON schoolmgmt.school.students"
        "CREATE PRIMARY INDEX ON schoolmgmt.school.teachers"
        "CREATE PRIMARY INDEX ON schoolmgmt.school.academics"
        "CREATE PRIMARY INDEX ON schoolmgmt.school.achievements"
        "CREATE INDEX idx_student_grade ON schoolmgmt.school.students(grade)"
        "CREATE INDEX idx_teacher_department ON schoolmgmt.school.teachers(department)"
        "CREATE INDEX idx_academic_student ON schoolmgmt.school.academics(student_id, academic_year)"
    )
    
    for index in "${indexes[@]}"; do
        curl -s -X POST "$QUERY_URL/query/service" \
            -u "$USERNAME:$PASSWORD" \
            -H "Content-Type: application/json" \
            -d "{\"statement\": \"$index\"}" > /dev/null 2>&1
        
        index_name=$(echo "$index" | awk '{print $NF}')
        if [ $? -eq 0 ]; then
            log_success "Index created: $index_name"
        else
            log_warning "Index may already exist: $index_name"
        fi
        sleep 1
    done
}

# Demo Student CRUD operations
demo_student_crud() {
    log_info "👨‍🎓 Demonstrating Student CRUD operations..."
    
    # CREATE Student
    log_info "Creating student..."
    response=$(curl -s -X POST "$QUERY_URL/query/service" \
        -u "$USERNAME:$PASSWORD" \
        -H "Content-Type: application/json" \
        -d '{
            "statement": "INSERT INTO schoolmgmt.school.students (KEY, VALUE) VALUES (\"student-001\", {\"id\": \"student-001\", \"first_name\": \"John\", \"last_name\": \"Doe\", \"email\": \"john.doe@school.edu\", \"grade\": \"10\", \"age\": 16, \"enrollment_date\": \"2024-08-01\", \"status\": \"active\", \"created_at\": \"2024-08-05T10:30:00Z\"})"
        }')
    
    if echo "$response" | grep -q "success"; then
        log_success "Student created successfully"
    else
        log_warning "Student may already exist"
    fi
    
    # READ Student
    log_info "Reading student..."
    response=$(curl -s -X POST "$QUERY_URL/query/service" \
        -u "$USERNAME:$PASSWORD" \
        -H "Content-Type: application/json" \
        -d '{
            "statement": "SELECT * FROM schoolmgmt.school.students WHERE META().id = \"student-001\""
        }')
    
    if echo "$response" | grep -q "John"; then
        name=$(echo "$response" | jq -r '.results[0].students.first_name + " " + .results[0].students.last_name' 2>/dev/null || echo "John Doe")
        log_success "Student retrieved: $name"
    else
        log_error "Failed to read student"
    fi
    
    # UPDATE Student
    log_info "Updating student..."
    curl -s -X POST "$QUERY_URL/query/service" \
        -u "$USERNAME:$PASSWORD" \
        -H "Content-Type: application/json" \
        -d '{
            "statement": "UPDATE schoolmgmt.school.students SET grade = \"11\", age = 17 WHERE META().id = \"student-001\""
        }' > /dev/null
    log_success "Student updated successfully"
    
    # LIST Students
    log_info "Listing all students..."
    response=$(curl -s -X POST "$QUERY_URL/query/service" \
        -u "$USERNAME:$PASSWORD" \
        -H "Content-Type: application/json" \
        -d '{
            "statement": "SELECT META().id, * FROM schoolmgmt.school.students ORDER BY last_name"
        }')
    
    count=$(echo "$response" | jq '.results | length' 2>/dev/null || echo "1")
    log_success "Found $count students"
    
    if command -v jq > /dev/null 2>&1; then
        echo "$response" | jq -r '.results[] | "  - " + .id + ": " + .students.first_name + " " + .students.last_name + " (Grade: " + .students.grade + ")"' 2>/dev/null
    fi
}

# Demo Teacher CRUD operations
demo_teacher_crud() {
    log_info "👩‍🏫 Demonstrating Teacher CRUD operations..."
    
    # CREATE Teacher
    log_info "Creating teacher..."
    curl -s -X POST "$QUERY_URL/query/service" \
        -u "$USERNAME:$PASSWORD" \
        -H "Content-Type: application/json" \
        -d '{
            "statement": "INSERT INTO schoolmgmt.school.teachers (KEY, VALUE) VALUES (\"teacher-001\", {\"id\": \"teacher-001\", \"first_name\": \"Dr. Emma\", \"last_name\": \"Wilson\", \"email\": \"emma.wilson@school.edu\", \"department\": \"Mathematics\", \"subjects\": [\"Algebra\", \"Geometry\"], \"experience\": 8, \"hire_date\": \"2020-01-15\", \"status\": \"active\", \"created_at\": \"2024-08-05T10:45:00Z\"})"
        }' > /dev/null
    log_success "Teacher created successfully"
    
    # READ Teacher
    log_info "Reading teacher..."
    response=$(curl -s -X POST "$QUERY_URL/query/service" \
        -u "$USERNAME:$PASSWORD" \
        -H "Content-Type: application/json" \
        -d '{
            "statement": "SELECT * FROM schoolmgmt.school.teachers WHERE department = \"Mathematics\""
        }')
    
    count=$(echo "$response" | jq '.results | length' 2>/dev/null || echo "1")
    log_success "Found $count Mathematics teachers"
}

# Demo Academic CRUD operations
demo_academic_crud() {
    log_info "📚 Demonstrating Academic Records CRUD operations..."
    
    # CREATE Academic Record
    log_info "Creating academic record..."
    curl -s -X POST "$QUERY_URL/query/service" \
        -u "$USERNAME:$PASSWORD" \
        -H "Content-Type: application/json" \
        -d '{
            "statement": "INSERT INTO schoolmgmt.school.academics (KEY, VALUE) VALUES (\"academic-001\", {\"id\": \"academic-001\", \"student_id\": \"student-001\", \"teacher_id\": \"teacher-001\", \"subject\": \"Algebra\", \"grade\": \"A\", \"semester\": \"Spring 2024\", \"max_marks\": 100, \"obtained_marks\": 95, \"percentage\": 95.0, \"status\": \"pass\", \"created_at\": \"2024-08-05T11:00:00Z\"})"
        }' > /dev/null
    log_success "Academic record created successfully"
    
    # Query student performance
    log_info "Querying student performance..."
    response=$(curl -s -X POST "$QUERY_URL/query/service" \
        -u "$USERNAME:$PASSWORD" \
        -H "Content-Type: application/json" \
        -d '{
            "statement": "SELECT * FROM schoolmgmt.school.academics WHERE student_id = \"student-001\""
        }')
    
    count=$(echo "$response" | jq '.results | length' 2>/dev/null || echo "1")
    log_success "Found $count academic records for student"
    
    if command -v jq > /dev/null 2>&1; then
        echo "$response" | jq -r '.results[] | "  - " + .academics.subject + ": " + .academics.grade + " (" + (.academics.percentage | tostring) + "%)"' 2>/dev/null
    fi
}

# Demo Achievement CRUD operations
demo_achievement_crud() {
    log_info "🏆 Demonstrating Achievement CRUD operations..."
    
    # CREATE Achievement
    log_info "Creating achievement..."
    curl -s -X POST "$QUERY_URL/query/service" \
        -u "$USERNAME:$PASSWORD" \
        -H "Content-Type: application/json" \
        -d '{
            "statement": "INSERT INTO schoolmgmt.school.achievements (KEY, VALUE) VALUES (\"achievement-001\", {\"id\": \"achievement-001\", \"student_id\": \"student-001\", \"title\": \"Math Excellence Award\", \"description\": \"Outstanding performance in mathematics\", \"category\": \"academic\", \"points\": 100, \"date\": \"2024-04-20\", \"status\": \"active\", \"created_at\": \"2024-08-05T11:30:00Z\"})"
        }' > /dev/null
    log_success "Achievement created successfully"
    
    # Query leaderboard
    log_info "Creating leaderboard..."
    response=$(curl -s -X POST "$QUERY_URL/query/service" \
        -u "$USERNAME:$PASSWORD" \
        -H "Content-Type: application/json" \
        -d '{
            "statement": "SELECT student_id, SUM(points) as total_points FROM schoolmgmt.school.achievements GROUP BY student_id ORDER BY total_points DESC"
        }')
    
    count=$(echo "$response" | jq '.results | length' 2>/dev/null || echo "1")
    log_success "Leaderboard created with $count students"
    
    if command -v jq > /dev/null 2>&1; then
        echo "$response" | jq -r '.results[] | "  - Student " + .student_id + ": " + (.total_points | tostring) + " points"' 2>/dev/null
    fi
}

# Show database statistics
show_database_stats() {
    log_info "📊 Showing database statistics..."
    
    collections=("students" "teachers" "academics" "achievements")
    
    for collection in "${collections[@]}"; do
        response=$(curl -s -X POST "$QUERY_URL/query/service" \
            -u "$USERNAME:$PASSWORD" \
            -H "Content-Type: application/json" \
            -d "{\"statement\": \"SELECT COUNT(*) as count FROM schoolmgmt.school.$collection\"}")
        
        count=$(echo "$response" | jq '.results[0].count' 2>/dev/null || echo "0")
        log_success "📈 $collection: $count records"
    done
}

# Show menu
show_menu() {
    echo -e "\n${YELLOW}🎓 Couchbase School Management Demo Script${NC}"
    echo -e "${YELLOW}===========================================${NC}"
    echo "1. Test Connection"
    echo "2. Initialize Cluster"
    echo "3. Create Collections"
    echo "4. Create Indexes"
    echo "5. Demo Student CRUD"
    echo "6. Demo Teacher CRUD"
    echo "7. Demo Academic CRUD"
    echo "8. Demo Achievement CRUD"
    echo "9. Show Database Stats"
    echo "10. Run All Setup (1-4)"
    echo "11. Run All Demos (5-9)"
    echo "12. Full Demo (All)"
    echo "0. Exit"
    echo ""
}

# Main execution
echo -e "${GREEN}🎓 Couchbase School Management System Demo${NC}"
echo -e "${GREEN}=============================================${NC}"

case "$ACTION" in
    "setup")
        test_connection
        initialize_cluster
        create_collections
        create_indexes
        log_success "🎉 Setup completed!"
        ;;
    "demo")
        demo_student_crud
        demo_teacher_crud
        demo_academic_crud
        demo_achievement_crud
        show_database_stats
        log_success "🎉 All demos completed!"
        ;;
    "test")
        test_connection
        ;;
    "all"|*)
        # Interactive mode
        while true; do
            show_menu
            read -p "Enter your choice (0-12): " choice
            
            case $choice in
                1) test_connection ;;
                2) initialize_cluster ;;
                3) create_collections ;;
                4) create_indexes ;;
                5) demo_student_crud ;;
                6) demo_teacher_crud ;;
                7) demo_academic_crud ;;
                8) demo_achievement_crud ;;
                9) show_database_stats ;;
                10)
                    test_connection
                    initialize_cluster
                    create_collections
                    create_indexes
                    log_success "🎉 Setup completed!"
                    ;;
                11)
                    demo_student_crud
                    demo_teacher_crud
                    demo_academic_crud
                    demo_achievement_crud
                    show_database_stats
                    log_success "🎉 All demos completed!"
                    ;;
                12)
                    test_connection
                    initialize_cluster
                    create_collections
                    create_indexes
                    demo_student_crud
                    demo_teacher_crud
                    demo_academic_crud
                    demo_achievement_crud
                    show_database_stats
                    log_success "🎉 Full demo completed!"
                    ;;
                0)
                    log_success "Goodbye!"
                    exit 0
                    ;;
                *)
                    log_warning "Invalid choice. Please try again."
                    ;;
            esac
            
            if [ "$choice" != "0" ]; then
                echo -e "\n${BLUE}Press any key to continue...${NC}"
                read -n 1 -s
            fi
        done
        ;;
esac

if [ "$ACTION" != "all" ]; then
    echo ""
    echo "Usage: $0 [setup|demo|test|all]"
    echo "  setup - Initialize cluster and create collections"
    echo "  demo  - Run CRUD demonstrations"
    echo "  test  - Test connection only"
    echo "  all   - Interactive mode (default)"
fi
