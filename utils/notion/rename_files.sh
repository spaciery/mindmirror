#!/bin/bash

# Check if a directory path is provided as an argument
if [ -z "$1" ]; then
  echo "Usage: $0 <directory_path>"
  exit 1
fi

# Directory to process
DIRECTORY="$1"

# Check if the directory exists
if [ ! -d "$DIRECTORY" ]; then
  echo "Directory not found. Please enter a valid directory path."
  exit 1
fi

# Function to rename files and directories
rename_items() {
  for item in "$1"/*; do
    # Extract the basename and dirname
    base_name=$(basename "$item")
    dir_name=$(dirname "$item")

    # Remove the alphanumerical part from the filename
    new_name=$(echo "$base_name" | sed -E 's/ [a-f0-9]{32}//')

    # Rename the file/folder if the name has changed
    if [ "$base_name" != "$new_name" ]; then
      mv "$item" "$dir_name/$new_name"
      echo "Renamed: $item -> $dir_name/$new_name"
      item="$dir_name/$new_name"  # Update the item path after renaming
    fi

    # Check if item is a directory and rename recursively
    if [ -d "$item" ]; then
      rename_items "$item"
    fi
  done
}

# Run the renaming function on the specified directory
rename_items "$DIRECTORY"
