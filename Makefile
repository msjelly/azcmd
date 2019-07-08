
# Image URL to use all building/pushing image targets
IMAGE ?= azcmd:latest

PROJECT_PATH = github.com/msjelly/azcmd
all: dependencies azcmd

# Install all go dependencies
dependencies: 
	./hack/update-deps.sh

# Build azcmd binary
azcmd: fmt vet
	go build -o bin/azcmd ${PROJECT_PATH}/cmd/azcmd

# Run against the configured Kubernetes cluster in ~/.kube/config
run: fmt vet
	go run ./cmd/azcmd/main.go 

# Setup local debugging of azcmd using ~/.kube/config
# debug: azcmd deploy-dev

# Run go fmt against code
fmt:
	go fmt -x ./pkg/... ./cmd/...

# Run go vet against code
vet:
	go vet ./pkg/... ./cmd/...

# Build the docker image
docker-build: 
	docker build . -t ${IMAGE}
	
# Push the docker image
docker-push:
	docker push ${IMAGE}

