# go-starter

## Resources
+ https://golang.org/doc/code.html
+ https://tour.golang.org/welcome/1

## Testing Locally
1. Build and add to bin:
```
go get -u github.com/golang/dep/cmd/dep
dep ensure
go install
```

2. Can call app name from anywhere now to run 

## Create Docker Image (NEEDS TO BE FIXED)
+ build a local binary (change myapp to name)
```
docker run --rm -it -v "$GOPATH":/gopath -v "$(pwd)":/app -v:"$(pwd)"/plots:/plots -e "GOPATH=/gopath" -w /app golang:1.4.2 sh -c 'CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o dimspan'
```

+ build the docker image (change myapp)
```
docker build -t dimspan .
```

+ run the app
```
docker run -it dimspan
```

+ can push to registry if you want (change myapp)
```
docker tag dimspan maxhoffman/dimspan:latest
docker push maxhoffman/dimspan:latest
```
