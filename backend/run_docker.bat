@echo off
docker build -t olympicsearch .
docker run -d -p 6563:8888 olympicsearch
pause