install-dep:
		go get -u github.com/golang/dep/cmd/dep

deps: install-dep
		dep ensure

deps-update: install-dep
		rm -rf ./vendor
		dep ensure -update

build:
		go build -ldflags="-w -s" -o bin/go-skilltest \
			main.go    \
			path.go    \
			helper.go  \
			server.go  \
			history.go \
			command.go

test:
		go test
