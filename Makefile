env_dir := $(shell go env GOPATH)
gob_dir := $(env_dir)/bin
sources := $(shell find . -type f -name "*.go")

$(gob_dir)/goauth: main/goauth.go $(sources)
	go build -o $@ $<

clean:
	$(RM) ${gob_dir}/goauth
