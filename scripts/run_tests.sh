#!/bin/bash

# This script runs the tests for the project.
go mod tidy

source ../.env

# Load environment variables from the .env file if it exists.
if [ -f .env ]; then
  export "$(grep -v '^#' .env | xargs)"
fi

echo $TEST_AUTH_TOKEN

# Run all tests in the project.
go test ../... -v