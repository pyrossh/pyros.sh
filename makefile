setup:
	go install github.com/mitranim/gow@latest
	go install github.com/pyros2097/gromer/cmd/gromer@latest

update:
	gromer

dev:
	gow run main.go

build:
	podman build -t pyros.sh:latest .

run:
	podman run -p 3000:3000 pyros.sh:latest

# TODO: Fix lighthouse perfs regarding images