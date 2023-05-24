#!/bin/bash

# MySQL credentials
MYSQL_USER="your_username"
MYSQL_PASSWORD="your_password"

# Database and table details
DATABASE_NAME="komvidz"
TABLE_NAME="videos"

# Prompt for user input
read -p "Enter ID: " id
read -p "Enter Title: " title
read -p "Enter FilePath: " filePath

# Generate the current date and time
dateTime=$(date +"%Y-%m-%d %H:%M:%S")

# Construct the SQL query
query="INSERT INTO $DATABASE_NAME.$TABLE_NAME (ID, Title, DateTime, FilePath) VALUES ('$id', '$title', '$dateTime', '$filePath');"

# Execute the SQL query
mysql -u $MYSQL_USER -p$MYSQL_PASSWORD -e "$query"
