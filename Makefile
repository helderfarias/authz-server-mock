default: clean build clean_build

clean:
	@echo "Clean all..."
	@rm authz-server-mock_*.zip || true
	@rm authz-server-mock || true

build: linux_64 darwin_osx
	@echo "Complete"

clean_build:
	@rm authz-server-mock || true

linux_64: clean_build
	@echo "Building for linux 64bits..."
	@GOOS=linux GOARCH=amd64 govendor build +local && zip authz-server-mock_linux_64.zip authz-server-mock

darwin_osx: clean_build
	@echo "Building for Darwin/OSX"
	@GOOS=darwin govendor build +local && zip authz-server-mock_darwin_osx.zip authz-server-mock