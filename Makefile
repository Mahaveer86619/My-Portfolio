build-css:
	npx tailwindcss -i static/css/input.css -o static/css/tailwind.css --minify

run: build-css
	go run main.go

docker-build:
	docker build -t portfolio-app .