# FAQ

## What is Email2Matrix?

Email2Matrix is an SMTP email server program, which receives messages to special predefined mailboxes (in your `config.json` [configuration file](./configuration.md)) and forwards those messages to rooms over the [Matrix](https://matrix.org) protocol.


## When do I need Email2Matrix?

You need Email2Matrix when you've got a system which is capable of sending email (notifications, reminders, etc.) and you'd like for that system to actually send a Matrix message instead.

Instead of redoing such systems (adding support for sending messages over the [Matrix](https://matrix.org) protocol to each one), you can just configure them to send emails to the Email2Matrix server and have those email messages relayed over to Matrix.


## Do I need to use a specific homeserver for this to work?

No. You can send via any homeserver to any room, even across federation.


## Why do I need to create rooms, access tokens, etc., manually?

This is built in the simplest way possible.
You need to do more manual work, so that Email2Matrix doesn't and can be kept simple.

A more easy to use version can be created.
One which acts as a bot, auto-accepts room invitations, manages mailbox mappings in a database, etc., but that hasn't been done yet.


## Do I need special DNS configuration to use this?

Not necessarily.

You can use `MX` DNS records if you wish, but you can avoid it as well.

If you're hosting Email2Matrix on another server which already has a DNS mapping for its hostname, you can [Configure](configuration.md) your `Smtp.Hostname` to match the hostname of the machine (example: `matrix.example.com`).

On the other hand, if you'd like to have more indirection, feel free to use `MX` DNS records.

## Can I run email2matrix on the same host with postfix?

Yes. Here is the [documentation describing that](setup_with_postfix.md).
