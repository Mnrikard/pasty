main_package_path = ./
binary_name = pasty

.PHONY: build
build:
	go build -o=/tmp/bin/${binary_name} ${main_package_path}

