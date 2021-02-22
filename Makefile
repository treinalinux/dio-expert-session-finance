build-image:
	docker build -t treinalinux/finance .


run-app:
	docker-compose -f .devops/app.yml up -d


lint:
	golint ./...
	go fmt -n ./...


unit_test:
	go test ./...
