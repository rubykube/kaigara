# kaigara
## Devops Swiss Army knife

[![Build Status](https://travis-ci.org/mod/kaigara.svg?branch=master)](https://travis-ci.org/mod/kaigara)

### Introduction

Kaigara is a tiny provisioning tool for container initialization
and with continuous updating in mind.

Documentation is available at [Kaigara.org](http://www.kaigara.org/)

### Installation

Paste the following lines into your Dockerfile:
```
## Kaigara being
  RUN curl -sL https://kaigara.org/get | sh
  COPY operations /opt/kaigara/operations
  COPY resources /etc/kaigara/resources
## Kaigara end
```

### Philosophy

Kaigara is a standard way for a container ENTRYPOINT. It can setup or upgrade
your application configuration before starting the process.

Kaigara operate with 3 main components Metadata, Operations and Resources.

*Metdata*: In development you can set default variables in a default.yaml file and mount it
In production or other environments mount the production config as well.
Default Folder: /etc/kaigara/metadata/

*Resources*: A directory copied on the target system containing templates or files to by manipulated
from Operations. The templates will be rendered using application metadata.
Default Folder: /etc/kaigara/resources/

*Operations*: A directory of scripts copied into the target system, they will be executed one by one in
alphabetical orders, we would usually use first operation for setting up application volumes for
persistant and application specific configurations. The sub-sequents operations are
configuration changes and upgrades, similar to database migrations.
Default Folder: /opt/kaigara/operations/

### Contribute

```
go get github/mod/kaigara
```
