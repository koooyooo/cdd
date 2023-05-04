.PHONY: install install-gh

install:
	@ go build -o cdd main.go && mv ./cdd "${GOPATH}/bin"

install-gh:
	@ go install github.com/koooyooo/cdd@latest