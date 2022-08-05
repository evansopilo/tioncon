FROM debian:stretch-slim

WORKDIR /

ADD bin /bin/

CMD ["/bin/api"]
