clean:
	go clean
	rm -r -f ./bin

compile: 
	@make clean
	@templ generate
	mkdir ./bin
	go build -o ./bin ./app

compile_linux: 
	@make clean
	@templ generate
	mkdir ./bin
	GOARCH=amd64 GOOS=linux go build -o ./bin ./app

run: 
	@make clean
	@npm run tailwind
	@npm run esbuild
	@templ generate
	@make compile
	./bin/app

install_deps:
	@go install github.com/bokwoon95/wgo@latest
	@go install github.com/a-h/templ/cmd/templ@latest
	@npm install

dev:
	@wgo -file=.go -file=.templ -file=.js -file=.ts -xdir=public -xfile=_templ.go templ generate :: npm run tailwind :: npm run esbuild :: go run app/main.go


