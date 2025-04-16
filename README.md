# GitHub API

## Description

This project is a REST API written in Go that interacts with the GitHub API to manage repositories and pull requests. It
provides endpoints to create, delete, and list repositories, as well as retrieve open pull requests and their
contributors for a specific repository.

## Features

- **Repository Management**:
    - Create a new repository.
    - Delete an existing repository.
    - List repositories for a user.

- **Pull Request Management**:
    - List open pull requests for a repository.
    - Retrieve the names of contributors on pull requests.

- **Root Endpoint**:
    - A simple health check endpoint.

## Requirements

- REST API to:
    - Create, destroy, and list repositories on GitHub.
    - List open pull requests and contributors for a specific repository.
- Deployment on Minikube.
- CI/CD pipeline to:
    - Run tests.
    - Perform linting and security checks.
    - Deploy the application.

## Technologies Used

- **Programming Language**: Go
- **Frameworks**:
    - [Gin](https://github.com/gin-gonic/gin) for building the REST API.
    - [Testify](https://github.com/stretchr/testify) for testing.
- **GitHub API**: Integration with the GitHub API using the `go-github` library.
- **Deployment**: Minikube for local Kubernetes deployment.

## Usage

1. Run the application:

```bash
./deploy.sh
```

2. Access the API at http://localhost:{PORT}.

## Endpoints

### Repository Management

- **Create Repository**: `POST /repositories/{auth-token}`
    - Request Body:
        ```json
        {
            "data": {
                "name": "string",
                "private": true,
                "url": "string",
                "owner": {
                    "name": "string"
                    ...
                }
                ...
            }
        }
        ```
    - Response:
        ```json
        {
            "data": {
                "name": "string",
                "private": true,
                "url": "string",
                "owner": {
                    "name": "string"
                    ...
                }
                ...
            }
        }
        ```
- **Delete Repository**: `DELETE /repositories/{auth-token}`
    - Request Body:
      ```json
        {
            "data": {
                "name": "string",
                "private": true,
                "url": "string",
                "owner": {
                    "name": "string"
                    ...
                }
                ...
            }
        }
      ```
- **List Repositories**: `GET /repositories/{auth-token}`
    - Response:
        ```json
        [{
            "data": {
                "name": "string",
                "private": true,
                "url": "string",
                "owner": {
                    "name": "string"
                    ...
                }
                ...
            }
        },
      ...]
        ```

### Pull Request Management

- **List Open Pull Requests**: `GET /repositories/:owner/:repo/pulls`
    - Response:
        ```json
        [
            {
                "title": "pull-request-title",
                "number": 1,
                "user": {
                    "login": "contributor-username"
                }
            }
        ]
        ```

## Testing
The `run_tests.sh` script is designed to automate the process of running tests for the GitHub API project. It performs the following tasks:

1. **Environment Setup**: Ensures that the necessary environment variables and dependencies are properly configured before running the tests.
2. **Run Unit Tests**: Executes all unit tests in the project using the `go test` command, which scans the project directories for test files and runs them.
3. **Generate Test Reports**: Optionally generates a test report or outputs the results in a structured format for further analysis.
4. **Error Handling**: Checks for any test failures and exits with a non-zero status code if errors are encountered, making it suitable for integration into CI/CD pipelines.

This script simplifies the testing process by consolidating all test-related tasks into a single command, ensuring consistency and reliability during development and deployment.

## Deployment
The deployment script (`deploy.sh`) is used to automate the process of deploying the GitHub API application. It performs the following tasks:

1. **Build the Application**: Compiles the Go application into an executable binary.
2. **Create Docker Image**: Builds a Docker image for the application, ensuring it can run in a containerized environment.
3. **Push Docker Image**: Pushes the built Docker image to a container registry (if configured).
4. **Deploy to Kubernetes**: Applies the Kubernetes deployment and service configurations using `kubectl`, deploying the application to a Minikube cluster.
5. **Expose the Service**: Ensures the application is accessible via Minikube's service.

This script simplifies the deployment process by combining multiple steps into a single command, making it easier to deploy and test the application locally or in a CI/CD pipeline.

## CI/CD
The `.github/workflows/ci-cd.yml` file defines a GitHub Actions workflow for the CI/CD pipeline of the project. It automates various tasks to ensure code quality, security, and deployment. Below is a description of its key components:

1. **Trigger Events**:
    - The workflow is triggered on `push` and `pull_request` events targeting the `dev-cdci` branch.

2. **Jobs**:
    - **Build**:
        - Runs on the `ubuntu-latest` environment.
        - Performs the following steps:
            1. **Checkout Code**: Clones the repository.
            2. **Set Up Go**: Configures the Go environment with version 1.24.
            3. **Install Dependencies**: Runs `go mod tidy` to install and clean up dependencies.
            4. **Run Tests**: Executes all tests in the project using `go test`.
            5. **Linting**: Uses `golangci-lint` to check for code quality issues.
            6. **Security Checks**: Runs `gosec` to identify potential security vulnerabilities.
            7. **Build Docker Image**: Builds a Docker image for the application and tags it as `ghcr.io/rsbguerra/github-api:latest`.
            8. **Set Up Minikube**: Installs and starts Minikube for local Kubernetes deployment.
            9. **Deploy to Minikube**: Applies Kubernetes deployment and service configurations using `kubectl`.

This workflow ensures that the application is tested, linted, secured, and deployed to a local Kubernetes cluster as part of the CI/CD process.