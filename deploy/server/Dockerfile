FROM ghcr.io/iter-x/iter-x:builder AS builder

COPY ../.. /iterx

WORKDIR /iterx

RUN make init
RUN make all
RUN make build

FROM debian:stable-slim

COPY --from=builder /iterx/bin /iterx/bin
COPY --from=builder /iterx/config /iterx/config
COPY --from=builder /iterx/i18n /iterx/i18n
COPY --from=builder /iterx/swagger /iterx/swagger

WORKDIR /iterx

RUN apt update && apt install -y ca-certificates curl wget vim && \
    update-ca-certificates && \
    rm -rf /var/lib/apt/lists/*

CMD ["./bin/server"]

