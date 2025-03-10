FROM ubuntu:22.04

RUN apt update && apt install -y ca-certificates curl wget vim && \
    update-ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY ./bin/server /iterx/bin/server
COPY ./config /iterx/config
COPY ./i18n /iterx/i18n

WORKDIR /iterx

CMD ["./bin/server"]