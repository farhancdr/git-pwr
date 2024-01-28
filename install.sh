#!/bin/bash

# Set the root directory
ROOT_DIR=~

# Set the project name
PROJECT_NAME="git-pwr"

cd ~

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Go is not installed. Installing Go..."
    
    # Install Go
    # This assumes you are on a macOS or Linux system
    # You may need to modify this part if your users are on a different OS
    if [[ $(uname) == "Darwin" ]]; then
        # Check if Homebrew is installed
        if ! command -v brew &> /dev/null; then
            echo "Homebrew is not installed. Installing Homebrew..."
            /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
        fi
        # Install Go using Homebrew
        brew install go
    elif [[ $(uname) == "Linux" ]]; then
        # Install Go on Linux using apt
        sudo apt update
        sudo apt install -y golang
    else
        echo "Unsupported operating system. Please install Go manually."
        exit 1
    fi
fi

# Clone the project
echo "Cloning the project..."
git clone https://github.com/farhancdr/$PROJECT_NAME.git

# Go to the project directory
cd $PROJECT_NAME

# Build the project
echo "Building the project..."
go build

# Move the executable to the bin directory
echo "Moving the executable to bin directory..."
sudo mv $PROJECT_NAME /usr/local/bin

cd ~
#removing the source project
rm -rf $PROJECT_NAME

echo "Installation complete. You can now use '$PROJECT_NAME' from the command line."
