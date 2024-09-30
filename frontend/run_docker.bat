@echo off
docker build -t olympicfrontend .
docker run --rm -p 80:80 olympicfrontend
pause