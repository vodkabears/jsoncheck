VERSION_FLAG := main.version=$(shell git describe --tags --abbrev=0)
DIST_FOLDER := dist

.PHONY: githooks
githooks:
	cp -f githooks/* .git/hooks/

.PHONY: install
install: githooks
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install
	go get -u github.com/mitchellh/gox
	go get -u github.com/xeipuuv/gojsonschema
	go get -u github.com/mgutz/ansi
	go install ./...

.PHONY: lint
lint:
	gometalinter ./... --enable-all --line-length=100 --vendor --sort=path --sort=line --sort=column --deadline=5m -t

.PHONY: clean
clean:
	rm -rf ${DIST_FOLDER}

.PHONY: build
build: clean
	gox -ldflags "-X ${VERSION_FLAG}" -os="linux darwin windows" -arch="386 amd64" -output="${DIST_FOLDER}/{{.Dir}}_{{.OS}}_{{if eq .Arch \"386\"}}i386{{else}}{{.Arch}}{{end}}" ./...

.PHONY: run
run:
	go run -ldflags "-X ${VERSION_FLAG}" *.go ${ARGS}
