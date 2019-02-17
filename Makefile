.PHONY: install test build serve clean pack deploy ship

TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)

export TAG

install:
	go get ./cmd/app

test: install
	go test ./...

build: install
	go build -ldflags "-X main.version=$(TAG)" -o app cmd/app/*

serve: build
	./app

clean:
	rm ./app

clean-watchtowers:
	kubectl delete deployment -n watchtower --selector=app=watchtower

pack:
	GOOS=linux make build
	docker build -t mattgarnett/auto-k8s:$(TAG) .

upload:
	docker push mattgarnett/auto-k8s:$(TAG)

deploy:
	kubectl delete deployment -n watchtower --all
	sed 's/REPLACE_WITH_TAG/$(TAG)/' kubernetes/server-deployment.yml | kubectl apply -f -
	kubectl apply -f kubernetes/server-load-balancer-service.yml

ship: pack upload deploy clean
