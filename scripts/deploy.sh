#!/bin/bash
set -e


# Step 1: Start Minikube
echo "Starting Minikube..."
minikube start

# Step 2: Use Minikube's Docker daemon
echo "Configuring Docker to use Minikube's environment..."
eval "$(minikube docker-env)"

# Step 3: Build the Docker image
echo "Building Docker image..."
docker build -t ghcr.io/rsbguerra/github-api:latest .

# Step 4: Apply Kubernetes manifests
echo "Applying Kubernetes manifests..."
kubectl apply -f manifests/deployment.yaml
kubectl apply -f manifests/service.yaml

# Step 5: Verify pod status
echo "Checking pod status..."
kubectl get pods

# Step 7: Expose the service
echo "Accessing the service..."
minikube service github-api-service