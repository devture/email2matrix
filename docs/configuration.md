# Email2Matrix Configuration

The `email2matrix` server configuration is a JSON document that looks like this:

```json
{
	"Smtp": {
		"ListenInterface": "0.0.0.0:25",
		"Hostname": "email2matrix.example.com",
		"Workers": 10
	},
	"Matrix": {
		"Mappings": [
			{
				"MailboxName": "test",
				"MatrixRoomId": "!ABCD:example.com",
				"MatrixHomeserverUrl": "https://matrix.example.com",
				"MatrixUserId": "@svc.test.sender:example.com",
				"MatrixAccessToken": "TOKEN_GOES_HERE",
				"IgnoreSubject": false,
				"IgnoreBody": false,
				"SkipMarkdown": false
			}
		]
	},
	"Misc": {
		"Debug": true
	}
}

```

## Fields

The configuration contains the following fields:

- `Smtp` - SMTP server related configuration

	- `ListenInterface` - the network address and port to listen on. If you're running this inside a container, use something like `0.0.0.0:25` (or `0.0.0.0:2525` with `docker run -p 25:2525 ...`).

	- `Hostname` - the hostname of this email server

	- `Workers` - the number of workers to run. Controls how many messages can be received simultaneously.


- `Matrix` - Matrix-related configuration

	- `Mappings` - a list of mappings specifying messages going to which mailbox should get sent to which Matrix room, using what kind of credentials, etc.

		- `MailboxName` - the mailbox name (e.g. `mailbox5`). Its full email address would be `MailboxName@Hostname` (`Hostname` is the `Smtp.Hostname` configuration value). All emails received to this mailbox will be forwarded to Matrix. Use a long and ungueassable name for the mailbox to prevent needless spam.

		- `MatrixRoomId` - the Matrix room id where messages should be sent

		- `MatrixHomeserverUrl` - the homeserver through which to send Matrix messages

		- `MatrixUserId` - the full user id through which the Matrix messages would be sent

		- `MatrixAccessToken` - a Matrix access token corresponding to the sender Matrix user (`MatrixUserId`)

		- `IgnoreSubject` - (`true` or `false`) specifies whether the subject should be forwarded to Matrix

		- `IgnoreBody` - (`true` or `false`) specifies whether the message body should be forwarded to Matrix

		- `SkipMarkdown` - (`true` or `false`) specifies whether to send a plain text message or a Markdown (actually HTML) message to Matrix

- `Misc` - miscellaneous configuration

	- `Debug` - whether to enable debug mode or not (enable for more verbose logs)
