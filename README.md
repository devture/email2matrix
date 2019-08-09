# Email2Matrix: SMTP server relaying messages to Matrix rooms

[email2matrix](https://github.com/devture/email2matrix) is an SMTP server (powered by [Go-Guerrilla](https://github.com/flashmob/go-guerrilla)), which receives messages to certain special (predefined) mailboxes and relays them to [Matrix](http://matrix.org/) rooms.

This is useful when you've got a system which is capable of sending email (notifications, reminders, etc.) and you'd like for that system to actually send a Matrix message instead.

Instead of redoing such systems (adding support for sending messages over the [Matrix](https://matrix.org) protocol to each one), you can just configure them to send emails to the Email2Matrix server and have those email messages relayed over to Matrix.

To learn more, refer to the [Documentation](./docs/README.md).


## Support

Matrix room: [#email2matrix:devture.com](https://matrix.to/#/#email2matrix:devture.com)

Github issues: [devture/email2matrix/issues](https://github.com/devture/email2matrix/issues)
