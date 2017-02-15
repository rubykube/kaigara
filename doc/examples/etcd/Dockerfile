FROM quay.io/coreos/etcd:v0.4.6

# Kaigara setup
ADD operations /opt/provision/operations
ADD resources /opt/provision/resources
ADD defaults.yml /opt/provision

ADD kaigara /usr/local/bin/

EXPOSE 5001 8001
