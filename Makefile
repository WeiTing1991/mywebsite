run:
	npx tailwindcss --config tailwind.config.js -i ./static/css/input.css -o ./static/css/tailwind.css
	templ generate
	go run cmd/main.go
