# Email2Matrix Setup

The easiest way to run Email2Matrix is using a [Docker](https://www.docker.com/) (or other) container.

If your Matrix server is installed using [matrix-docker-ansible-deploy](https://github.com/spantaleev/matrix-docker-ansible-deploy), it's easiest if you refer to [the playbook's setup instructions for Email2Matrix](https://github.com/spantaleev/matrix-docker-ansible-deploy/blob/master/docs/configuring-playbook-email2matrix.md) instead of following these.


## Prerequisites

- [Docker](https://www.docker.com/) or other (e.g. [Podman](https://podman.io/)) container runtime engine

- **port 25 available** on the host machine where you'd like to run this. If not available, you may also be able to make your existing email server on port 25 relay messages to Email2Matrix running on another port (e.g. 2525), but that's a more complicated setup. A simplified documentation can be found [here](setup_with_postfix.md)

- Matrix server configuration:

	- you would most likely wish to create one or more dedicated users that would be doing the sending (e.g. `@email2matrix:DOMAIN`)

	- for each of those sender users, you would need to obtain a Matrix Access Token manually. This can happen with a command like this:

		```
		curl \
		--data '{"identifier": {"type": "m.id.user", "user": "email2matrix" }, "password": "MATRIX_PASSWORD_FOR_THE_USER", "type": "m.login.password", "device_id": "Email2Matrix", "initial_device_display_name": "Email2Matrix"}' \
		https://matrix.DOMAIN/_matrix/client/r0/login
		```

- A configuration file (`config.json`) created as shown in the [Configuration](configuration.md) documentation page

	- `Smtp.Hostname` in your configuration file should either match the hostname leading to your server (using a regular `A` record), or it should be an MX record that eventually leads to your server. In any case, what you see in `Smtp.Hostname` is the domain that needs to be in the `to` field of the emails you would be sending

	- some mappings need to be defined (mailbox names and where those mailboxes lead on the Matrix side)


## Running

```
docker run \
	-it \
	--rm \
	--name=email2matrix \
	-p 25:2525 \
	--mount type=bind,src=/path/to/your/config.json,dst=/config.json,ro \
	--network=email2matrix_default \
	devture/email2matrix:latest
```

It's better to use a specific Docker image tag (not `:latest`).
