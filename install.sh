#!/bin/bash

# Replace these placeholders with your actual values
REPO_URL="https://github.com/farhancdr/git-pwr"
CLI_NAME="git-pwr"
BIN_DIRECTORY="/usr/local/bin"
GO_VERSION="1.21"  # Change to the required Go version

cd ~

# Function to display error and exit
error_exit() {
    echo "Error: $1" >&2
    exit 1
}

# Check if Go is installed
if ! command -v go &> /dev/null; then
    # Install Go if not present
    echo "Installing Go ${GO_VERSION}..."
    wget https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz || error_exit "Failed to download Go"
    sudo tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz || error_exit "Failed to extract Go"
    rm go${GO_VERSION}.linux-amd64.tar.gz
    export PATH=$PATH:/usr/local/go/bin
fi

# Clone the GitHub repository
echo "Cloning the repository..."
git clone ${REPO_URL} || error_exit "Failed to clone the repository"
cd ${CLI_NAME} || error_exit "Failed to enter the project directory"

# Build the project
echo "Building the project..."
go build || error_exit "Failed to build the project"

# Move the executable to the bin directory
echo "Moving the executable to ${BIN_DIRECTORY}..."
sudo mv ${CLI_NAME} ${BIN_DIRECTORY} || error_exit "Failed to move the executable"
cd ~
rm -rf git-pwr

echo "Installation completed successfully!"