
setup:
	go get github.com/pilu/fresh && \
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.44.0 &&\
    chmod +x quality.sh

quality:
	./quality.sh

run:
	fresh -c fresh-custom-runner.conf