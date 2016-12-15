---
date: 2016-09-17T20:27:11+02:00
title: Getting Started
---
## How to install

### Using inline script
```
curl https://kaigara.org/get | sh
```
### Using wget
```
export KAIGARA_VERSION=v0.0.4
wget --quiet https://github.com/mod/kaigara/releases/download/$KAIGARA_VERSION/kaigara-linux-amd64-$KAIGARA_VERSION.tar.gz
```

### Using go
```
go get -v -u github.com/mod/kaigara
```

## Create an Application tree

```
mkdir -p work/application/myapp
cd work/application/myapp
mkdir resources operations
touch defaults.yml
vim Dockerfile
```

## Edit your Dockerfile

Paste the following line in your dockerfile as root user
```
## <[ Kaigara
RUN curl https://kaigara.org/get | sh

COPY operations   /opt/kaigara/operations
COPY resources    /etc/kaigara/resources
COPY defaults.yml /etc/kaigara/defaults.yml

ENTRYPOINT ["kaigara"]
CMD ["start", "myapp.sh"]
## Kaigara ]>
```

You can build your container :
```
docker build -t company/myapp:0.0.1 --rm .
```

### Run inside a container

Kaigara should be the default entrypoint
you can access any command using CMD

```
# Display help
docker run -it --rm company/myapp help

# Run all operations
docker run -it --rm company/myapp provision

# Render a template from resources on stdout
docker run -it --rm company/myapp render config.conf

# Execute bash or any unix program
docker run -it --rm company/myapp exec bash

# Display installed kaigara version
docker run -it --rm company/myapp version
```

### Usage

```
Kaigara is a lightweight provisioning system for unix

By embeding operations and resources into your application containers
kaigara will run all your provisioning scripts before starting the app.

Usage:
  kaigara [command]

Available Commands:
  exec        Execute an executable in a child process
  provision   Provisioning using operations
  render      Generate a file on STDOUT from a TEMPLATE
  start       Runs a <PROGRAM> after provision
  version     Return kaigara version

Use "kaigara [command] --help" for more information about a command.
```
