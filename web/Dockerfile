FROM debian:latest
MAINTAINER MoD <mod@kaigara.org>

ENV HUGO_VERSION=0.16

RUN apt-get update -y \
      && apt-get install -y \
      wget git ca-certificates \
      && rm -rf /var/lib/apt/lists/*

RUN wget https://github.com/spf13/hugo/releases/download/v${HUGO_VERSION}/hugo_${HUGO_VERSION}_linux-64bit.tgz && \
      tar xzf hugo_${HUGO_VERSION}_linux-64bit.tgz && \
      rm -r hugo_${HUGO_VERSION}_linux-64bit.tgz && \
      mv hugo /usr/local/bin/hugo

RUN groupadd web
RUN useradd -g web -m web

USER web
ADD . /home/web

WORKDIR /home/web

EXPOSE 1313

ENTRYPOINT ["hugo"]

CMD ["server"]
