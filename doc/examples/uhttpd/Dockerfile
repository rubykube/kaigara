FROM alpine
MAINTAINER Louis Bellet

# Kaigara setup
ADD operations /opt/kaigara/operations
ADD resources /opt/kaigara/resources
ADD scripts/kaigara /usr/local/bin/
ADD uhttpd /usr/local/bin/
ADD scripts/entrypoint.sh /usr/local/bin/

EXPOSE 9090

VOLUME ["/var/lib/uhttpd", "/etc/uhttpd"]

ENTRYPOINT ["entrypoint.sh"]

CMD ["kaigara", "execute", "uhttpd"]
