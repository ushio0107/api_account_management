#!/bin/bash
MONGO_INITDB_ROOT_USERNAME="root"
MONGO_INITDB_ROOT_PASSWORD="123456"
DB_HOST="mongo"
DB_PORT="27017"

# Function to pull image from Docker Hub and deploy
pull_docker_image_and_deploy() {
    docker network create api
    docker pull ushio0107/account_management_api
    docker pull mongo:4.4
    docker run --name ${DB_HOST} \
    --network api \
	-e MONGO_INITDB_ROOT_USERNAME=${MONGO_INITDB_ROOT_USERNAME} \
	-e MONGO_INITDB_ROOT_PASSWORD=${MONGO_INITDB_ROOT_PASSWORD} \
	-v ./data:/data/db \
	-p 27017:27017 \
	-d mongo:4.4 
    docker run --name api --network api -e DB_HOST=${DB_HOST} -e DB_PORT=${DB_PORT} -e DB_USER=${MONGO_INITDB_ROOT_USERNAME} -e DB_PASSWORD=${MONGO_INITDB_ROOT_PASSWORD} ushio0107/account_management_api
}

# Function to clone repo and start containers
clone_repo_and_start_containers() {
    git clone git@github.com:ushio0107/api_account_management.git
    cd api_account_management
    vi .env # Set your environment variable if need.
    docker-compose up
}

# Main script
echo "Choose deployment method:"
echo "1. Pull Docker image from Docker Hub and run"
echo "2. Clone repo and run using Docker Compose"
read -p "Enter your choice (1/2): " choice

case $choice in
    1)
        pull_docker_image_and_deploy

        ;;
    2)
        clone_repo_and_start_containers
        ;;
    *)
        echo "Invalid choice. Exiting."
        ;;
esac
