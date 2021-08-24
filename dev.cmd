docker build --file=Dockerfile.dev -t golang-hello-world . 
docker run --rm -it -v "%~dp0:/go/src/github.com/dirkarnez/golang-hello-world" golang-hello-world
pause