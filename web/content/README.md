---
date: 2016-09-17T20:09:24+02:00
title: kaigara
type: index
weight: 0
---
## Devops Swiss Army knife

[![Build Status](https://travis-ci.org/mod/kaigara.svg?branch=master)](https://travis-ci.org/mod/kaigara)

### Introduction

Kaigara is a tiny provisioning tool for container initialization
and with continuous migration in production in mind.

Kaigara is a standard entrypoint for all your containers, it simply run
an ordered list of operations before starting your application.

Operations are a means for setting up, configuring, contextualizing your application
using metadata. Operations also enable container version upgrade managing data migration
or configuration edits.

### Installation

Paste the following lines into your Dockerfile:
```
## <[ Kaigara
ENV KAIGARA_VERSION v0.0.2
RUN wget --quiet https://github.com/mod/kaigara/releases/download/$KAIGARA_VERSION/kaigara-linux-amd64-$KAIGARA_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf kaigara-linux-amd64-$KAIGARA_VERSION.tar.gz

    COPY operations /opt/provision/operations
    COPY resources /opt/provision/resources
    COPY metadata.yml /opt/provision/metadata.yml
## Kaigara ]>
```

### Download

Linux amd64/i386 are available but also Mac OS X release:

[https://github.com/mod/kaigara/releases]

### Philosophy

Kaigara is a standard way for a container ENTRYPOINT. It can setup or upgrade
your application configuration before starting the process.

Kaigara operate with 3 main components Metadata, Operations and Resources.

*Metdata*: In development you can set default variables into your metadata.yml file
In production or other stages you can either use shell environment or better etcd/consul.

*Resources*: A directory copied on the target system containing templates or files to by manipulated
from Operations. The templates will be rendered using application metadata.

*Operations*: A directory of scripts copied into the target system, they will be executed one by one in
alphabetical orders, we would usually use first operation for setting up application volumes for
persistant and application specific configurations. The sub-sequents operations are
configuration changes and upgrades, similar to database migrations.

Kaigara the stupid provisioning tool.

### Contribute

```
go get -v github.com/mod/kaigara
```

