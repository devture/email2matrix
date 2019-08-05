# Architecture


```
 Email sending system
   |
   |
   |          +---------------------+                     +----------------------+
   |          |                     |                     |                      |
   | SMTP 25  |     Email2Matrix    | HTTPS (client API)  |   Matrix Homeserver  |
   +--------> |                     | ------------------> |    (e.g. Synapse)    |
              |                     |                     |                      |
              +---------------------+                     +----------------------+
```

Things to note:

- `email2matrix` receives email messages sent from another system

- the mailbox that a message gets delivered to (e.g. `mailbox5@email2matrix.example.com`) designates where the message will be forwarded to on the Matrix side (such mappings are defined in `config.json` manually)

- `email2matrix` then uses the [Matrix Client-Server API](https://matrix.org/docs/spec/client_server/r0.5.0) with a pre-created user and access token in order to send a Matrix message to a specific room (as defined in `config.json`)
