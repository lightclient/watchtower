.PHONY: install test build serve clean pack deploy ship

TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)

export TAG

install:
	go get ./manager/cmd/app

test: install
	go test ./...

build: install
	go build -ldflags "-X main.version=$(TAG)" -o manager/app manager/cmd/app/*

serve: build
	./app

clean:
	rm ./manager/app

clean-watchtowers:
	kubectl delete deployment -n watchtower --selector=app=watchtower

pack:
	GOOS=linux make build
	docker build -t mattgarnett/auto-k8s:$(TAG) manager

upload:
	docker push mattgarnett/auto-k8s:$(TAG)

deploy:
	kubectl delete deployment -n watchtower server-deployment
	sed 's/REPLACE_WITH_TAG/$(TAG)/' manager/kubernetes/server-deployment.yml | kubectl apply -f -
	kubectl apply -f manager/kubernetes/server-load-balancer-service.yml

ship: pack upload deploy clean


build-watchtower:
	docker build -t mattgarnett/watchtower:$(TAG) .
