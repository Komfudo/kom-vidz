#!/bin/bash

# MySQL credentials
MYSQL_USER="vidzdb"
MYSQL_PASSWORD="cockballs"

# Database and table details
DATABASE_NAME="komvidz"
TABLE_NAME="videos"

# Prompt for user input
read -p "Enter Title: " title
read -p "Enter FilePath: " filePath

# Generate the current date and time
dateTime=$(date +"%Y-%m-%d %H:%M:%S")

# Construct the SQL query
query="INSERT INTO $DATABASE_NAME.$TABLE_NAME (Title, DateTime, FilePath) VALUES ('$title', '$dateTime', '$filePath');"

# Execute the SQL query
mysql -u $MYSQL_USER -p$MYSQL_PASSWORD -e "$query"
