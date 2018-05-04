build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build service.go types.go

image: build
	docker build -t etcdlet .
	docker save etcdlet -o etcdlet.tar
