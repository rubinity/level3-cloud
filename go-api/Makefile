RED="\\033[91m"
GRE="\\033[92m"
YEL="\\033[93m"
END="\\033[0m"

API_BIN=api
API_SRC=cmd/api/*.go
INTERNAL=internal/api/api.go internal/api/handlers.go internal/api/helpers.go internal/api/requests.go \
internal/auth/auth.go internal/auth/middleware.go internal/auth/redis.go
IMAGE_NAME= api_image

all: ${API_BIN}

${API_BIN}: ${API_SRC} ${INTERNAL} go.mod go.sum
	GOOS=linux GOARCH=amd64 go build -o ${API_BIN} ${API_SRC} 

run:
	./${API_BIN}

test: 
	@echo "$(YEL)Test API$(END)"
	@go test -coverpkg=./internal/api -coverprofile coverage-unit.out ./internal/api/... || (echo "$(RED)ERROR$(END) unit tests failed"; exit 1)
	@go tool cover -html=coverage-unit.out -o coverage-unit.html
	@echo "Unit test coverage report in $$(pwd)/coverage-unit.html"
	@echo "$(GRE)OK$(END) unit tests passed"


docker:
# 	docker buildx build --platform linux/amd64 -t $(IMAGE_NAME) .
	docker run -d -it -p 8080:80 --name api_cont $(IMAGE_NAME)
# 	docker run -it --name $(IMAGE_NAME)_cont $(IMAGE_NAME)

clean:
	rm -rf ${API_BIN}

./PHONY: run
