#!/bin/bash
# Executes all three analysis
echo "Analyse with safesql"
safesql goretech/comparison
echo "Analyse with gas"
gas -include=G201,G202 ./...
echo "Analyse with our tool"
analysis -src="main.go" -ssf="sqlQuery.txt" -pkgs="goretech/comparison/sqlInjections"

