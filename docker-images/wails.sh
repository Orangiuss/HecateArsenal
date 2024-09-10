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
    echo "  -r, --rebuild, rebuild     Rebuild the Docker image"
    echo "  -h, --help, help           Display this help message"
    echo "  -d, --down, down           Tear down the deployment"
    echo "  -s, --shell, shell         Start a shell in the Docker container"
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
        -r|--rebuild|rebuild)
            rebuild_image
            shift
            ;;
        -h|--help|help)
            display_help
            exit 0
            ;;
        -d|--down|down)
            tear_down
            exit 0
            ;;
        -s|--shell|shell)
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

# Allow X11 forwarding
xhost +local:docker

# Default behavior: build and run the Docker container with DISPLAY and X11 forwarding
rebuild_image
docker run -d -p 5555:5555 -p 5173:5173 -v tmpdocker:/app -e DISPLAY=$DISPLAY -v /tmp/.X11-unix:/tmp/.X11-unix --name $IMAGE_NAME $IMAGE_NAME

docker exec -it $IMAGE_NAME /bin/bash -c "bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)"
docker exec -it $IMAGE_NAME /bin/bash -c "source /root/.gvm/scripts/gvm; gvm install go1.20"
docker exec -it $IMAGE_NAME /bin/bash -c "source /root/.gvm/scripts/gvm; gvm use go1.20 --default"


# Test if dev environment is working
# docker exec -it $IMAGE_NAME /bin/bash -c "source /root/.gvm/scripts/gvm; gvm use go1.20; wails init -n myproject -t vue; cd myproject; wails dev"

# chown for the script
# chmod +x setup.sh