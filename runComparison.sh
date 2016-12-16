#!/bin/bash
# Executes all three analysis
echo "Analyse with safesql"
safesql github.com/akwick/sqinco 
echo "Analyse with gas"
gas -include=G201,G202 ./...
echo "Analyse with our tool"
gotcha -src="main.go" -ssf="sqlQuery.txt" -pkgs="github.com/akwick/sqinco/sqlInjections"

