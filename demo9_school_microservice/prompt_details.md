# Copilot Prompts – Demo 9: School Microservice

## Primary Prompt (Ask Mode)
```
Need to generate a Go microservice that can handle operations of a school.
Service should maintain list of students, teachers, classes, academics and achievements.
This service will be accessed by the external services.
Need to provide REST interfaces for CRUD operations.
Need to build suitable logic for academics, exams, and achievements.
Utilize Couchbase as the storage database.
Expected traffic is about 200 tps in the peak hours 10 am to 5 pm.
Provide me the details of possible approaches detailing each one's pros & cons.
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
