#!/bin/bash

# Define variables
IMAGE_NAME="taskly-js-server"
CONTAINER_NAME="taskly-api"
HOST_PORT=8000
CONTAINER_PORT=8000
ENV_FILE=".env"

# Build the Docker image
echo "Building the Docker image..."
docker build -t $IMAGE_NAME .

# Check if the container is already running
if [ $(docker ps -q -f name=$CONTAINER_NAME) ]; then
  echo "Stopping the existing container..."
  docker stop $CONTAINER_NAME
fi

# Remove the container if it exists
if [ $(docker ps -aq -f name=$CONTAINER_NAME) ]; then
  echo "Removing the existing container..."
  docker rm $CONTAINER_NAME
fi

# Run the Docker container
echo "Running the Docker container..."
docker run -d \
  --name $CONTAINER_NAME \
  -p $HOST_PORT:$CONTAINER_PORT \
  --env-file $ENV_FILE \
  $IMAGE_NAME

echo "Container is running at http://localhost:$HOST_PORT"
