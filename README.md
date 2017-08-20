# go-starter

## Resources
+ https://golang.org/doc/code.html
+ https://tour.golang.org/welcome/1

## Testing Locally
1. Build and add to bin:
```
go install
```

2. Can call app name from anywhere now to run 

## Create Docker Image
+ build a local binary (change myapp to name)
```
docker run --rm -it -v "$GOPATH":/gopath -v "$(pwd)":/app -e "GOPATH=/gopath" -w /app golang:1.4.2 sh -c 'CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o myapp'
```

+ build the docker image (change myapp)
```
docker build -t myapp .
```

+ run the app
```
docker run -p 8080:8080 myapp
```

+ can push to registry if you want (change myapp)
```
docker tag myapp maxhoffman/myapp:latest
docker push maxhoffman/myapp:latest
```
