#!/bin/bash
cd /app/cmd
go get .
cd migrate && go run migrate.go db init && go run . db migrate
cd ..
go build -o /app/app_exe .
/app/app_exe