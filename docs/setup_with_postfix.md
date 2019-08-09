# Running email2matrix on the same host as postfix

It's possible to run email2matrix on the same host with postfix. email2matrix can run on a loopback interface and postfix will forward the messges to email2matrix.
This is a simple example what needs to be configured on email2matrix and postfix side to work together on one host.

## Architecture

```
 Email sending system
   |
   |
   |          +--------------------+ 
   | SMTP 25  |                    |
   +--------> |     Postfix        |
              |        |           |
              | .......|.......... |                      +----------------------+
              |        |           |                      |                      |
              |        V           |  HTTPS (client API)  |   Matrix Homeserver  |
              |    email2matrix    |  ------------------> |    (e.g. Synapse)    |
              |                    |                      |                      |
              +--------------------+                      +----------------------+
```

## Configuration

## email2maitrix

Make sure you run email2matrix on a unused port. For example on port 2525. You can even run it on a lookback interface.

Example config (partly):
```
{
  "Smtp": {
    "ListenInterface": "127.0.0.1:2525",
    "Hostname": "email2matrix.example.com",
    "Workers": 10
  },
  "Matrix": ...
```

## Postfix

* Postfix needs to know that he is responsible for the domain
  * main.cf: `mydestination = email2matrix.example.com, localhost`
* Set the transport maps on the destination domain
  * main.cf: `transport_maps = hash:/etc/postfix/transport`
  * transport `email2matrix.example.com	smtp:127.0.0.1:2525`
    * create lookup db from plain file `postmap transport` (this creates transport.db)
* Restart postfix
* Look at `/var/log/mail.log` for errors
