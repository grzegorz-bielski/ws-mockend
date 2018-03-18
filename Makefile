dev-dockerfile = -f docker-compose.yml -f docker-compose.dev.yml
prod-dockerfile = -f docker-compose.yml

.PHONY: build-dev
build-dev:
	docker-compose $(dev-dockerfile) build
	$(MAKE) install-node-dependencies

.PHONY: build-prod
build-prod:
	docker-compose $(prod-dockerfile) build
	$(MAKE) install-node-dependencies
	$(MAKE) build-statics

.PHONY: install-node-dependencies
install-node-dependencies:
	docker-compose run --rm --no-deps mock-client npm install

.PHONY: build-statics
build-statics:
	docker-compose run --rm --no-deps mock-client npm run build

.PHONY: dev
dev:
	docker-compose $(dev-dockerfile) up --remove-orphans

.PHONY: prod
prod:
	docker-compose $(prod-dockerfile) up -d --remove-orphans
