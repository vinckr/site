#!/bin/bash

# List of excluded files
excluded_files=("quotes.md" "links.md")

# Build the find command with excluded files
find_cmd="find content -iname \"*.md\""
for file in "${excluded_files[@]}"; do
	find_cmd+=" -not -iname \"$file\""
done

# Run the find command and execute vale for each file
eval "$find_cmd -exec vale {} \;"
