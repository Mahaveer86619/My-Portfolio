build-css:
	npx tailwindcss -i static/css/input.css -o static/css/tailwind.css --minify

start-app: build-css
	docker compose up --build

docker-build:
	docker build -t portfolio-app .


