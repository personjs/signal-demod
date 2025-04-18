all: clean build save

build:
	docker compose build
save:
	docker save signal-demod -o signal-demod.tar
clean:
	@docker system prune -f