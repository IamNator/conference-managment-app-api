
setup:
	go install github.com/pilu/fresh@latest && \
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.44.0 &&\
    chmod +x quality.sh &&\
    go install github.com/golang/mock/mockgen@latest


quality:
	./quality.sh

run:
	fresh -c fresh-custom-runner.conf