#!/bin/bash

# Name of the Docker image
IMAGE_NAME=ha-wails-image
NAME_SCRIPT=wails.sh
# Empty if Dockerfile file name is Dockerfile
DOCKERFILE_NAME=Dockerfile_wails

# Function to display help message
display_help() {
    echo "Usage: $NAME_SCRIPT [OPTIONS]"
    echo "Options:"
    echo "  -r, --rebuild     Rebuild the Docker image"
    echo "  -h, --help        Display this help message"
    echo "  -d, --down        Tear down the deployment"
    echo "  -s, --shell       Start a shell in the Docker container"
}

# Function to rebuild the Docker image
rebuild_image() {
    if [ ! -s $DOCKERFILE_NAME ]; then
        docker build -t $IMAGE_NAME .
    else
        docker build -t $IMAGE_NAME -f $DOCKERFILE_NAME .
    fi
}

# Function to tear down the deployment
tear_down() {
    docker stop $IMAGE_NAME
    docker rm $IMAGE_NAME
}

# Parse command line options
while [[ $# -gt 0 ]]; do
    case $1 in
        -r|--rebuild)
            rebuild_image
            shift
            ;;
        -h|--help)
            display_help
            exit 0
            ;;
        -d|--down)
            tear_down
            exit 0
            ;;
        -s|--shell)
            docker exec -it $IMAGE_NAME /bin/bash
            exit 0
            ;;
        *)
            echo "Invalid option: $1"
            display_help
            exit 1
            ;;
    esac
done

# Default behavior: build and run the Docker container
rebuild_image
docker run -d --name $IMAGE_NAME $IMAGE_NAME

# chown for the script
# chmod +x setup.sh