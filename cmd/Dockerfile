FROM golang:1.19

ARG apiPort

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV BUNDEBUG=2

EXPOSE ${apiPort}

COPY start.sh /usr/bin/start.sh
RUN chmod +x /usr/bin/start.sh

ENTRYPOINT ["/usr/bin/start.sh"]