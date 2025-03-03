FROM ubuntu:22.04

COPY ./bin/server /iterx/bin/server
COPY ./config /iterx/config
COPY ./i18n /iterx/i18n

WORKDIR /iterx

CMD ["./bin/server"]