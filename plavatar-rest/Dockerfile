FROM alpine:latest

RUN addgroup -S plavatar && adduser -S plavatar -G plavatar

USER plavatar
COPY artifacts/plavatar-rest /usr/bin/plavatar-rest

WORKDIR /etc/plavatar/

USER root
RUN chown -R plavatar:plavatar /etc/plavatar/
RUN chmod 755 /usr/bin/plavatar-rest
#RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

USER plavatar
ENTRYPOINT /usr/bin/plavatar-rest
