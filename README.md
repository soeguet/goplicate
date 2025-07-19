# goplicate

small cli utility to check if any duplicates reside within the current working directory (recursivly)

## features
list out name, filepath and hash for duplicates

## build
```sh
go build -o out/goplicate
```

### Makefile
```sh
make build
```
### Dockerfile
#### docker
```sh
docker build -t goplicate:v1 .
docker run goplicate:v1
```
#### podman
```sh
podman build -t goplicate:v1 .
podman run goplicate:v1
```

## extra
### goroutines
#### without goroutines
```
Counter:  2223
Time passed: 284.386941ms
```
#### with goroutines
```
Counter:  2224
Time passed: 97.631632ms
```