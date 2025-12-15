# Copilot Prompts – Demo 9: School Microservice

## Primary Prompt (Ask Mode)
```
## SYSTEM ROLE
You are a senior backend architect and system designer.
You are NOT allowed to write code in this response.

## TASK
Analyze the problem and propose MULTIPLE ARCHITECTURAL APPROACHES
for building a Go-based school management microservice.

## IMPORTANT CONSTRAINTS (MUST FOLLOW)
- ❌ DO NOT generate any code
- ❌ DO NOT provide APIs, structs, or Go examples
- ❌ DO NOT assume a final design
- ✅ ONLY discuss design approaches and trade-offs
- ✅ Provide at least 3 clearly distinct approaches
- ✅ Each approach must have pros, cons, and when to use it

## PROBLEM STATEMENT
We need a Go microservice for school operations:
- Manage students, teachers, classes, academics, exams, achievements
- Exposed via REST APIs for CRUD
- Uses Couchbase as the data store
- Accessed by external services
- Expected traffic: ~200 TPS during peak hours (10 AM – 5 PM)

## REQUIRED OUTPUT FORMAT
For EACH approach, strictly follow this format:

### Approach <N>: <Approach Name>
**Description**
- High-level architecture
- Data modeling philosophy
- How academics, exams, achievements are handled

**Pros**
- Bullet points

**Cons**
- Bullet points

**Best Fit When**
- Bullet points

## FINAL SECTION
Add a short comparison table summarizing all approaches.

## REMINDER
If you write any code or API definitions, the response is invalid.
```

## Copilot Agent Mode
```
Build the microservice using approach #
Create a changelog.md file to keep track of the changes and always keep it updated.
Make Couchbase credentials configurable in the application.
Implement the necessary code changes for all microservices to support this configuration.
Create a Docker Compose file to simplify starting the service.
Provide step-by-step instructions for launching the service using Docker Compose.
Demonstrate that the service is running successfully.
```

## Follow-up Prompts
```
Explain the function <selected>
Optimize <func>
Analyze <func> for concurrency issues or race conditions
```

## Kubernetes Deployment Prompts
```
Generate Dockerfile and helm charts for this service for deployment in kubernetes cluster in the namespace school-demo.
Update the changelog.md file accordingly.
```

## Image Justification
```
Explain Dockerfile. Include details of the base os image version and justify your choice. Suggest alternatives.
```

## Documentation
```
[Edit mode] [select the api] Generate API documentation in go doc style.
Update the changelog.md file accordingly.
Generate a high level design document as a markdown file.
```

## Unit Test Case Generation
```
Generate unit tests for the StudentService using the Go testing framework. Additionally, configure the project to generate test coverage reports. Finally, update the changelog.md file to document the addition of unit tests and test coverage configuration.
```
