# Script on Docker Events

Script on Docker Events is a small utility that will execute a set of commands when a Docker event occurs.
I made this to automate a few things on my Unraid server. Rather than maintain your own Docker container for
specific applications on Unraid, this should allow you to rely on the existing community applications without
having to create your own Docker containers to bake in whatever automation you want to add.

## Requirements

* go 1.15+
* bash

## Configuration

The event processing config consists of:

* A free-form identifier for the event. This is for your use, so you should make this as descriptive as you
  need.
* An object type. This corresponds to Docker object types from events, which can be seen in more detail
  [here](https://docs.docker.com/engine/reference/commandline/events/). Note that, on this page, the object
  types are listed as plural (containers, images) but in actuality they are singular (container, image)
* An action which corresponds to a Docker event action. This describes what action has occurred to the
  object type. A few example object type/action pairs would be: `container:start`, `image:pull`.
* A list of commands to run. These will be passed as an argument to `bash -c <command>` and will be run
  asynchronously.

At present, the Docker client is only configurable via the environment.

### Initialization event

There is a special event that can be added to the config called "init." If you label the type and action as
"init," it will be executed prior to any other events. The commands run as part of init are synchronous to ensure
that any following command execute as expected.

## Subcommands

### echo-config

```
script-on-docker-events echo-config --config <config-file>
```

echo-config will take your config file and just echo a processed version of it to standard out. This is
useful for ensuring that things are being parsed as expected.

### process-events

```
script-on-docker-events process-events --config <config-file> [--start-minutes-ago <number>]
```

process-events will take your config and actually monitor docker events and run the configured commands.
By default, process-events will examinute the previous 5 minutes of events on startup. This is to avoid race conditions
where this utility starts up slightly after an intended target container and misses the actual event as it happens.
This time can be adjusted using the `--start-minutes-ago` argument.
