---
date: 2016-09-17T20:09:24+02:00
title: kaigara
type: index
weight: 0
---
# kaigara
## Devops Swiss Army knife

[![Build Status](https://travis-ci.org/mod/kaigara.svg?branch=master)](https://travis-ci.org/mod/kaigara)

### Introduction

Kaigara is a tiny provisioning tool for container initialization
and with continuous migration in production in mind.

### Installation

Paste the following lines into your Dockerfile:
```
## <[ Kaigara
ENV KAIGARA_VERSION v0.0.1
RUN wget --quiet https://github.com/mod/kaigara/releases/download/$KAIGARA_VERSION/kaigara-linux-amd64-$KAIGARA_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf kaigara-linux-amd64-$KAIGARA_VERSION.tar.gz

    COPY operations /opt/provision/operations
    COPY resources /opt/provision/resources
## Kaigara ]>
```

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

### Contribute

```
go get github/mod/kaigara
```
