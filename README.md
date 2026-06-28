# yaca

**Yet Another Chat Application**

A playground project for experimenting with architecture and design.

## Usage

Note that your system may need to support `Taskfile`s to use this application.
`Taskfile` cli may be called either `task` or `go-task`. Examples in this
section use `task` cli but you should use `go-task` if you have it instead.

### CLI

> Using chat application via command line interface.

The CLI uses local file as message storage so server is not required.

#### Posting a message.

Run `task cli-post DATA=<your message>` to post a message.

#### Viewing messages.

Run `task cli-view` to view all messages.
