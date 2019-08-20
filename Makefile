help: ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

_prepare_services: var/.env
	mkdir -p var/matrix-synapse-media-store var/matrix-synapse-postgres

var/.env:
	mkdir -p var
	echo 'CURRENT_USER_UID='`id -u` > var/.env;
	echo 'CURRENT_USER_GID='`id -g` >> var/.env

services-start: _prepare_services ## Starts all services (Postgres, Synapse, Riot)
	docker-compose --project-directory var -f etc/services/docker-compose.yaml -p email2matrix up -d

services-stop: _prepare_services ## Stops all services (Postgres, Synapse, Riot)
	docker-compose --project-directory var -f etc/services/docker-compose.yaml -p email2matrix down

services-tail-logs: _prepare_services ## Tails the logs for all running services
	docker-compose --project-directory var -f etc/services/docker-compose.yaml -p email2matrix logs -f

create-sample-receiver-user: _prepare_services ## Creates a receiver user
	docker-compose --project-directory var -f etc/services/docker-compose.yaml -p email2matrix \
		exec synapse \
		register_new_matrix_user \
		-a \
		-u receiver \
		-p password \
		-c /data/homeserver.yaml \
		http://localhost:8008

create-sample-sender-user: _prepare_services ## Creates a sender user
	docker-compose --project-directory var -f etc/services/docker-compose.yaml -p email2matrix \
		exec synapse \
		register_new_matrix_user \
		-a \
		-u sender \
		-p password \
		-c /data/homeserver.yaml \
		http://localhost:8008

obtain-sample-sender-access-token: _prepare_services ## Obtain an access token for the sender user
	docker run \
	-it \
	--rm \
	--network=email2matrix_default \
	alpine:3.10 \
	/bin/sh -c "apk add --no-cache curl && curl --data '{\"identifier\": {\"type\": \"m.id.user\", \"user\": \"sender\" }, \"password\": \"password\", \"type\": \"m.login.password\", \"device_id\": \"Sender\", \"initial_device_display_name\": \"Sender\"}' http://synapse:8008/_matrix/client/r0/login"

run-locally: build-locally ## Builds and runs email2matrix-server locally (no containers)
	./email2matrix-server

build-locally: ## Builds the email2matrix-server code locally (no containers)
	go get -u -v github.com/ahmetb/govvv
	rm -f email2matrix-server
	go build -a -ldflags "`~/go/bin/govvv -flags`" email2matrix-server.go

test: ## Runs the tests locally (no containers)
	go test ./...

build-container-image: ## Builds a Docker container image
	docker build -t devture/email2matrix:latest -f etc/docker/Dockerfile .

run-in-container: build-container-image ## Runs email2matrix in a container
	docker run \
	-it \
	--rm \
	--name=email2matrix \
	-p 40025:2525 \
	--mount type=bind,src=`pwd`/config.json,dst=/config.json,ro \
	--network=email2matrix_default \
	devture/email2matrix:latest

send-sample-email-to-test-mailbox: ## Sends a sample email to email2matrix
	docker run \
	-it \
	--rm \
	--network=email2matrix_default \
	alpine:3.10 \
	/bin/sh -c "apk add --no-cache ssmtp && sed -i s/mailhub=mail/mailhub=email2matrix:2525/ /etc/ssmtp/ssmtp.conf && echo -e \"Subject: this is the subject\n\nthis is the body\" | ssmtp test@email2matrix.127.0.0.1.xip.io"
