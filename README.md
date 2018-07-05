# gofn Example

## Create docker images

```console
docker build -t sum sum/.
docker build -t encode encode/.
docker build -t upper upper/.
```

## Run docker images via gofn

```console
go run main.go -image=sum -p1=2 -p2=2
go run main.go -image=upper -p1="Hello, world"
go run main.go -image=encode -p1="string to be encoded"
```