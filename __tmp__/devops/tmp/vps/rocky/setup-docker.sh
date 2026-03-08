#!/bin/bash

# Function to print status messages
print_status() {
  echo "====================================="
  echo "$1"
  echo "====================================="
}

print_status "Installing yum-utils..."
sudo yum install -y yum-utils

print_status "Adding Docker repository to Yum sources..."
sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo

print_status "Installing Docker..."
sudo yum install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

print_status "Starting Docker service..."
sudo systemctl start docker

print_status "Adding Docker group (if not already exists)..."
sudo groupadd docker

print_status "Adding user $USER to Docker group..."
sudo usermod -aG docker $USER

print_status "Verifying Docker installation..."

# Check Docker version
docker_version=$(docker --version)
if [[ $? -eq 0 ]]; then
  echo "Docker was successfully installed! Version: $docker_version"
else
  echo "Failed to verify Docker installation. Please ensure Docker is installed correctly."
fi

# Final message to the user
echo "Log out and log back in so that your group membership is re-evaluated, or run 'newgrp docker' to activate the changes immediately."

