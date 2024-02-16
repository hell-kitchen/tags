.PHONY: gen
gen:
	protoc --go_out=. --go_opt=paths=import \
          --go-grpc_out=. --go-grpc_opt=paths=import \
          ./api/tags.proto

.PHONY: dock
dock:
	docker build . --file=infra/tags.dockerfile --tag="vladmarlo/tags_backend:latest"
	docker build . --file=infra/migrator.dockerfile --tag="vladmarlo/tags_migrator:latest"

.PHONY: dock/push
dock/push: dock
	docker push vladmarlo/tags_backend:latest
	docker push vladmarlo/tags_migrator:latest

.PHONY: dock/run
dock/run:
	docker-compose up -d

.PHONY: lines
lines:
	git ls-files | xargs wc -l

.DEFAULT_GOAL := build