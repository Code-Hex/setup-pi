all: gen pi

gen:
	@go generate
pi:
	@GOOS=linux GOARCH=arm GOARM=7 go build -o setup cmd/setup/main.go
mac:
	@go build -o setup cmd/setup/main.go