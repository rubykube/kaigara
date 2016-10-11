---
date: 2016-09-17T20:27:11+02:00
title: Getting Started
---
## How to install

### Using wget
```
export KAIGARA_VERSION=v0.0.2
wget --quiet https://github.com/mod/kaigara/releases/download/$KAIGARA_VERSION/kaigara-linux-amd64-$KAIGARA_VERSION.tar.gz
```

### Using go
```
go get -v -u github.com/mod/kaigara
```

## Create an Application

```
mkdir -p work/application/myapp
cd work/application/myapp
kaigara init
```

### Run inside a container

Kaigara should be the default entrypoint
you can access any command using CMD

```
# Display help
docker run -it --rm kaigara/box help

# Run all operations
docker run -it --rm kaigara/box provision

# Render a template from resources on stdout
docker run -it --rm kaigara/box render config.conf.tmpl
```

### Usage

```
igara is a lightweight provisioning system for unix

  By embeding operations and resources into your application containers
  kaigara will run all your provisioning scripts before starting the app.

Usage:
  kaigara [command]

Available Commands:
  create      Create a Kaigara default docker project
  exec        Execute an executable in a child process
  provision   Provisioning using operations
  render      Generate a file from a template
  start       Runs a <command> after provision

Flags:
      --color           enable colorized output
      --config string   config file (default is $HOME/.kairc)
  -h, --help            help for kaigara

Use "kaigara [command] --help" for more information about a command.

```
