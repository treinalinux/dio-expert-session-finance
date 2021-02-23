# build-image: faz o build da nossa imagem docker.
build-image:
	docker build -t treinalinux/finance .

# run-app : executa nosso docker-compose a partir do app.yml, levantando no container em background
run-app:
	docker-compose -f .devops/app.yml up -d

# lint : checar erros no go e assim mantem um padrão de qualidade.
lint:
	golint ./...
	go fmt -n ./...

# unit_test : chama e executa os testes que tivermos em nosso código go.
unit_test:
	go test ./...
