clean:
	@echo "Cleaning previus build"
	rm backend/static -rf && mkdir backend/static
	rm bin/ -rf && mkdir bin/

build:
	$(MAKE) clean
	@echo "Building frontend"
	cd frontend/webapp && npm run build
	cp -r frontend/webapp/dist/* backend/static
	cp backend/config/dev.yaml backend/static
	@echo "Building the backend app (LINUX)"
	$(MAKE) build_linux
	@echo "Building the backend app (WINDOWS)"
	$(MAKE) build_windows
	mv backend/taskr bin/
	mv backend/taskr.exe bin/
	# Remove the static folder
	rm backend/static -rf && mkdir backend/static
	@echo "Build completed, please run ./bin/taskr"
	cd bin/ && ls -l --block-size=M

build_linux:
	cd backend && packr build -ldflags "-s -w"
build_windows:
	cd backend && GOOS=windows GOARCH=386 packr build -ldflags "-s -w"

.DEFAULT_GOAL := build