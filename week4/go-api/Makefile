API_BIN=api
API_SRC=cmd/api/*.go
INTERNAL=internal/api/api.go internal/api/handlers.go internal/api/helpers.go internal/api/requests.go
IMAGE_NAME= api_image

all: ${API_BIN}

${API_BIN}: ${API_SRC} ${INTERNAL} go.mod go.sum
	GOOS=linux GOARCH=amd64 go build -o ${API_BIN} ${API_SRC} 

run:
	./${API_BIN}

docker:
# 	docker buildx build --platform linux/amd64 -t $(IMAGE_NAME) .
	docker run -d -it -p 8080:80 --name api_cont $(IMAGE_NAME)
# 	docker run -it --name $(IMAGE_NAME)_cont $(IMAGE_NAME)

clean:
	rm -rf ${API_BIN}

./PHONY: run
