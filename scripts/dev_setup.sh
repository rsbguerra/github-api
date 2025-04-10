#!/bin/bash
# This script sets up the development environment for the project.

# Install necessary packages and tools.
yay -Syu docker minikube kubectl docker-buildx

# Enable Docker service and add user to Docker group.
sudo systemctl enable docker
sudo usermod -aG docker "$USER" && newgrp docker

go get github.com/gin-gonic/gin
go get github.com/go-git/go-git/v5
go get github.com/google/go-github/v50/github
go get golang.org/x/oauth2