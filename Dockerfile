FROM alpine:latest

ADD . /opt/.

ADD animal  /usr/local/bin/

RUN chmod +x /usr/local/bin/animal

WORKDIR /opt/
