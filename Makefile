build_editor:
	echo "Building editor..."
	rm -rf packages/editor/dist/*
	cd packages/editor && yarn build

watch_editor:
	cd packages/editor && yarn dev

build_app:
	echo "Building app..."
	rm -rf views/*.go
	echo "package views" > views/views.go
	go run packages/prebuild/prebuild.go --force
	go run . bundlestatic
	go mod tidy

build_app_dev:
	echo "Building app dev..."
	echo "package views" > views/views.go
	go run packages/prebuild/prebuild.go
	go run . bundlestatic
	go run . run

test_all:
	echo "Testing..."
	go test -coverprofile ./test/coverage.txt ./... && go tool cover -html=./test/coverage.txt -o ./test/coverage.html

prerelease: build_app test_all

releaselocal: prerelease
	echo "Releasing..."
	rm -rf dist
	goreleaser release --snapshot --rm-dist

release: build_editor prerelease
	echo "Releasing..."
	rm -rf dist
	goreleaser release