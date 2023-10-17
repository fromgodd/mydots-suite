#!/bin/bash

# Destination directory for source code and build
BUILD_DIR="$HOME/gnome-build"
mkdir -p "$BUILD_DIR"

# List of GNOME 3.38 source tarballs from the GNOME website
SOURCE_URL="https://download.gnome.org/core/3.38/3.38.4/sources/"
SOURCE_TARBALLS=(
    "gnome-shell-3.38.4.tar.xz"
    "gnome-session-3.38.4.tar.xz"
    # Add more tarballs as needed
)

# Download and extract source tarballs
for tarball in "${SOURCE_TARBALLS[@]}"; do
    # Download the tarball
    wget "${SOURCE_URL}${tarball}" -P "$BUILD_DIR"
    
    # Extract the tarball
    tar -xf "${BUILD_DIR}/${tarball}" -C "$BUILD_DIR"
done

# Build and install GNOME components
for tarball in "${SOURCE_TARBALLS[@]}"; do
    # Extract the directory name from the tarball filename
    dir_name="${tarball%.tar.xz}"
    
    # Navigate to the source directory
    cd "${BUILD_DIR}/${dir_name}"
    
    # Configure, build, and install
    ./configure
    make
    sudo make install
    
    # Optional: Clean up build artifacts
    make clean
done
