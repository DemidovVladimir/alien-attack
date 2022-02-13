.PHONY: build

dependencies:
	go mod download

build-mocks:
	@go get github.com/golang/mock/gomock
	@go install github.com/golang/mock/mockgen
	@~/go/bin/mockgen -source=entity/world/interface.go -destination=entity/world/mock/world.go -package=mock
	@~/go/bin/mockgen -source=entity/alien/interface.go -destination=entity/alien/mock/alien.go -package=mock

test:
	go test -v ./... -coverprofile=coverage.out

run-dev: 
	go run ./... -s=5 -w=static/world.txt -a=static/aliens.txt

build:
	go install github.com/VladimirDemidov/alien-attack  

run-prod:
	alien-attack -s=5 -w=static/world.txt -a=static/aliens.txt 