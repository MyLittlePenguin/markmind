MAIN := cmd/main/main.go

.PHONY: run
run: templ tailwind
	go run ${MAIN}

.PHONY: build
build: templ tailwind
		go build -o bin/markmind ${MAIN}

.PHONY: templ
templ:
	templ generate

.PHONY: test
test:
	go test ./test/...
######

.PHONY: tailwind
tailwind:
	tailwindcss -i static/css/custom.css -o static/css/style.css

.PHONY: tailwind_watch
tailwind_watch:
	tailwindcss -w -o static/css/style.css
