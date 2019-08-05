# Email2Matrix: SMTP server relaying messages to Matrix rooms

[email2matrix](https://github.com/devture/email2matrix) is an SMTP server (powered by [Go-Guerrilla](https://github.com/flashmob/go-guerrilla)), which receives messages to certain special (predefined) mailboxes and relays them to [Matrix](http://matrix.org/) rooms.


# Using this Docker image

Start off by [creating your configuration](https://github.com/devture/email2matrix/blob/master/docs/configuration.md).

Since you're running this in a container, make sure `ListenInterface` configuration uses the `0.0.0.0` interface.

To start the container:

```bash
docker run \
-it \
--rm \
-p 127.0.0.1:41080:41080 \
-v /local/path/to/config.json:/config.json:ro \
devture/email2matrix:latest
```

**Hint**: using a tagged/versioned image, instead of `latest` is recommended.
