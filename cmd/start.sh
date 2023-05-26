#!/bin/bash
cd /app/cmd
go get .
go build -o /app/app_exe .
/app/app_exe