build:
	@echo Vamos compilar, galerinha
	go build -o ftpd.exe  main.go

run: build
	./ftpd.exe

fmt: 
	gofmt -w main.go