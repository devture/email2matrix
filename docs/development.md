# Development / Experimenting

If you'd like to contribute code to this project or give it a try locally (before deploying it), you need to:

- clone this repository

- get [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/) -- used for running a local Matrix Synapse + riot-web setup, for testing

- start all dependency services (Postgres, Matrix Synapse, riot-web): `make services-start`. You can stop them later with `make services-stop` or tail their logs with `make services-tail-logs`

- create a sample "sender" user: `make create-sample-sender-user`

- create a sample "receiver" user: `make create-sample-receiver-user`

- you should now be able to log in with user `receiver` and password `password` to the [riot-web instance](http://email2matrix.127.0.0.1.xip.io:41465)

- using riot-web, from that receiver (`receiver`) user: create a room or two with the `sender` user (full Matrix user id is: `@sender:email2matrix.127.0.0.1.xip.io`)

- in another browser session (new container tab, private tab, another browser, etc.), log in to [riot-web](http://email2matrix.127.0.0.1.xip.io:41465) with user `sender` and password `password` and accept those room invitations

- copy the sample configuration: `cp config.json.dist config.json`

- obtain an access token for the `sender` user using `make obtain-sample-sender-access-token`. You will need the value of the `access_token` field below

- create a new mapping in `config.json` (see the [Configuration documentation](configuration.md))

Example:
```json
{
	"Mappings": [
		{
			"MailboxName": "test",
			"MatrixRoomId": "!ROOM_ID_HERE:email2matrix.127.0.0.1.xip.io",
			"MatrixHomeserverUrl": "http://synapse:8008",
			"MatrixUserId": "@sender:email2matrix.127.0.0.1.xip.io",
			"MatrixAccessToken": "SENDER_ACCESS_TOKEN_HERE",
			"IgnoreSubject": false,
			"IgnoreBody": false,
			"SkipMarkdown": false
		}
	]
}
```

- build and run the `email2matrix` program by executing: `make run-in-container`

- send a test email by executing: `make send-sample-email-to-test-mailbox`

- you should now see that email message relayed to the Matrix room created above


For local development, it's best to install a [Go](https://golang.org/) compiler (version 1.12 or later is required) locally.
Some tests are available and can be executed with: `make test`.
