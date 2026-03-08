#!/bin/bash

# Function to print status messages
print_status() {
  echo "====================================="
  echo "$1"
  echo "====================================="
}

print_status "Updating package lists..."
sudo apt-get update

print_status "Installing dependencies..."
sudo apt-get install -y ca-certificates curl

print_status "Creating directory for Docker's GPG key..."
sudo install -m 0755 -d /etc/apt/keyrings

print_status "Downloading Docker's official GPG key..."
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc

print_status "Setting permissions for the GPG key..."
sudo chmod a+r /etc/apt/keyrings/docker.asc

print_status "Adding Docker repository to Apt sources..."
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

print_status "Updating package lists with Docker repository..."
sudo apt-get update

print_status "Installing Docker..."
sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

print_status "Starting Docker..."
sudo service docker start

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

