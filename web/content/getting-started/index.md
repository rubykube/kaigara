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

