setup:
	go install github.com/mitranim/gow@latest
	go install github.com/pyros2097/gromer/cmd/gromer@latest
	npm i -g tailwindcss

update:
	gromer

run: export PORT=3000
run:
	gow run main.go

css:
	tailwindcss -i ./assets/css/config.css -o ./assets/css/styles.css --watch

docker-build:
	docker build -t pyros.sh:latest .

docker-run:
	docker run -p 3000:3000 -e PORT=3000 pyros.sh:latest