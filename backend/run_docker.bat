@echo off
docker build -t olympicsearch .
docker run --rm -p 6563:8888 olympicsearch
pause